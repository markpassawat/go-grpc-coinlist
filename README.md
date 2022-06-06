# go-grpc-coinlist

## Installation 
Clone this repository with git:
```bash
  git clone https://github.com/markpassawat/go-grpc-coinlist.git
```
## Usage
Run both command for run server:
```bash
  source run gateway
```
For run service:
```bash
  source run coin_list
```

Use Postman to testing APIs on `localhost:8080` :

* GetCoin : `GET /coins/:coin_id`
* GetCoins : `GET /coins`
* CreateCoins : `POST /coins/:coin_id`
* UpdateCoins : `PATCH /coins`
* DeleteCoin : `DELETE /coins/:coin_id`


## FAQ
If there are any problem, try run this command and do it again.
```bash
  export PATH="$PATH:$(go env GOPATH)/bin"
```
