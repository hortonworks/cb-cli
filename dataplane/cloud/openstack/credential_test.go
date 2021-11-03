package openstack

import (
	"testing"

	"github.com/hortonworks/cb-cli/dataplane/cloud"
)

var provider cloud.CloudProvider = new(OpenstackProvider)

func init() {
	cloud.SetProviderType(cloud.OPENSTACK)
}

func TestValidateAndGetDefaultValue(t *testing.T) {
	t.Parallel()

	actualValue, _ := validateAndGet("", []string{"default", "not-default"})
	expectedValue := "default"

	if expectedValue != actualValue {
		t.Errorf("default value not match default == %s", actualValue)
	}
}

func TestValidateAndGetValidValue(t *testing.T) {
	t.Parallel()

	actualValue, _ := validateAndGet("valid", []string{"not-valid", "valid"})
	expectedValue := "valid"

	if expectedValue != actualValue {
		t.Errorf("allowed value not match valid == %s", actualValue)
	}
}

func TestValidateAndGetInvalidValue(t *testing.T) {
	t.Parallel()

	_, err := validateAndGet("not-valid", []string{"valid1", "valid2"})

	if err == nil {
		t.Error("error doesn't occurred")
	}
}
