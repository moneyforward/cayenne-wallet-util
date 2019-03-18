package sign

import (
	"io/ioutil"
	"testing"

	wrsa "github.com/mf-financial/cayenne-wallet-util/rsa"
)

func BenchmarkSignRSAPSS(b *testing.B) {
	privateKeyFile, err := ioutil.ReadFile("testdata/private.pem")
	if err != nil {
		b.Fatal("failed to read private key file")
	}
	privateKey, err := wrsa.ReadRSAPrivateKeyFromBytes(privateKeyFile)
	if err != nil {
		b.Fatal("failed to create private key from bytes")
	}

	rsaSignature := &RSASignature{
		privatekey: privateKey,
	}

	testMessage := "ghiosehgiosh"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsaSignature.Sign(testMessage)
	}
}

func BenchmarkVerifyRSAPSS(b *testing.B) {
	privateKeyFile, err := ioutil.ReadFile("testdata/private.pem")
	if err != nil {
		b.Fatal("failed to read private key file")
	}
	privateKey, err := wrsa.ReadRSAPrivateKeyFromBytes(privateKeyFile)
	if err != nil {
		b.Fatal("failed to create private key from bytes")
	}

	rsaSignature := &RSASignature{
		privatekey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}

	testMessage := "ghiosehgiosh"
	sig, err := rsaSignature.Sign(testMessage)
	if err != nil {
		b.Fatal("failed to sign")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsaSignature.Verify(testMessage, sig)
	}
}
