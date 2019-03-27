package main

import (
	"context"
	crypto "crypto/rsa"
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mf-financial/cayenne-wallet-util/rsa"
	"github.com/mf-financial/cayenne-wallet-util/storage"
)

const (
	extensionTxt    = "txt"
	extensionPem    = "pem"
	privateFileName = "private"
	publicFileName  = "public"
)

var (
	parentFolderPath string
	storageType      string
	bucketName       string
	generateType     string
)

// RSAの鍵生成 のためのコマンド。
// 複数回実行すると鍵が上書きされるので要注意
// Usage of ./keygen:
//   -b string
//         bucket name
//   -o string
//         parent folder to output (default ".")
//   -s string
//         storage type (default "local")
//   -g string
//         generate type (default "byte")

func main() {
	parseFlag()

	privateKey, err := rsa.GeneratePrivateKey()
	if err != nil {
		fmt.Printf("GeneratePrivateKey Error: %v", err)
		return
	}
	publicKey := rsa.GeneratePublickey(privateKey)
	privateBytes := encodePrivateKey(privateKey)
	publicBytes := encodePublicKey(publicKey)

	uploader, err := getUploader()
	if err != nil {
		fmt.Printf("getUploader Error: %v", err)
		return
	}
	// public keyとprivate keyがあるので親階層のフォルダまで
	privataKeyPath, publicKeyPath, err := getUploadPath()
	if err != nil {
		fmt.Printf("getUploadPath Error: %v", err)
		return
	}

	if err := uploader.Upload(privataKeyPath, privateBytes); err != nil {
		fmt.Printf("Upload PrivateKey Error: %v", err)
		return
	}
	if err := uploader.Upload(publicKeyPath, publicBytes); err != nil {
		fmt.Printf("Upload PrivateKey Error: %v", err)
		return
	}
}

func parseFlag() {
	flag.StringVar(&parentFolderPath, "o", ".", "parent folder to output")
	flag.StringVar(&storageType, "s", "local", "storage type")
	flag.StringVar(&bucketName, "b", "", "bucket name")
	flag.StringVar(&generateType, "g", "", "generate type")
	flag.Parse()
}

func encodePublicKey(publicKey *crypto.PublicKey) []byte {
	switch strings.ToLower(generateType) {
	case "string":
		return []byte(rsa.EncodePublicKeyToPemBase64String(publicKey))
	case "byte":
		return rsa.EncodePublicKeyToPem(publicKey)
	default:
		return rsa.EncodePublicKeyToPem(publicKey)
	}
}

func encodePrivateKey(privateKey *crypto.PrivateKey) []byte {
	switch strings.ToLower(generateType) {
	case "string":
		return []byte(rsa.EncodePrivateKeyToPemBase64String(privateKey))
	case "byte":
		return rsa.EncodePrivateKeyToPem(privateKey)
	default:
		return rsa.EncodePrivateKeyToPem(privateKey)
	}
}

func getUploader() (storage.Uploader, error) {
	switch strings.ToLower(storageType) {
	case "gcs":
		return storage.NewGCSUploader(context.Background(), bucketName)
	case "local":
		return storage.NewLocalUploader(), nil
	default:
		return nil, errors.New("unrecognized storage type")
	}
}

func getUploadPath() (privatePath string, publicPath string, err error) {

	privateFullFileName, publicFullFileName := getFullFileName()
	if parentFolderPath == "" {
		return privateFullFileName, publicFullFileName, nil
	}

	switch strings.ToLower(storageType) {
	case "gcs":
		return fmt.Sprintf("%s/%s", parentFolderPath, privateFullFileName), fmt.Sprintf("%s/%s", parentFolderPath, publicFullFileName), nil
	case "local":
		privateFullPath, err := filepath.Abs(fmt.Sprintf("%s/%s", parentFolderPath, privateFullFileName))
		if err != nil {
			return "", "", err
		}
		publicFullPath, err := filepath.Abs(fmt.Sprintf("%s/%s", parentFolderPath, publicFullFileName))
		if err != nil {
			return "", "", err
		}
		return privateFullPath, publicFullPath, nil

	default:
		return "", "", nil
	}
}

func getFullFileName() (string, string) {

	var privateFullFileName, publicFullFileName string
	switch strings.ToLower(generateType) {
	case "string":
		privateFullFileName = fmt.Sprintf("%s.%s", privateFileName, extensionTxt)
		publicFullFileName = fmt.Sprintf("%s.%s", publicFileName, extensionTxt)
	case "byte":
		privateFullFileName = fmt.Sprintf("%s.%s", privateFileName, extensionPem)
		publicFullFileName = fmt.Sprintf("%s.%s", publicFileName, extensionPem)
	default:
		privateFullFileName = fmt.Sprintf("%s.%s", privateFileName, extensionPem)
		publicFullFileName = fmt.Sprintf("%s.%s", publicFileName, extensionPem)
	}
	return privateFullFileName, publicFullFileName
}
