package cli

import (
	"errors"

	"github.com/hortonworks/hdc-cli/cli/cloud"
	_ "github.com/hortonworks/hdc-cli/cli/cloud/aws"

	"fmt"
	"regexp"
	"strconv"

	swagerrors "github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	ss "github.com/hortonworks/hdc-cli/client_cloudbreak/smartsensesubscriptions"
)

func (s *ClusterSkeleton) Validate(cbClient *Cloudbreak) error {
	var res []error
	provider := cloud.GetProvider()

	if err := validate.RequiredString("ClusterName", "body", string(s.ClusterName)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("HDPVersion", "body", string(s.HDPVersion)); err != nil {
		res = append(res, err)
	}
	if err := validateHDPVersion(string(s.HDPVersion)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterType", "body", string(s.ClusterType)); err != nil {
		res = append(res, err)
	}
	if s.Master.InstanceCount != 0 && s.Master.InstanceCount != 1 {
		res = append(res, swagerrors.New(1, "The instance count must be 1 for the 'master' group"))
	}
	if err := validate.RequiredNumber("InstanceCount", "worker", float64(s.Worker.InstanceCount)); err != nil {
		res = append(res, err)
	} else if s.Worker.InstanceCount < 1 {
		res = append(res, swagerrors.New(1, "The instance count has to be greater than 0"))
	}
	if s.Compute.InstanceCount >= 0 {
		res = append(res)
	} else {
		res = append(res, swagerrors.New(1, "The instance count has to be not less than 0"))
	}
	if s.Compute.SpotPrice != "" {
		if f, err := strconv.ParseFloat(s.Compute.SpotPrice, 64); err != nil || f <= 0 {
			res = append(res, swagerrors.New(1, "SpotPrice must be numeric and greater than 0"))
		}
	}
	if err := validate.RequiredString("SSHKeyName", "body", string(s.SSHKeyName)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("RemoteAccess", "body", string(s.RemoteAccess)); err != nil {
		res = append(res, err)
	}
	if err := validate.Required("WebAccess", "body", s.WebAccess); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterAndAmbariUser", "body", string(s.ClusterAndAmbariUser)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterAndAmbariPassword", "body", string(s.ClusterAndAmbariPassword)); err != nil {
		res = append(res, err)
	}
	if s.Network != nil {
		if err := provider.ValidateNetwork(s.Network); err != nil {
			for _, e := range err {
				res = append(res, e)
			}
		}
	}
	if s.HiveMetastore != nil {
		if err := s.HiveMetastore.Validate(); err != nil {
			for _, e := range err {
				res = append(res, e)
			}
		}
	}
	if s.AmbariDatabase != nil {
		if err := s.AmbariDatabase.Validate(); err != nil {
			for _, e := range err {
				res = append(res, e)
			}
		}
	}
	if len(s.Master.Recipes) != 0 {
		if e := validateRecipes(s.Master.Recipes); len(e) != 0 {
			res = append(res, e...)
		}
	}

	if len(s.Worker.Recipes) != 0 {
		if e := validateRecipes(s.Worker.Recipes); len(e) != 0 {
			res = append(res, e...)
		}
	}

	if len(s.Compute.Recipes) != 0 {
		if e := validateRecipes(s.Compute.Recipes); len(e) != 0 {
			res = append(res, e...)
		}
	}

	if err := validateMasterRecoveryMode("master", s.Master.RecoveryMode); err != nil {
		res = append(res, err)
	}
	if err := validateEncryption(s.Master.Encrypted, s.Master.VolumeType); err != nil {
		res = append(res, err)
	}

	if err := validateRecoveryMode("worker", s.Worker.RecoveryMode); err != nil {
		res = append(res, err)
	}
	if err := validateEncryption(s.Worker.Encrypted, s.Worker.VolumeType); err != nil {
		res = append(res, err)
	}

	if err := validateSpotRecoveryMode("compute", s.Compute.RecoveryMode, s.Compute.SpotPrice); err != nil {
		res = append(res, err)
	}
	if err := validateEncryption(s.Compute.Encrypted, s.Compute.VolumeType); err != nil {
		res = append(res, err)
	}

	if err := provider.ValidateTags(s.Tags); len(err) > 0 {
		res = append(res, err...)
	}

	if err := s.Autoscaling.Validate(); err != nil {
		res = append(res, err...)
	}

	if err := validateFlexSubscription(cbClient, s); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return swagerrors.CompositeValidationError(res...)
	}
	return nil
}

func (a *AutoscalingSkeletonBase) Validate() []error {
	var res []error = nil

	if a != nil {
		if a.Configuration != nil {
			conf := a.Configuration
			if conf.ClusterMinSize > conf.ClusterMaxSize {
				res = append(res, errors.New("The minimum cluster size cannot be greater than the maximum cluster size"))
			}
			if conf.ClusterMinSize < 1 {
				res = append(res, errors.New("The minimum cluster size cannot be less than 1"))
			}
			if conf.ClusterMaxSize == 0 {
				res = append(res, errors.New("The maximum cluster size cannot be 0"))
			}
			if conf.ClusterMaxSize > 1000 {
				res = append(res, errors.New("The maximum cluster size cannot be greater than 1000"))
			}
			if conf.CooldownTime < 1 {
				res = append(res, errors.New("The cooldown time cannot be less than 1 minute"))
			}
		}

		policies := a.Policies
		if policies != nil {
			pattern := "^[a-zA-Z0-9]+$"
			for _, p := range policies {
				name := p.PolicyName

				if len(name) < 5 || len(name) > 20 {
					res = append(res, errors.New(fmt.Sprintf("The policy's name's (%s) length must be between 5 and 20 characters", name)))
				}

				if match, _ := regexp.Match(pattern, []byte(name)); !match {
					res = append(res, errors.New(fmt.Sprintf("The policy's name (%s) contains invalid characters. "+
						"Allowed characters are letters and numbers representable in UTF-8", name)))
				}

				if p.ScalingAdjustment == 0 {
					res = append(res, errors.New("The scaling adjustment cannot be 0"))
				}

				if p.ScalingDefinition == nil {
					res = append(res, errors.New("The scaling definition must be provided"))
				}

				if len(p.Operator) == 0 {
					res = append(res, errors.New("The operator must be provided"))
				}

				if p.Period < 1 {
					res = append(res, errors.New("The period must be at least 1 minute"))
				}

				if p.NodeType != WORKER && p.NodeType != COMPUTE {
					res = append(res, errors.New("The nodeType must be one of [worker, compute]"))
				}

			}
		}
	}

	return res
}

func (h *HiveMetastore) Validate() []error {
	var res []error = nil

	if h.isNew() {
		if err := validate.RequiredString("Name", "hivemetastore", h.Name); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("DatabaseType", "hivemetastore", h.DatabaseType); err != nil {
			res = append(res, err)
		} else if h.DatabaseType != POSTGRES {
			res = append(res, errors.New("Invalid database type. Accepted value is: POSTGRES"))
		}
		if err := validate.RequiredString("Password", "hivemetastore", h.Password); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("Username", "hivemetastore", h.Username); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("URL", "hivemetastore", h.URL); err != nil {
			res = append(res, err)
		}
	}

	return res
}

func (h *RangerMetastore) Validate() []error {
	var res []error = nil
	if err := validate.RequiredString("Name", "RangerMetastore", h.Name); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Password", "RangerMetastore", h.Password); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Username", "RangerMetastore", h.Username); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("URL", "RangerMetastore", h.URL); err != nil {
		res = append(res, err)
	}

	return res
}

func (r *Recipe) Validate() []error {
	var res []error = nil

	if err := validate.RequiredString("URI", "recipe", r.URI); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Phase", "recipe", r.Phase); err != nil {
		res = append(res, err)
	}

	if r.Phase != PRE && r.Phase != POST {
		res = append(res, errors.New(fmt.Sprintf("Valid recipe phases: %s, %s", PRE, POST)))
	}

	return res
}

func validateRecipes(recipes []Recipe) []error {
	var res []error = make([]error, 0)
	for _, recipe := range recipes {
		for _, e := range recipe.Validate() {
			res = append(res, e)
		}
	}
	return res
}

func (h *HiveMetastore) isNew() bool {
	return len(h.DatabaseType) > 0 || len(h.Username) > 0 || len(h.Password) > 0 || len(h.URL) > 0
}

func validateHDPVersion(version string) error {
	if hdp, err := strconv.ParseFloat(version, 10); err != nil || !isVersionSupported(hdp) {
		return errors.New(fmt.Sprintf("Invalid HDP version. Accepted value(s): %v", SUPPORTED_HDP_VERSIONS))
	}
	return nil
}

func isVersionSupported(version float64) bool {
	for _, v := range SUPPORTED_HDP_VERSIONS {
		if v == version {
			return true
		}
	}
	return false
}

func validateMasterRecoveryMode(hostGroup string, recoveryMode string) error {
	if recoveryMode != "" && recoveryMode != "MANUAL" {
		return errors.New(fmt.Sprintf("Invalid recoveryMode [%s] on hostgroup [%s], supported revorery mode on master hostgroups is MANUAL only",
			recoveryMode, hostGroup))
	}
	return nil
}

func validateEncryption(encryption *bool, volumetype string) error {
	if encryption != nil && *encryption == true && volumetype == "ephemeral" {
		return errors.New(fmt.Sprintf("Invalid encryption because [%s] volume does not support encryption", volumetype))
	}
	return nil
}

func validateRecoveryMode(hostGroup string, recoveryMode string) error {
	if recoveryMode != "" && recoveryMode != "MANUAL" && recoveryMode != "AUTO" {
		return errors.New(fmt.Sprintf("Invalid recoveryMode [%s] on Hostgroup [%s], supported revorery modes are MANUAL or AUTO", recoveryMode, hostGroup))
	}
	return nil
}

func validateSpotRecoveryMode(hostGroup string, recoveryMode string, spotPrice string) error {
	if spotPrice != "" && recoveryMode != "" && recoveryMode != "MANUAL" {
		return errors.New(
			fmt.Sprintf("Invalid recoveryMode [%s] on Hostgroup [%s] with spotprice [%s], supported revorery mode for nodes with spotprice is MANUAL only",
				recoveryMode, hostGroup, spotPrice))
	}
	return validateRecoveryMode(hostGroup, recoveryMode)
}

func validateFlexSubscription(cbClient *Cloudbreak, skeleton *ClusterSkeleton) error {
	flexSubscription := skeleton.FlexSubscription
	if flexSubscription != nil && len(flexSubscription.Name) > 0 {
		ss, err := cbClient.Cloudbreak.Smartsensesubscriptions.GetSmartSenseSubscription(&ss.GetSmartSenseSubscriptionParams{})
		if err != nil {
			return err
		} else if *ss.Payload.AutoGenerated {
			return errors.New("Flex subscription could not be configured, because your SmartSense subscription is generated.")
		}
	}

	return nil
}

func (h *AmbariDatabase) Validate() []error {
	var res []error = nil
	if err := validate.RequiredString("DatabaseName", "AmbariDatabase", h.DatabaseName); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Password", "AmbariDatabase", h.Password); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Username", "AmbariDatabase", h.Username); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("Host", "AmbariDatabase", h.Host); err != nil {
		res = append(res, err)
	}
	if h.Port == 0 {
		res = append(res, errors.New("Ambari database port cannot be 0"))
	}
	if err := validate.RequiredString("DatabaseType", "AmbariDatabase", h.DatabaseType); err != nil {
		res = append(res, err)
	} else {
		for _, db := range DB_TYPES {
			if db == h.DatabaseType {
				return res
			}
		}
		res = append(res, errors.New(fmt.Sprintf("Invalid database type. Accepted values are: %s", DB_TYPES)))
	}

	return res
}
