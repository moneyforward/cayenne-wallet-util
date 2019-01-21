package encryption_test

import (
	"testing"

	. "github.com/mf-financial/cayenne-wallet-util/encryption"
)

func TestEncryption(t *testing.T) {
	key := "PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3"
	iv := "@~Pp-6sC3<M8x@RA"

	str := "test@gmail.com"

	NewCryptWithParam(key, iv)
	crypt := GetCrypt()

	result1 := crypt.EncryptBase64(str)
	if result1 != "SpJqzcL176g9aBq88pkKQw==" {
		t.Errorf("EncryptBase64() result: %s", result1)
	}

	result2, _ := crypt.DecryptBase64(result1)
	if result2 != str {
		t.Errorf("DecryptBase64() result: %s", result2)
	}
}
