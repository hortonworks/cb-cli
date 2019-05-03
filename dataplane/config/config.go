package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/dataplane/common"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server       string `json:"server" yaml:"server"`
	Workspace    string `json:"workspace,omitempty" yaml:"workspace,omitempty"`
	Output       string `json:"output,omitempty" yaml:"output,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty" yaml:"refreshtoken,omitempty"`
	ApiKeyID     string `json:"apikeyid" yaml:"apikeyid"`
	PrivateKey   string `json:"privatekey" yaml:"privatekey"`
}

type ConfigList map[string]Config

func (c Config) Yaml() string {
	j, _ := yaml.Marshal(c)
	return string(j)
}

func (c ConfigList) Json() string {
	j, _ := json.MarshalIndent(c, "", "  ")
	return string(j)
}

func (c ConfigList) Yaml() string {
	j, _ := yaml.Marshal(c)
	return string(j)
}

func CheckConfigAndCommandFlags(c *cli.Context) error {
	err := fl.CheckRequiredFlagsAndArguments(c)
	if err == nil {
		resp := configRead(c)
		validateContext(c, []fl.StringFlag{fl.FlServerOptional, fl.FlApiKeyIDOptional, fl.FlPrivateKeyOptional})
		setWorkspaceInContext(c)
		return resp
	}
	return err
}

func CheckConfigAndCommandFlagsWithoutWorkspace(c *cli.Context) error {
	err := fl.CheckRequiredFlagsAndArguments(c)
	if err == nil {
		resp := configRead(c)
		validateContext(c, []fl.StringFlag{fl.FlServerOptional, fl.FlApiKeyIDOptional, fl.FlPrivateKeyOptional})
		return resp
	}
	return err
}

func CheckConfigAndCommandFlagsDP(c *cli.Context) error {
	err := fl.CheckRequiredFlagsAndArguments(c)
	if err == nil {
		resp := configRead(c)
		validateContext(c, []fl.StringFlag{fl.FlServerOptional})
		return resp
	}
	return err
}

func validateContext(c *cli.Context, flagsTocheck []fl.StringFlag) {
	for _, f := range flagsTocheck {
		if len(c.String(f.Name)) == 0 {
			log.Error(fmt.Sprintf("configuration is not set, see: dp configure --help or provide the following flags: %v",
				[]string{"--" + f.Name}))
			os.Exit(1)
		}
	}
}

func setWorkspaceInContext(c *cli.Context) {
	set := func(name, value string) {
		if err := c.Set(name, value); err != nil {
			log.Debug(err)
		}
	}
	set(fl.FlWorkspaceOptional.Name, strconv.FormatInt(0, 10))
}

func configRead(c *cli.Context) error {
	args := c.Args()
	if args.Present() {
		name := args.First()
		if k := c.App.Command(name); k != nil {
			// this is a sub-command invocation
			return nil
		}
	}

	server := c.String(fl.FlServerOptional.Name)
	output := c.String(fl.FlOutputOptional.Name)
	profile := c.String(fl.FlProfileOptional.Name)
	apiKeyID := c.String(fl.FlApiKeyIDOptional.Name)
	privateKey := c.String(fl.FlPrivateKeyOptional.Name)

	if len(profile) == 0 {
		profile = "default"
	}

	config, err := ReadConfig(GetHomeDirectory(), profile)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	set := func(name, value string) {
		if err = c.Set(name, value); err != nil {
			log.Debug(err)
		}
	}

	if len(output) == 0 {
		set(fl.FlOutputOptional.Name, config.Output)
	}
	PrintConfig(*config)
	if len(server) == 0 {
		set(fl.FlServerOptional.Name, config.Server)
	}
	if len(apiKeyID) == 0 {
		set(fl.FlApiKeyIDOptional.Name, config.ApiKeyID)
	}
	if len(privateKey) == 0 {
		set(fl.FlPrivateKeyOptional.Name, config.PrivateKey)
	}

	return nil
}

func PrintConfig(cfg Config) {
	cfg.PrivateKey = "*"
	log.Infof("[CheckConfigAndCommandFlags] Config read from file, setting as global variable:\n%s", cfg.Yaml())
}

func GetHomeDirectory() string {
	homeDir, err := homedir.Dir()
	if err != nil || len(homeDir) == 0 {
		utils.LogErrorMessageAndExit("failed to determine the home directory")
	}
	return homeDir
}

func ReadConfig(baseDir string, profile string) (*Config, error) {
	configDir := baseDir + "/" + common.Config_dir
	configFile := configDir + "/" + common.Config_file

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return new(Config), nil
	}
	log.Infof("[ReadConfig] found config file: %s", configFile)

	var configList ConfigList
	err := readConfigToList(configFile, &configList)
	if err != nil {
		return nil, err
	}

	if config, valid := configList[profile]; valid {
		log.Infof("[ReadConfig] selected profile: %s", profile)
		return &config, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Invalid profile selected: %s", profile))
	}
}

func WriteConfigToFile(baseDir, server, output, profile, workspace, apiKeyID, privateKey string) error {
	configDir := baseDir + "/" + common.Config_dir
	configFile := configDir + "/" + common.Config_file
	if len(profile) == 0 {
		profile = "default"
	}
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		log.Infof("[writeConfigToFile] create dir: %s", configDir)
		err = os.MkdirAll(configDir, 0700)
		if err != nil {
			return err
		}
	} else {
		log.Infof("[writeConfigToFile] dir already exists: %s", configDir)
	}

	log.Infof("[writeConfigToFile] writing credentials to file: %s", configFile)

	var configList = make(ConfigList)
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		err = readConfigToList(configFile, &configList)
		if err != nil {
			return err
		}
	}

	configList[profile] = Config{
		Server:     server,
		Output:     output,
		ApiKeyID:   apiKeyID,
		PrivateKey: privateKey,
	}

	// in the case the token is empty and command is of type caas we don't want to overide workspace value
	if len(workspace) != 0 {
		var p = configList[profile]
		p.Workspace = workspace
		configList[profile] = p
	}

	err := ioutil.WriteFile(configFile, []byte(configList.Yaml()), 0600)
	if err != nil {
		return err
	}

	return nil
}

func readConfigToList(configPath string, configList *ConfigList) error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, configList)
	if err != nil {
		return err
	}
	return nil
}
