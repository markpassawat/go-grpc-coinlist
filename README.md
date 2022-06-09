# go-grpc-coinlist

## Installation 
Clone this repository with git:
```bash
git clone https://github.com/markpassawat/go-grpc-coinlist.git
```
## Usage
Create docker for postgeSQL on port 5432 with password and max connection setting:

```bash
docker run --name coin-list-database -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres -N 300
```

Create database and insert 150 default coins:

```bash
go run cmd/create-database/main.go
```

Run both command for run server:

```bash
source run gateway
```
For run service:
```bash
source run coin_list
```

Use Postman to testing APIs on `localhost:8080` with (grps-coin-list.postman_collection.json) file :

* GetCoin : `GET /coins/:coin_id`
* GetCoins : `GET /coins`
* CreateCoins : `POST /coins/:coin_id`
* SearchCoins : `GET /coins/:search_text`


## FAQ
If there are any problem, try run this command and do it again.
```bash
  export PATH="$PATH:$(go env GOPATH)/bin"
```
