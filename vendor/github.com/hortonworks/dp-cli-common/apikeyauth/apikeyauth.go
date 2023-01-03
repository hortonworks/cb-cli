package apikeyauth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"

	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	ed "golang.org/x/crypto/ed25519"
)

const (
	altusActorCrnHeader = "x-cdp-actor-crn"
	altusAuthHeader     = "x-altus-auth"
	altusDateHeader     = "x-altus-date"
	contentTypeHeader   = "content-type"
	signPattern         = "%s\napplication/json\n%s\n%s\n%s"
	layout              = "Mon, 02 Jan 2006 15:04:05 GMT"
	authMethod          = "ed25519v1"
	authMethodECDSA     = "ecdsav1"
)

type metastr struct {
	AccessKey  string `json:"access_key_id"`
	AuthMethod string `json:"auth_method"`
}

func newMetastr(accessKeyID, authMethod string) *metastr {
	return &metastr{accessKeyID, authMethod}
}

func GetAPIKeyAuthTransport(address, baseAPIPath, accessKeyID, privateKey string) *utils.Transport {
	address, basePath := utils.CutAndTrimAddress(address)
	cbTransport := &utils.Transport{client.New(address, basePath+baseAPIPath, []string{"https"})}
	cbTransport.Runtime.DefaultAuthentication = altusAPIKeyAuth(baseAPIPath, accessKeyID, privateKey)
	cbTransport.Runtime.Transport = utils.LoggedTransportConfig
	return cbTransport
}

func GetActorCrnAuthTransport(address, baseAPIPath, actorCrn string) *utils.Transport {
	address, basePath := utils.CutAndTrimAddress(address)
	cbTransport := &utils.Transport{client.New(address, basePath+baseAPIPath, []string{"http"})}
	cbTransport.Runtime.DefaultAuthentication = altusActorCrnAuth(baseAPIPath, actorCrn)
	cbTransport.Runtime.Transport = utils.LoggedTransportConfig
	return cbTransport
}

func altusActorCrnAuth(baseAPIPath, actorCrn string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam(altusActorCrnHeader, actorCrn)
	})

}

func altusAPIKeyAuth(baseAPIPath, accessKeyID, privateKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		date := formatdate()
		err := r.SetHeaderParam(altusAuthHeader, authHeader(accessKeyID, privateKey, r.GetMethod(), resourcePath(baseAPIPath, r.GetPath(), r.GetQueryParams().Encode()), date))
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(contentTypeHeader, "application/json")
		if err != nil {
			return err
		}
		return r.SetHeaderParam(altusDateHeader, date)
	})
}

func resourcePath(baseAPIPath, path, query string) string {
	base := escapePath(strings.ReplaceAll(baseAPIPath+path, "//", "/"))
	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", base, query)
	}
	return base
}

func escapePath(path string) string {
	spl := strings.Split(path, "/")
	encoded := []string{}
	for _, e := range spl {
		encoded = append(encoded, url.PathEscape(e))
	}
	return strings.Join(encoded, "/")
}

func authHeader(accessKeyID, privateKey, method, path, date string) string {
	return fmt.Sprintf("%s.%s", urlsafeMeta(accessKeyID, privateKey), urlsafeSignature(privateKey, method, path, date))
}

func urlsafeSignature(seedBase64, method, path, date string) string {
	seed, err := base64.StdEncoding.DecodeString(seedBase64)
	if err != nil {
		log.Debugf("Error during decoding v2 key, trying v3 format.")
		return urlsafeECDSASignature(seedBase64, method, path, date)
	}
	k := ed.NewKeyFromSeed(seed)
	message := fmt.Sprintf(signPattern, method, date, path, authMethod)
	log.Debugf("Message to sign: \n%s\n", message)
	signature := ed.Sign(k, []byte(message))
	return urlsafeBase64Encode(signature)
}

func urlsafeECDSASignature(privateKey, method, path, date string) string {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		utils.LogErrorAndExit(errors.New("privKey no pem data found"))
	}
	pk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	pkECDSA, ok := pk.(*ecdsa.PrivateKey)
	if !ok {
		utils.LogErrorAndExit(errors.New("cannot convert private key to ecdsa key"))
	}
	// Hash the input to get a summary of the information
	h := sha512.New()
	message := fmt.Sprintf(signPattern, method, date, path, authMethodECDSA)
	log.Debugf("Message to sign: \n%s\n", message)
	_, err = io.WriteString(h, message)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	hash := h.Sum(nil)
	// ECDSA Signing
	signature, err := ecdsa.SignASN1(rand.Reader, pkECDSA, hash)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return urlsafeBase64Encode(signature)
}

func urlsafeMeta(accessKeyID, privatekey string) string {
	_, err := base64.StdEncoding.DecodeString(privatekey)
	method := authMethod
	if err != nil {
		log.Debugf("Cannot decode v2 key, trying v3 format.")
		method = authMethodECDSA
	}
	b, err := json.Marshal(newMetastr(accessKeyID, method))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return urlsafeBase64Encode(b)
}

func urlsafeBase64Encode(data []byte) string {
	return string(base64.URLEncoding.EncodeToString(data))
}

func formatdate() string {
	return string(time.Now().UTC().Format(layout))
}
