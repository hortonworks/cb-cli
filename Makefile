BINARY=dp

VERSION ?= $(shell git describe --tags --abbrev=0)-snapshot
PLUGIN_ENABLED ?= false
BUILD_TIME=$(shell date +%FT%T)
LDFLAGS=-ldflags "-X github.com/hortonworks/cb-cli/dataplane/common.Version=${VERSION} -X github.com/hortonworks/cb-cli/dataplane/common.BuildTime=${BUILD_TIME} -X github.com/hortonworks/cb-cli/plugin.Enabled=${PLUGIN_ENABLED}"
LDFLAGS_NOVER=-ldflags "-X github.com/hortonworks/cb-cli/dataplane/common.Version=snapshot -X github.com/hortonworks/cb-cli/dataplane/common.BuildTime=${BUILD_TIME} -X github.com/hortonworks/cb-cli/plugin.Enabled=${PLUGIN_ENABLED}"
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")
CB_IP = $(shell echo \${IP})
ifeq ($(CB_IP),)
        CB_IP = 192.168.64.1
endif
CB_PORT = $(shell echo \${PORT})
ifeq ($(CB_PORT),)
        CB_PORT = 9091
endif

deps: deps-errcheck
	go get -u golang.org/x/tools/cmd/goimports
	curl -o $(GOPATH)/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/v0.17.2/swagger_$(shell echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
	chmod +x $(GOPATH)/bin/swagger

deps-errcheck:
#	go get -u github.com/kisielk/errcheck

formatcheck:
	([ -z "$(shell gofmt -d $(GOFILES_NOVENDOR))" ]) || (echo "Source is unformatted"; exit 1)

format:
	@gofmt -w ${GOFILES_NOVENDOR}

vet:
	GO111MODULE=on go vet ./...

test:
	GO111MODULE=on go test -timeout 30s -race ./...

errcheck:
#	Module support is not on master yet: https://github.com/kisielk/errcheck/issues/152#issuecomment-415945206
#	errcheck -ignoretests -exclude errcheck_excludes.txt ./...

coverage:
	go test github.com/hortonworks/cb-cli/dataplane/... -cover

coverage-html:
	go test github.com/hortonworks/cb-cli/dataplane/... -coverprofile fmt
	@go tool cover -html=fmt
	@rm -f fmt

build: errcheck formatcheck vet test build-darwin build-linux build-windows

build-version: errcheck format vet test build-darwin-version build-linux-version build-windows-version

build-docker:
	@#USER_NS='-u $(shell id -u $(whoami)):$(shell id -g $(whoami))'
	docker run --rm ${USER_NS} -v "${PWD}":/go/src/github.com/hortonworks/cb-cli -w /go/src/github.com/hortonworks/cb-cli -e VERSION=${VERSION} -e GO111MODULE=on golang:1.12 make build

build-darwin:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS_NOVER} -o build/Darwin/${BINARY} main.go

dev-debug-darwin:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS_NOVER} -o /usr/local/bin/${BINARY} main.go

build-linux:
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS_NOVER} -o build/Linux/${BINARY} main.go

build-windows:
	GOOS=windows GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS_NOVER} -o build/Windows/${BINARY}.exe main.go

build-darwin-version:
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS} -o build/Darwin/${BINARY} main.go

build-linux-version:
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS} -o build/Linux/${BINARY} main.go

build-windows-version:
	GOOS=windows GO111MODULE=on CGO_ENABLED=0 go build -a ${LDFLAGS} -o build/Windows/${BINARY}.exe main.go

_init-swagger-generation:
	rm -rf dataplane/api/client dataplane/api/model
	rm -f build/swagger.json
	curl -sL http://$(CB_IP):$(CB_PORT)/cb/api/swagger.json -o build/swagger.json

generate-swagger: _init-swagger-generation
	swagger generate client -f build/swagger.json -c client -m model -t dataplane/api

generate-swagger-dp: 
	rm -rf dataplane/oauthapi/client dataplane/oauthapi/model
	swagger generate client -f http://$(DP_IP):$(DP_PORT)/spec/api-docs/swagger.json -c client -m model -t dataplane/oauthapi

generate-swagger-docker: _init-swagger-generation
	@docker run --rm -it -v "${GOPATH}":"${GOPATH}" -v ${PWD}/build/swagger.json:${PWD}/build/swagger.json  -w "${PWD}" -e GOPATH --net=host quay.io/goswagger/swagger:v0.17.2 \
	generate client -f ${PWD}/build/swagger.json -c client -m model -t dataplane/api

release: build
	rm -rf release
	mkdir release
	git tag v${VERSION}
	git push https://${GITHUB_ACCESS_TOKEN}@github.com/hortonworks/cb-cli.git v${VERSION}
	tar -zcvf release/dp-cli_${VERSION}_Darwin_x86_64.tgz -C build/Darwin "${BINARY}"
	tar -zcvf release/dp-cli_${VERSION}_Linux_x86_64.tgz -C build/Linux "${BINARY}"
	tar -zcvf release/dp-cli_${VERSION}_Windows_x86_64.tgz -C build/Windows "${BINARY}.exe"

release-version: build-version
	rm -rf release
	mkdir release
	git tag v${VERSION}
	git push https://${GITHUB_ACCESS_TOKEN}@github.com/hortonworks/cb-cli.git v${VERSION}
	tar -zcvf release/dp-cli_${VERSION}_Darwin_x86_64.tgz -C build/Darwin "${BINARY}"
	tar -zcvf release/dp-cli_${VERSION}_Linux_x86_64.tgz -C build/Linux "${BINARY}"
	tar -zcvf release/dp-cli_${VERSION}_Windows_x86_64.tgz -C build/Windows "${BINARY}.exe"

release-docker:
	@USER_NS='-u $(shell id -u $(whoami)):$(shell id -g $(whoami))'
	docker run --rm ${USER_NS} -v "${PWD}":/go/src/github.com/hortonworks/cb-cli -w /go/src/github.com/hortonworks/cb-cli -e VERSION=${VERSION} -e GITHUB_ACCESS_TOKEN=${GITHUB_TOKEN} -e GO111MODULE=on golang:1.12 bash -c "make release"

release-docker-version:
	@USER_NS='-u $(shell id -u $(whoami)):$(shell id -g $(whoami))'
	docker run --rm ${USER_NS} -v "${PWD}":/go/src/github.com/hortonworks/cb-cli -w /go/src/github.com/hortonworks/cb-cli -e VERSION=${VERSION} -e GITHUB_ACCESS_TOKEN=${GITHUB_TOKEN} -e GO111MODULE=on golang:1.12 bash -c "make release-version"

upload_s3:
	ls -1 release | xargs -I@ aws s3 cp release/@ s3://dp-cli/@ --acl public-read

linux-test: build-linux
	docker run --rm -it -v ${PWD}/build/Linux/:/usr/sbin/ --name "${BINARY}" alpine sh

# Build DP-CLI locally
# Start a new DataPlane Mock with new Swagger JSON and renewed Mock IP
#   For custom version apply like: 'GIT_FIRST_PARENT=2.8.0-dev.374 make integration-test'
# Execute just one scenario you can start this with: 'CLI_TEST_FILES=spec/integration/credential.rb make integration-test'
integration-test: build-docker
	make -C tests all

# Create then source your local E2E testing environment variables like 'tests/localvars'
# Execute just one scenario you can start this with: 'CLI_TEST_FILES=spec/e2e/credential.rb make e2e-test'
e2e-test:
	make -C tests e2e-test

mod-tidy:
	@docker run --rm -v "${PWD}":/go/src/github.com/hortonworks/cb-cli -w /go/src/github.com/hortonworks/cb-cli -e GO111MODULE=on golang:1.12 make _mod-tidy

_mod-tidy:
	go mod tidy -v

.DEFAULT_GOAL := build

.PHONY: build release