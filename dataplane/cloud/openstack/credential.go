package openstack

import (
	"errors"
	"fmt"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

var FACINGS = []string{"public", "admin", "internal"}
var SCOPES = []string{"project", "domain"}

func (p *OpenstackProvider) GetCredentialRequest(stringFinder func(string) string, govCloud bool) (*model.CredentialV1Request, error) {
	return nil, nil
}

func validateAndGet(value string, values []string) (string, error) {
	if len(value) == 0 {
		return values[0], nil
	} else {
		for _, v := range values {
			if v == value {
				return value, nil
			}
		}
	}
	return "", errors.New(fmt.Sprintf("%s not allowed", value))
}
