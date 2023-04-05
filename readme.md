# NFT Marketplace
api for nft market place

## Prerequisites
- Golang 1.16
- Mysql

## Running Aplication
running this script and aplication will be runnning in port 8080. 
```shell script
source ./start.sh
```

## Environment
if there is difference configuration you can edit file `config.json` to configure aplication:
```json
  {
    "database": {
      "host": "localhost",
      "port": "3306",
      "user": "root",
      "pass": "password",
      "name": "nft_marketplace"
    }
  }
```