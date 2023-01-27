package main

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
				FileToWrite{
					Path:    "/home/u/.ssh/authorized_keys",
					Content: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar",
					Append:  true,
				},
			}},
			`write_files:
    - path: /home/u/.ssh/authorized_keys
      content: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar
      append: true
`,
			false,
		},
		{
			"some_multilin",
			WriteFiles{[]FileToWrite{
				FileToWrite{
					Path: "/some-multiline",
					Content: `foo
bar
foobar`,
					Append: false,
				},
				FileToWrite{
					Path: "/some-multiline2",
					Content: `foo
bar
foobar`,
					Append: false,
				},
			}},
			`write_files:
    - path: /some-multiline
      content: |-
        foo
        bar
        foobar
      append: false
    - path: /some-multiline2
      content: |-
        foo
        bar
        foobar
      append: false
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("running %s\n", testCase.name)
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
content: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2LVzkp5iPHX8x== foo@bar
append: true
`,
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("running %s\n", testCase.name)
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
			t.Logf("running %s\n", testCase.name)
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
			t.Logf("running %s\n", testCase.name)
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
			t.Logf("running %s\n", testCase.name)
			actualResult := testCase.testData.String()
			if (actualResult != testCase.expResult) != testCase.wantErr {
				t.Logf("%+v\n", string(actualResult))
				t.Errorf("failed, expected %s", testCase.expResult)
			}
		})
	}
}
