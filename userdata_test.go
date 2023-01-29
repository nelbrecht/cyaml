package cyaml

import (
	"testing"
)

func TestWriteFiles(t *testing.T) {
	tests := []struct {
		name      string
		testData  WriteFiles
		expResult string
		wantErr   bool
	}{
		{
			"authorized_keys",
			WriteFiles{[]FileToWrite{
				{
					Path:    "/home/u/.ssh/authorized_keys",
					Content: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar",
					Append:  true,
				},
			}},
			`write_files:
    - path: /home/u/.ssh/authorized_keys
      append: true
      content: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar
`,
			false,
		},
		{
			"some_multilin",
			WriteFiles{[]FileToWrite{
				{
					Path: "/some-multiline",
					Content: `foo
bar
foobar`,
					Append: false,
				},
				{
					Path: "/some-multiline2",
					Content: `foo
bar
foobar`,
					Append: false,
				},
			}},
			`write_files:
    - path: /some-multiline
      append: false
      content: |-
        foo
        bar
        foobar
    - path: /some-multiline2
      append: false
      content: |-
        foo
        bar
        foobar
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}

func TestFileToWrite(t *testing.T) {
	tests := []struct {
		name      string
		testData  FileToWrite
		expResult string
		wantErr   bool
	}{
		{
			"one_file",
			FileToWrite{
				Path:    "/home/u/.ssh/authorized_keys",
				Content: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar",
				Append:  true,
			},
			`path: /home/u/.ssh/authorized_keys
append: true
content: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar
`,
			false,
		},
		{
			"full_file",
			FileToWrite{
				Append:      true,
				Content:     "dGVzdAo=",
				Defer:       true,
				Encoding:    "b64",
				Owner:       "syslog:staff",
				Path:        "/home/u/full_file",
				Permissions: "4755",
			},
			`path: /home/u/full_file
append: true
content: dGVzdAo=
defer: true
encoding: b64
owner: syslog:staff
permissions: "4755"
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}
func TestFileContentToWrite(t *testing.T) {
	tests := []struct {
		name          string
		testData      FileToWrite
		localFileName string
		expResult     string
		wantErr       bool
	}{
		{
			"read_file",
			FileToWrite{
				Path:   "/home/u/testfile",
				Append: true,
			},
			"testfile",
			`path: /home/u/testfile
append: true
content: YXNkZgpmb28gYmFy
encoding: b64
`,
			false,
		},
		{
			"full_file",
			FileToWrite{
				Path:        "/home/u/full_file",
				Append:      true,
				Defer:       true,
				Owner:       "syslog:staff",
				Permissions: "4755",
			},
			"testfile",
			`path: /home/u/full_file
append: true
content: YXNkZgpmb28gYmFy
defer: true
encoding: b64
owner: syslog:staff
permissions: "4755"
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.testData.AddLocalFile(testCase.localFileName)
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name      string
		testData  RunCmd
		expResult string
		wantErr   bool
	}{
		{
			"cmd",
			RunCmd("cmd"),
			"cmd\n",
			false,
		},
		{
			"multicmd",
			RunCmd(`date
id`),
			"|-\n    date\n    id\n",
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}

func TestRunCmds(t *testing.T) {
	tests := []struct {
		name      string
		testData  RunCmds
		expResult string
		wantErr   bool
	}{
		{
			"single commands",
			RunCmds{[]RunCmd{
				"foo", "bar", "foobar",
			}},
			"runcmd:\n    - foo\n    - bar\n    - foobar\n",
			false,
		},
		{
			"multiple commands",
			RunCmds{[]RunCmd{
				"foo", "bar\nbarfoo\nbarbar", "foobar",
			}},
			"runcmd:\n    - foo\n    - |-\n      bar\n      barfoo\n      barbar\n    - foobar\n",
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}

func TestUserData(t *testing.T) {
	tests := []struct {
		name      string
		testData  UserData
		expResult string
		wantErr   bool
	}{
		{
			"empty",
			UserData{},
			"#cloud-config\npackage_update: false\npackage_upgrade: false\n",
			false,
		},
		{
			"sparse",
			UserData{
				RunCmds: []RunCmd{
					"foo", "bar\nbarfoo\nbarbar", "foobar",
				}},
			`#cloud-config
package_update: false
package_upgrade: false
runcmd:
    - foo
    - |-
      bar
      barfoo
      barbar
    - foobar
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}
