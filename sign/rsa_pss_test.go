package sign

import (
	"crypto/rsa"
	"io/ioutil"
	"reflect"
	"testing"

	wrsa "github.com/mf-financial/cayenne-wallet-util/rsa"
)

func TestNewRSASignatureWithPublicKey(t *testing.T) {
	type args struct {
		publicKey *rsa.PublicKey
	}
	publicKeyFile, err := ioutil.ReadFile("./testdata/public.pem")
	if err != nil {
		t.Fatal("failed to read public key file")
	}
	publicKey, err := wrsa.ReadRSAPublicKeyFromBytes(publicKeyFile)
	if err != nil {
		t.Fatal("failed to create public key from bytes")
	}
	tests := []struct {
		name    string
		args    args
		want    Signature
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{publicKey},
			want: &RSASignature{
				publicKey: publicKey,
			},
		},
		{
			name:    "public keyがnil",
			args:    args{nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRSASignatureWithPublicKey(tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReceiveCustomerCryptoDepositDecryptedAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRSASignatureWithPublicKey() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestNewRSASignatureWithPrivatekey(t *testing.T) {
	type args struct {
		privateKey *rsa.PrivateKey
	}
	privateKeyFile, err := ioutil.ReadFile("./testdata/private.pem")
	if err != nil {
		t.Fatal("failed to read private key file")
	}
	privateKey, err := wrsa.ReadRSAPrivateKeyFromBytes(privateKeyFile)
	if err != nil {
		t.Fatal("failed to create private key from bytes")
	}
	tests := []struct {
		name    string
		args    args
		want    Signature
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{privateKey},
			want: &RSASignature{
				publicKey:  &privateKey.PublicKey,
				privatekey: privateKey,
			},
		},
		{
			name:    "privat keyがnil",
			args:    args{nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRSASignatureWithPrivatekey(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRSASignatureWithPrivatekey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRSASignatureWithPrivatekey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSASignature_Sign(t *testing.T) {
	type fields struct {
		publicKey  *rsa.PublicKey
		privatekey *rsa.PrivateKey
	}
	privateKeyFile, err := ioutil.ReadFile("./testdata/private.pem")
	if err != nil {
		t.Fatal("failed to read private key file")
	}
	privateKey, err := wrsa.ReadRSAPrivateKeyFromBytes(privateKeyFile)
	if err != nil {
		t.Fatal("failed to create private key from bytes")
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "正常系",
			fields: fields{privatekey: privateKey},
			args:   args{"aaaa"},
		},
		{
			name:    "private keyがnil",
			fields:  fields{privatekey: nil},
			args:    args{"bbbb"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RSASignature{
				publicKey:  tt.fields.publicKey,
				privatekey: tt.fields.privatekey,
			}
			_, err := s.Sign(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSASignature.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRSASignature_Verify(t *testing.T) {
	type fields struct {
		publicKey  *rsa.PublicKey
		privatekey *rsa.PrivateKey
	}
	privateKeyFile, err := ioutil.ReadFile("./testdata/private.pem")
	if err != nil {
		t.Fatal("failed to read private key file")
	}
	privateKey, err := wrsa.ReadRSAPrivateKeyFromBytes(privateKeyFile)
	if err != nil {
		t.Fatal("failed to create private key from bytes")
	}
	type args struct {
		message string
		sig     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "正常系",
			fields: fields{privatekey: privateKey, publicKey: &privateKey.PublicKey},
			args:   args{message: "aaaa"},
			want:   true,
		},
		{
			name:    "public keyがnil",
			fields:  fields{privatekey: privateKey, publicKey: nil},
			args:    args{message: "bbbb"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RSASignature{
				publicKey:  tt.fields.publicKey,
				privatekey: tt.fields.privatekey,
			}

			sig, err := s.Sign(tt.args.message)
			if err != nil {
				t.Error("failed to sign")
				return
			}
			tt.args.sig = sig
			got, err := s.Verify(tt.args.message, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSASignature.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RSASignature.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
