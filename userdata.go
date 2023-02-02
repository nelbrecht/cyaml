package cyaml

import (
	"encoding/base64"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type FileToWrite struct {
	Path        string `yaml:"path"`
	Append      bool   `yaml:"append"`
	Content     string `yaml:"content"`
	Defer       bool   `yaml:"defer,omitempty"`
	Encoding    string `yaml:"encoding,omitempty"`
	Owner       string `yaml:"owner,omitempty"`
	Permissions string `yaml:"permissions,omitempty"`
}

func (cyaml *FileToWrite) AddLocalFile(path string) {
	if path != "" {
		fileContent, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("file %s could not be read\n%s", path, err.Error())
			fileContent = []byte("")
		}
		cyaml.Content = base64.StdEncoding.EncodeToString(fileContent)
		cyaml.Encoding = "b64"
	}
}

func (cyaml *FileToWrite) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type WriteFiles struct {
	FilesToWrite []FileToWrite `yaml:"write_files"`
}

func (cyaml *WriteFiles) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type CliCmd string

func (cyaml *CliCmd) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type RunCmds struct {
	CommandsToRun []CliCmd `yaml:"runcmd,omitempty"`
}

func (cyaml *RunCmds) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type BootCmds struct {
	CommandsToRun []CliCmd `yaml:"bootcmd,omitempty"`
}

func (cyaml *BootCmds) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type UserData struct {
	PackageUpdate  bool          `yaml:"package_update"`
	PackageUpgrade bool          `yaml:"package_upgrade"`
	WriteFiles     []FileToWrite `yaml:"write_files,omitempty"`
	RunCmds        []CliCmd      `yaml:"runcmd,omitempty"`
	BootCmds       []CliCmd      `yaml:"bootcmd,omitempty"`
}

func (cyaml *UserData) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return "#cloud-config\n" + string(result[:])
}
