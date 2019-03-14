# cayenne-wallet-util
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

### generate .env file 
```bash
enc -generate -target input.txt
```

#### input.txt format

* column separation is half-width space
* Line separation is line feed code(LF)
* Comma-delimited if value is an array

example test format
```bash
ENV_KEY_01 ENV_VALUE_01
ENV_KEY_02 ENV_VALUE_02,ENV_VALUE_03,ENV_VALUE_04 # value is an array
ENV_KEY_03 ENV_VALUE_05
ENV_KEY_04 ENV_VALUE_05
```
