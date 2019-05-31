# ads.api

## Check out

```
mkdir -p $GOPATH/src/github.com/lob-inc
cd $GOPATH/src/github.com/lob-inc
git clone https://github.com/lob-inc/rad.api
cd rad.api
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
