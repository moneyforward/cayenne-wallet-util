package encryption_test

import (
	"testing"

	. "github.com/mf-financial/wallet-util/encryption"
)

func TestEncryption(t *testing.T) {
	size := 16
	key := "8#75F%R+&a5ZvM_<"
	iv := "@~wp-7hPs<WEx@R4"

	str := "test@gmail.com"

	NewCryptWithParam(size, key, iv)
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
