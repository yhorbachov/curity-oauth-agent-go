package utils_test

import (
	"encoding/hex"
	"testing"

	"github.com/yhorbachov/curity-oauth-agent-g/pkg/utils"
)

func TestEcrypt(t *testing.T) {
	type args struct {
		key   []byte
		value []byte
	}

	validKey, _ := hex.DecodeString("4e4636356d65563e4c73233847503e3b21436e6f7629724950526f4b5e2e4e50")

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Invalid key",
			args: args{
				key:   []byte(""),
				value: []byte(""),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid key",
			args: args{
				key:   validKey,
				value: []byte(""),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.Encrypt(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		key   []byte
		value string
	}

	validKey, _ := hex.DecodeString("4e4636356d65563e4c73233847503e3b21436e6f7629724950526f4b5e2e4e50")
	validValue, _ := utils.Encrypt(validKey, []byte("Secret sauce"))

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Invalid key",
			args: args{
				key:   []byte(""),
				value: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid value",
			args: args{
				key:   validKey,
				value: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid key and value",
			args: args{
				key:   validKey,
				value: validValue,
			},
			want:    "Secret sauce",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.Decrypt(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
