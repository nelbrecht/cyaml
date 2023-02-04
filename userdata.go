package cyaml

import (
	"encoding/base64"
	"log"
	"os"
	"time"

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

type User struct {
	Default           string   `yaml:"default,omitempty"`
	Name              string   `yaml:"name,omitempty"`
	Expiredate        string   `yaml:"expiredate,omitempty"`
	Gecos             string   `yaml:"gecos,omitempty"`
	Groups            string   `yaml:"groups,omitempty"`
	Homedir           string   `yaml:"homedir,omitempty"`
	Inactive          int      `yaml:"inactive,omitempty"`
	LockPasswd        bool     `yaml:"lock_passwd,omitempty"`
	NoCreateHome      bool     `yaml:"no_create_home,omitempty"`
	NoLogInit         bool     `yaml:"no_log_init,omitempty"`
	NoUserGroup       bool     `yaml:"no_user_group,omitempty"`
	Passwd            string   `yaml:"passwd,omitempty"`
	PrimaryGroup      string   `yaml:"primary_group,omitempty"`
	SelinuxUser       string   `yaml:"selinux_user,omitempty"`
	Shell             string   `yaml:"shell,omitempty"`
	SshAuthorizedKeys string   `yaml:"ssh_authorized_keys,omitempty"`
	SshImportId       []string `yaml:"ssh_import_id,omitempty"`
	SshRedirectUser   bool     `yaml:"ssh_redirect_user,omitempty"`
	Sudo              string   `yaml:"sudo,omitempty"`
	System            bool     `yaml:"system,omitempty"`
}

func (cyaml *User) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return string(result)
}

func (cyaml *User) SetExpireDate(t time.Time) {
	cyaml.Expiredate = t.Format("2006-01-02")
}

type Users struct {
	UserToAdd []User `yaml:"users,omitempty"`
}

func (cyaml *Users) String() string {
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
	Users          []User        `yaml:"users,omitempty"`
}

func (cyaml *UserData) String() string {
	result, err := yaml.Marshal(cyaml)
	if err != nil {
		return ""
	}

	return "#cloud-config\n" + string(result[:])
}
