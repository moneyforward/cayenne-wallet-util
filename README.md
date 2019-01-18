# cayenne-wallet-util
This repo is common packages for both cayenne and wallet-btc/wallet-eth

### encryption
AESによる暗号化/復号化パッケージ  

- 暗号化/復号化に利用するkeyを`.bashrc`もしくは`.zshrc`内の環境変数に設定しておくこと  
- ENV_KEYはサイズに応じてAES-128(key:16byte),AES-192(key:24byte),AES-256(key:32byte)と強度が変わる
- ENV_IVは16バイト固定
```
export 'ENC_KEY=PBc1h^fjKd3Mrug3PBc1h^fjKd3Mrug3'
export 'ENC_IV=@~Pp-6sC3<M8x@RA'
```


## crypto-address-validator
validator for crypto addresses like bitcoin, bitcoin cash, ethereum and  so on


## tools
### encryption
stringをエンコード/デコードするためのツール

```bash
make bld

#encode
enc -e abc

#decode
enc -d xxx
```
