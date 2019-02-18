package encryption

import (
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	//Result by `go test -bench=. -v ./...`
	//BenchmarkEncrypt/Encrypto_with_base64_Bitcoin_Address-8         	 3000000	       534 ns/op
	//BenchmarkEncrypt/Encrypto_with_base64_BitcoinCash_Address-8     	 3000000	       602 ns/op
	//BenchmarkEncrypt/Encrypto_with_base64_Ethereum_Address-8        	 3000000	       522 ns/op
	//BenchmarkEncrypt/Encrypto_Bitcoin_Address-8                     	 3000000	       390 ns/op
	//BenchmarkEncrypt/Encrypto_BitcoinCash_Address-8                 	 3000000	       444 ns/op
	//BenchmarkEncrypt/Encrypto_Ethereum_Address-8                    	 5000000	       392 ns/op
	//PASS
	//ok  	github.com/mf-financial/wallet-util/encryption	6.162s

	key := "PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3"
	iv := "@~Pp-6sC3<M8x@RA"

	crypt, err := NewCrypt(key, iv)
	if err != nil {
		b.Fatal(err)
	}

	var benchmarks = []struct {
		name     string
		isBase64 bool
		address  string
	}{
		{"Encrypto with base64 Bitcoin Address", true, "3P3QsMVK89JBNqZQv5zMAKG8FK3kJM4rjt"},
		{"Encrypto with base64 BitcoinCash Address", true, "bitcoincash:qpcu3wz0kln63yck9vyyz7ddxy4uuzh4mqumj9wa63"},
		{"Encrypto with base64 Ethereum Address", true, "0x407d73d8a49eeb85d32cf465507dd71d507100c1"},
		{"Encrypto Bitcoin Address", false, "3P3QsMVK89JBNqZQv5zMAKG8FK3kJM4rjt"},
		{"Encrypto BitcoinCash Address", false, "bitcoincash:qpcu3wz0kln63yck9vyyz7ddxy4uuzh4mqumj9wa63"},
		{"Encrypto Ethereum Address", false, "0x407d73d8a49eeb85d32cf465507dd71d507100c1"},
	}

	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if bm.isBase64 {
					crypt.EncryptBase64(bm.address)
				} else {
					crypt.Encrypt([]byte(bm.address))
				}
			}
		})
	}
	b.StopTimer()
}
