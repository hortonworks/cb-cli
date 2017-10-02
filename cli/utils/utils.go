package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SafeInt32Convert(value *int32) int32 {
	if value == nil {
		return 0
	}
	return *value
}

func SafeStringConvert(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func EscapeStringToJson(input string) string {
	return strings.Replace(strings.Replace(input, "\\", "\\\\", -1), "\"", "\\\"", -1)
}

func ReadFile(fileLocation string) []byte {
	bp, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		LogErrorAndExit(err)
	}
	return bp
}

func ReadFromUrl(urlLocation string) []byte {
	client := new(http.Client)
	resp, httperr := client.Get(urlLocation)
	if httperr != nil {
		LogErrorAndExit(httperr)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		LogErrorMessageAndExit(fmt.Sprintf("Couldn't download blueprint from URL, response code is not %d.", resp.StatusCode))
	}
	bodyBytes, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		LogErrorAndExit(ioerr)
	}
	return bodyBytes
}
