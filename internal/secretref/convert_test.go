package secretref

import (
	"reflect"
	"testing"
)

// Test ParseVaultType
func TestParseVaultType(t *testing.T) {
	// Test cases
	tests := []struct {
		name string
		str  string
		want VaultType
	}{
		{
			name: "AWS",
			str:  "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
			want: VaultTypeAws,
		},
		{
			name: "Unknown",
			str:  "Any other string",
			want: VaultTypeUnknown,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseVaultType(tt.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVaultType() = %v, want %v", got, tt.want)
			}
		})
	}
}
