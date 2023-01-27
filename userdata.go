package main

import (
	"gopkg.in/yaml.v3"
)

type FileToWrite struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
	Append  bool   `yaml:"append"`
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

type RunCmd string

func (cyaml *RunCmd) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

type RunCmds struct {
	CommandsToRun []RunCmd `yaml:"runcmd,omitempty"`
}

func (cyaml *RunCmds) String() string {
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
	RunCmds        []RunCmd      `yaml:"runcmd,omitempty"`
}

func (cyaml *UserData) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return "#cloud-config\n" + string(result[:])
}
