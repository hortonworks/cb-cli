package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hortonworks/cb-cli/dataplane/api/model"

	"github.com/hortonworks/cb-cli/dataplane/api/client/v4utils"
	"github.com/hortonworks/dp-cli-common/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
)

type utilClient interface {
	CheckClientVersionV4(params *v4utils.CheckClientVersionV4Params) (*v4utils.CheckClientVersionV4OK, error)
}

func CheckClientVersion(client utilClient, version string) {
	resp, err := client.CheckClientVersionV4(v4utils.NewCheckClientVersionV4Params().WithVersion(&version))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	valid := resp.Payload.VersionCheckOk
	message := resp.Payload.Message
	if !valid {
		commonutils.LogErrorAndExit(errors.New(message))
	}
}

func SafeClusterDescriptionConvert(s *model.StackV4Response) string {
	if s.Cluster != nil {
		return s.Cluster.Description
	}
	return ""
}

func SafeClusterViewStatusConvert(s *model.StackViewV4Response) string {
	if s.Cluster != nil {
		return s.Cluster.Status
	}
	return ""
}

func SafeClusterStatusConvert(s *model.StackV4Response) string {
	if s.Cluster != nil {
		return s.Cluster.Status
	}
	return ""
}

func CheckServerAddress(address string) {
	if len(address) == 0 {
		utils.LogErrorMessageAndExit("Server address is empty. Please set it with the `dp configure` command or the `--server` option.")
	}
}

func ConvertToInt32Ptr(source string) *int32 {
	if strings.TrimSpace(source) == "" {
		return nil
	}
	temporaryInt, err := strconv.ParseInt(source, 10, 32)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	result := int32(temporaryInt)
	return &result
}

func ConvertToFloat64Ptr(source string) *float64 {
	if strings.TrimSpace(source) == "" {
		return nil
	}
	temporaryInt, err := strconv.ParseFloat(source, 64)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	result := float64(temporaryInt)
	return &result
}
