package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mf-financial/cayenne-wallet-util/encryption"
)

var (
	encode     = flag.String("encode", "", "for encode")
	decode     = flag.String("decode", "", "for decode")
	generate   = flag.Bool("generate", false, "for generate to env")
	targetFile = flag.String("target", "", "target file for convert")
)

var usage = `Usage: %s [options...]
Options:
  -encode  for encode
  -decode  for decode
  -convert  for convert to env, required option is import.go
e.g.:
  enc -encode xxxxxxxx
    or
  enc -decode xxxxxxxx
    or
  enc -generate -target xxxxxxxx
`

// Params is parameter for template file
type Params struct {
	Name      string
	Uppercase string
}

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	if *encode == "" && *decode == "" && !*generate {
		flag.Usage()
		os.Exit(1)
		return
	}
}

func setup() encryption.Crypter {
	key := os.Getenv("ENC_KEY")
	iv := os.Getenv("ENC_IV")

	if key == "" || iv == "" {
		fmt.Println("set Environment Valuable: ENC_KEY, ENC_IV in your .bashrc or .zshrc")
		os.Exit(1)
	}

	crypt, err := encryption.NewCrypt(key, iv)
	if err != nil {
		panic(err)
	}
	return crypt
}

func main() {
	crypt := setup()

	//encode
	if *encode != "" {
		fmt.Println(crypt.EncryptBase64(*encode))
	}

	//decode
	if *decode != "" {
		str, err := crypt.DecryptBase64(*decode)
		if err != nil {
			fmt.Printf("failed to call crypt.DecryptBase64(%s) error: %s\n", *decode, err)
			return
		}
		fmt.Println(str)
	}

	//generate
	if *generate {
		if *targetFile == "" {
			fmt.Printf("targetFile is required as option -target")
			return
		}

		fileName, err := crypt.GenerateToEnv(*targetFile)
		if err != nil {
			fmt.Printf("failed to call crypt.ConvertToEnv(%s) error: %s\n", *targetFile, err)
			return
		}
		fmt.Printf("[generate env file]: %s\n", fileName)
	}
}
