
###############################################################################
# Build on local
###############################################################################

# 対象環境
ENV=${env}

bld-enc:
	go build -i -v -o ${GOPATH}/bin/enc ./tools/encryption/main.go

bld-keygen:
	go build -i -v -o ${GOPATH}/bin/keygen ./tools/rsa/keygen/main.go

bld: bld-enc bld-keygen

crypt-env:
	@if [ ! -e "encrypt_key.env" ]; then \
		echo "encrypt_key.envを配置して下さい"; exit 1 ; \
	fi
	@if [ ! -e "input.txt" ]; then \
		echo "input.txtを配置して下さい"; exit 1 ; \
	fi
	@if [ -z ${ENV} ]; then \
		echo "実行時引数env={dev|stg|prod}を指定してください"; exit 1 ; \
	fi

	. encrypt_key.env
	enc -generate -target input.txt

	# KMS暗号化
	gcloud kms encrypt --ciphertext-file=.env.enc --plaintext-file=.env --key wallet-env-key --keyring=wallet-build-keyring --location=global
	gcloud kms encrypt --ciphertext-file=encrypt_key.env.enc --plaintext-file=encrypt_key.env --key wallet-env-key --keyring=wallet-build-keyring --location=global

	# GCSへアップロード
	gsutil cp .env.enc gs://cayenne-wallet-$(ENV)-env-bucket/envfile/.env.enc
	gsutil cp encrypt_key.env.enc gs://cayenne-wallet-$(ENV)-env-bucket/envfile/encrypt_key.env.enc
