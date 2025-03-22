package envvars

// Test ReplaceEnvVarSecrets
// func TestReplaceEnvVarSecrets(t *testing.T) {
// 	// Test cases
// 	tests := []struct {
// 		name       string
// 		rawEnv     []string
// 		wantEnv    []string
// 		wantRedact []string
// 		wantErr    bool
// 	}{
// 		{
// 			name: "AWS",
// 			rawEnv: []string{
// 				"FOO=arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret",
// 			},
// 			wantEnv: []string{
// 				"FOO=My secret",
// 			},
// 			wantRedact: []string{
// 				"My secret",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	// Run tests
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotEnv, gotRedact, err := ReplaceEnvVarSecrets(tt.rawEnv)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ReplaceEnvVarSecrets() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotEnv, tt.wantEnv) {
// 				t.Errorf("ReplaceEnvVarSecrets() = %v, want %v", gotEnv, tt.wantEnv)
// 			}
// 			if !reflect.DeepEqual(gotRedact, tt.wantRedact) {
// 				t.Errorf("ReplaceEnvVarSecrets() = %v, want %v", gotRedact, tt.wantRedact)
// 			}
// 		})
// 	}
// }
