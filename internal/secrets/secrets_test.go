package secrets

import (
	"reflect"
	"testing"
)

// Test GetIdentifierType
func TestGetIdentifierType(t *testing.T) {
	// Test cases
	tests := []struct {
		name string
		str  string
		want SecretIdentifierType
	}{
		{
			name: "AWS",
			str:  "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
			want: SecretIdentifierTypeAwsArn,
		},
		{
			name: "AWS",
			str:  "rsec://000000000000.sm.aws/MyTestSecret?region=us-east-1",
			want: SecretIdentifierTypeRef,
		},
		{
			name: "Unknown",
			str:  "Any other string",
			want: SecretIdentifierTypeUnknown,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIdentifierType(tt.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVaultType() = %v, want %v", got, tt.want)
			}
		})
	}
}
