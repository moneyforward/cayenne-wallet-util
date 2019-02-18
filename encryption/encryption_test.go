package encryption_test

import (
	"testing"

	. "github.com/mf-financial/cayenne-wallet-util/encryption"
)

func TestEncryptionNewCrypto(t *testing.T) {
	var tests = []struct {
		key         string
		iv          string
		isErr       bool
		description string
	}{
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", false, "key is 32 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fj", "@~Pp-6sC3<M8x@RA", false, "key is 24 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", false, "key is 16 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3V", "@~Pp-6sC3<M8x@RA", true, "key is 33 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug", "@~Pp-6sC3<M8x@RA", true, "key is 31 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjG", "@~Pp-6sC3<M8x@RA", true, "key is 25 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^f", "@~Pp-6sC3<M8x@RA", true, "key is 23 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3G", "@~Pp-6sC3<M8x@RAw", true, "key is 17 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrub", "@~Pp-6sC3<M8x@RAw", true, "key is 15 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RvW", true, "key is 32 bytes, iv is 17 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@R", true, "key is 32 bytes, iv is 15 bytes"},
	}

	for _, val := range tests {
		_, err := NewCrypt(val.key, val.iv)
		if err != nil && !val.isErr {
			t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n Unexpectedly error occorred. %v", val.description, val.key, val.iv, err)
		}
		if err == nil && val.isErr {
			t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n Error is expected. However nothing happened.", val.description, val.key, val.iv)
		}
	}

}

func TestEncryptDecrypt(t *testing.T) {
	var tests = []struct {
		key         string
		iv          string
		source      string
		expected    string
		isInvalid   bool
		description string
	}{
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", "test@gmail.com", "LsWFRCOS66xULAxQ1Sf4Xw==", false, "key is 32 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", "testAABBDD$*%&@gmail.com", "TKgJT5muIed5gml1PSK9/OmwjeonAZz7zesxCbl+hKU=", false, "key is 32 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fj", "@~Pp-6sC3<M8x@RA", "testAABBDD$*%&@gmail.com", "TKgJT5muIed5gml1PSK9/OmwjeonAZz7zesxCbl+hKU=", true, "key is 24 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3PBc1h^fj", "@~Pp-6sC3<M8x@RA", "testAABBDD$*%&@gmail.com", "19XQ7PIOnIfRUu6Mvo6Ji33yb1nQVkLnW5tZp3LNbG4=", false, "key is 24 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", "testAABBDD$*%&@gmail.com", "TKgJT5muIed5gml1PSK9/OmwjeonAZz7zesxCbl+hKU=", true, "key is 16 bytes, iv is 16 bytes"},
		{"PBc1h^fjKd3Mrug3", "@~Pp-6sC3<M8x@RA", "testAABBDD$*%&@gmail.com", "Lepfg7eDOGqfXbPPt8i3TKDiYlkgEJlh9ASBR5b26D8=", false, "key is 16 bytes, iv is 16 bytes"},
	}

	for _, val := range tests {
		crypt, err := NewCrypt(val.key, val.iv)
		if err != nil {
			t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n failed to call NewCryptWithParam(). %v", val.description, val.key, val.iv, err)
		}
		// Encrypt
		result1 := crypt.EncryptBase64(val.source)
		if !val.isInvalid && result1 != val.expected {
			t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n failed to call EncryptBase64(): got: %s, want: %s", val.description, val.key, val.iv, result1, val.expected)
		}
		if val.isInvalid && result1 == val.expected {
			t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n result should be different in EncryptBase64()", val.description, val.key, val.iv)
		}

		// Decrypt
		if !val.isInvalid {
			result2, _ := crypt.DecryptBase64(result1)
			if !val.isInvalid && result2 != val.source {
				t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n failed to call DecryptBase64(): got: %s, want: %s", val.description, val.key, val.iv, result2, val.source)
			}
			if val.isInvalid && result2 == val.source {
				t.Errorf("[Test:%s]\n[key]%s, [iv]%s\n result should be different in DecryptBase64()", val.description, val.key, val.iv)
			}
		}
	}
}
