/*
Copyright © 2024 Jean-Marc Meessen jean-marc@meessen-web.org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_performHonorContributorSelection_params(t *testing.T) {
	type args struct {
		dataDir           string
		outputFileName    string
		monthToSelectFrom string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"inexistent data directory",
			args{
				dataDir:           "inexistentDir",
				monthToSelectFrom: "2024-04",
			},
			true,
		},
		{
			"valid data directory and month",
			args{
				dataDir:           "../test-data",
				monthToSelectFrom: "2024-04",
			},
			false,
		},
		{
			"invalid month",
			args{
				monthToSelectFrom: "junkMonth",
				dataDir:           "../test-data",
			},
			true,
		},
		{
			"invalid header in input file",
			args{
				dataDir:           "../test-data",
				monthToSelectFrom: "2024-03",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := performHonorContributorSelection(tt.args.dataDir, tt.args.outputFileName, tt.args.monthToSelectFrom); (err != nil) != tt.wantErr {
				t.Errorf("performHonorContributorSelection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_honorCommand_paramCheck_noMonth(t *testing.T) {
	//Setup environment
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	var commandArguments []string
	commandArguments = append(commandArguments, "honor", "--data_dir=../test-data")
	rootCmd.SetArgs(commandArguments)

	// execute command
	error := rootCmd.Execute()

	// check results
	assert.ErrorContains(t, error, "\"month\" argument is missing.", "Call should have failed with expected error.")
}

func Test_honorCommand_paramCheck_invalidMonth(t *testing.T) {
	//Setup environment
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	var commandArguments []string
	commandArguments = append(commandArguments, "honor", "junkMonth", "--data_dir=../test-data")
	rootCmd.SetArgs(commandArguments)

	// execute command
	error := rootCmd.Execute()

	// check results
	assert.ErrorContains(t, error, "\"junkMonth\" is not a valid month.", "Call should have failed with expected error.")
}

// FIXME: do this in a temporary directory
func Test_honorCommand_integrationTest_verbose(t *testing.T) {
	//Setup environment
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	var commandArguments []string
	commandArguments = append(commandArguments, "honor", "2024-04", "--data_dir=../test-data", "--verbose")
	rootCmd.SetArgs(commandArguments)

	// execute command
	error := rootCmd.Execute()

	// check results
	assert.NoError(t, error, "Call should not have failed")

}

func Test_stringifySlice(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"happy case",
			args{s: []string{"aaa", "bbb", "ccc"}},
			"aaa bbb ccc",
		},
		{
			"Single item case",
			args{s: []string{"aaa"}},
			"aaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringifySlice(tt.args.s); got != tt.want {
				t.Errorf("stringifySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
