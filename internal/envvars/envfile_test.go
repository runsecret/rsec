package envvars

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test Read
func TestReadEnvFile(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		path    string
		want    []string
		wantErr bool
	}{
		{
			name: "valid env file",
			path: "testdata/envfile.env",
			want: []string{
				"FOO=bar",
				"BAZ=qux",
			},
			wantErr: false,
		},
		{
			name:    "invalid env file",
			path:    "testdata/invalid.env",
			want:    []string{},
			wantErr: true,
		},
		{
			name:    "non-existent env file",
			path:    "testdata/non-existent.env",
			want:    nil,
			wantErr: true,
		},
		// Test case for file with comments
		{
			name: "env file with comments",
			path: "testdata/envfile_with_comments.env",
			want: []string{
				"FOO=bar",
				"BAZ=qux",
			},
			wantErr: false,
		},
		// Test case for file with empty lines
		{
			name: "env file with empty lines",
			path: "testdata/envfile_with_empty_lines.env",
			want: []string{
				"FOO=bar",
				"BAZ=qux",
			},
			wantErr: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readEnvFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadEnvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadEnvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateEnvFile(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		envVars []string
		wantErr bool
	}{
		{
			name: "valid env file",
			envVars: []string{
				"FOO=bar",
				"BAZ=qux",
			},
			wantErr: false,
		},
		{
			name: "invalid env file, missing value",
			envVars: []string{
				"FOO=bar",
				"BAZ=",
			},
			wantErr: true,
		},
		{
			name: "invalid env file, missing key",
			envVars: []string{
				"FOO=bar",
				"=qux",
			},
			wantErr: true,
		},
		{
			name: "invalid env file, missing key and value",
			envVars: []string{
				"FOO=bar",
				"BAZ",
			},
			wantErr: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateEnvFile(tt.envVars)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
