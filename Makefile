.PHONY:deps
deps:
	(cd ./libs/rad && git submodule update)
	(cd ./libs/rad/server && go mod vendor && GO111MODULE=off make bindata)
	(cd ./libs/rad/server/api && make clean && RSSP_API_PACKAGE=. make generate_gohandler)
	go mod vendor
