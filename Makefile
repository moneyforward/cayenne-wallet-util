
###############################################################################
# Build on local
###############################################################################
bld:
	go build -i -v -o ${GOPATH}/bin/enc ./tools/encryption/main.go

crypt-env:
	@if [ ! -e "encrypt_key.env" ]; then \
		echo "require encrypt_key.env"; exit 1 ; \
	fi
	@if [ ! -e "input.txt" ]; then \
		echo "require input.txt"; exit 1 ; \
	fi

	. encrypt_key.env
	enc -generate -target input.txt

	# KMS暗号化
	gcloud kms encrypt --ciphertext-file=.env.enc --plaintext-file=.env --key wallet-env-key --keyring=wallet-build-keyring --location=global
	gcloud kms encrypt --ciphertext-file=encrypt_key.env.enc --plaintext-file=encrypt_key.env --key wallet-env-key --keyring=wallet-build-keyring --location=global

	# GCSへアップロード
	gsutil cp .env.enc gs://wallet-env-bucket/envfile/.env.enc
	gsutil cp encrypt_key.env.enc gs://wallet-env-bucket/envfile/encrypt_key.env.enc
