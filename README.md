# ads.api

## Check out

```
mkdir -p $GOPATH/src/github.com/lob-inc
cd $GOPATH/src/github.com/lob-inc
git clone https://github.rakops.com/gatd/rad.api
cd rad.api
```

## Install deps

```
export GO111MODULE=on
make deps
```

## Start

```
export SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)"
docker-compose up
```

Access API with

```
curl localhost:8081/hc
```

Access RSSP API with

```
curl -iv -XPOST http://localhost:8080/v1/login -H 'Content-Type: application/json' -d '{"mail": "test@test.com","password":"securepassword"}'
```

Access AdServer API with
```
curl -XGET http://localhost:8080/v1/ads/campaigns -H 'Content-Type: application/json' -d '{"id": 1}'
```
