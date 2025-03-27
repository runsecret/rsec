package secrets

import (
	"reflect"
	"testing"
)

// Test ParseRefType
func TestGetRefType(t *testing.T) {
	// Test cases
	tests := []struct {
		name string
		str  string
		want SecretRefType
	}{
		{
			name: "AWS",
			str:  "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
			want: SecretRefTypeAwsArn,
		},
		{
			name: "AWS",
			str:  "aws://us-east-1/000000000000/MyTestSecret",
			want: SecretRefTypeAwsRef,
		},
		{
			name: "Unknown",
			str:  "Any other string",
			want: SecretRefTypeUnknown,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRefType(tt.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVaultType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test GetVaultReference
func TestGetVaultAddress(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		str       string
		wantVault VaultType
		wantRef   string
	}{
		{
			name:      "AWS ARN",
			str:       "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
			wantVault: VaultTypeUnknown,
			wantRef:   "Invalid secret reference",
		},
		{
			name:      "AWS Ref",
			str:       "aws://us-east-1/000000000000/MyTestSecret",
			wantVault: VaultTypeAws,
			wantRef:   "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
		},
		{
			name:      "Unknown",
			str:       "Any other string",
			wantVault: VaultTypeUnknown,
			wantRef:   "Invalid secret reference",
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotRef := GetVaultAddress(tt.str)
			if !reflect.DeepEqual(gotType, tt.wantVault) {
				t.Errorf("GetVaultReference() = %v, want %v", gotType, tt.wantVault)
			}
			if !reflect.DeepEqual(gotRef, tt.wantRef) {
				t.Errorf("GetVaultReference() = %v, want %v", gotRef, tt.wantRef)
			}
		})
	}
}
