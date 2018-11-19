package utils

import (
	"errors"

	"github.com/hortonworks/cb-cli/cloudbreak/api/client/v1util"
	commonutils "github.com/hortonworks/cb-cli/dps-common/utils"
)

type utilClient interface {
	CheckClientVersion(params *v1util.CheckClientVersionParams) (*v1util.CheckClientVersionOK, error)
}

func CheckClientVersion(client utilClient, version string) {
	resp, err := client.CheckClientVersion(v1util.NewCheckClientVersionParams().WithVersion(version))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	valid := resp.Payload.VersionCheckOk
	message := resp.Payload.Message
	if valid == nil || !*valid {
		commonutils.LogErrorAndExit(errors.New(message))
	}
}
