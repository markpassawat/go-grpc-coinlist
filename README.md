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

Use postman to testing APIs:

[GetCoin] : `GET /coins/{symbol}`
[GetCoins] : `GET /coins`
[CreateCoins] : `POST /coins`
[UpdateCoins] : `PATCH /coins`
[DeleteCoin] : `DELETE /coins/{symbol}`


## FAQ
If there are any problem, try run this command and do it again.
```bash
  export PATH="$PATH:$(go env GOPATH)/bin"
```
