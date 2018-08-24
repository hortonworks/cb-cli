package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteConfigToFileDirExists(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configwritetest")
	defer os.RemoveAll(tempDirName)
	os.MkdirAll(tempDirName+string(filepath.Separator)+Config_dir, 0700)

	WriteConfigToFile(tempDirName, "server", "user", "password", "output", "default", "oauth2", "organization")

	validateConfigContent(tempDirName, t)
}

func TestWriteConfigToFileDirNotExists(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configwritetest")
	defer os.RemoveAll(tempDirName)

	WriteConfigToFile(tempDirName, "server", "user", "password", "output", "default", "oauth2", "organization")

	validateConfigContent(tempDirName, t)
}

func validateConfigContent(tempDirName string, t *testing.T) {
	content, _ := ioutil.ReadFile(tempDirName + string(filepath.Separator) + Config_dir + string(filepath.Separator) + Config_file)

	expected := "default:\n  username: user\n  password: password\n  server: server\n  authType: oauth2\n  organization: organization\n  output: output\n"
	if string(content) != expected {
		t.Errorf("content not match %s == %s", expected, string(content))
	}
}

func TestReadConfig(t *testing.T) {
	t.Parallel()

	tempDirName, _ := ioutil.TempDir("", "configreadtest")
	defer os.RemoveAll(tempDirName)

	os.MkdirAll(tempDirName+string(filepath.Separator)+Config_dir, 0700)
	password := "§±!@#$%^&*()_+-=[]{};'\\:\"/.,?><`~"
	ioutil.WriteFile(tempDirName+string(filepath.Separator)+Config_dir+string(filepath.Separator)+Config_file, []byte("default:\n  username: user\n  password: "+password+"\n  server: server\n  organization: organization\n  output: output\n"), 0700)

	config, err := ReadConfig(tempDirName, "default")

	if err != nil {
		t.Errorf("unable to read file: %s", err.Error())
	}
	if config.Server != "server" {
		t.Errorf("server not match server == %s", config.Server)
	}
	if config.Username != "user" {
		t.Errorf("user not match user == %s", config.Username)
	}
	if config.Password != password {
		t.Errorf("password not match %s == %s", password, config.Password)
	}
	if config.Output != "output" {
		t.Errorf("output not match output == %s", config.Output)
	}
	if config.Organization != "organization" {
		t.Errorf("organization not match organization == %s", config.Organization)
	}
}
