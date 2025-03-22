package envfile

import (
	"reflect"
	"testing"
)

// Test Read
func TestRead(t *testing.T) {
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
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.path)
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
