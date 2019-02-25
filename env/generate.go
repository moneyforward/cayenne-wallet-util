package env

import (
	"fmt"
	"strings"

	"github.com/mf-financial/cayenne-wallet-util/encryption"
)

// Generate is generate an encrypted .env file
func Generate(enc encryption.Crypter, importFile string) (string, error) {
	data, err := ImportData(importFile)
	if err != nil {
		fmt.Printf("failed to call env.ImportData(%s) error: %s\n", importFile, err)
	}

	envData := generateEnvData(enc, data)
	return WriteEnvWithMultipleData(envData, true)
}

func generateEnvData(enc encryption.Crypter, data []string) map[string][]string {
	var envData = map[string][]string{}

	for _, row := range data {

		rows := strings.Split(row, " ")

		// envのキーを取得
		envKey := rows[0]
		envValues := strings.Split(rows[1], ",") // envの値を取得（配列の場合はカンマ区切り）

		encryptedData := make([]string, 0, len(envValues))
		for _, envValue := range envValues {
			encryptedData = append(encryptedData, enc.EncryptBase64(envValue))
		}
		envData[envKey] = encryptedData
	}
	return envData
}
