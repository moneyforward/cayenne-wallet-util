package rsa

import (
	"crypto/rsa"
	"io/ioutil"
	"testing"
)

func TestReadRSAPrivateKeyFromBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	privateFile, err := ioutil.ReadFile("./testdata/private.pem")
	if err != nil {
		t.Fatal("failed to read public key file")
	}
	textFile, err := ioutil.ReadFile("./testdata/test.txt")
	if err != nil {
		t.Fatal("failed to read text file")
	}
	tests := []struct {
		name    string
		args    args
		want    *rsa.PrivateKey
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{privateFile},
		},
		{
			name:    "ファイルがpublic keyではない",
			args:    args{textFile},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadRSAPrivateKeyFromBytes(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRSAPrivateKeyFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReadRSAPublicKeyFromBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	publicFile, err := ioutil.ReadFile("testdata/public.pem")
	if err != nil {
		t.Fatal("failed to read public key file")
	}
	textFile, err := ioutil.ReadFile("testdata/test.txt")
	if err != nil {
		t.Fatal("failed to read text file")
	}
	tests := []struct {
		name    string
		args    args
		want    *rsa.PublicKey
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{publicFile},
		},
		{
			name:    "ファイルがpublic keyではない",
			args:    args{textFile},
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadRSAPublicKeyFromBytes(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRSAPublicKeyFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
