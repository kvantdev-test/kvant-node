GOTOOLS = \
	github.com/mitchellh/gox \
	github.com/golang/dep/cmd/dep \
	github.com/alecthomas/gometalinter \
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/gobuffalo/packr/packr
PACKAGES=$(shell go list ./... | grep -v '/vendor/')
BUILD_TAGS?=kvant
BUILD_FLAGS=-ldflags "-s -w -X kvant/version.GitCommit=`git rev-parse --short=8 HEAD`"

all: check build test install

check: check_tools ensure_deps

docker: build_docker

########################################
### Build

build:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o build/kvant ./cmd/kvant/

build_c:
	CGO_ENABLED=1 go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS) gcc cleveldb' -o build/kvant ./cmd/kvant/

install:
	CGO_ENABLED=0 go install $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' ./cmd/kvant


########################################
### Tools & dependencies

test:
	CGO_ENABLED=1 CGO_LDFLAGS="-lsnappy" go test --count 1 --tags "gcc cleveldb" ./...

check_tools:
	@# https://stackoverflow.com/a/25668869
	@echo "Found tools: $(foreach tool,$(notdir $(GOTOOLS)),\
        $(if $(shell which $(tool)),$(tool),$(error "No $(tool) in PATH")))"

get_tools:
	@echo "--> Installing tools"
	./scripts/get_tools.sh

update_tools:
	@echo "--> Updating tools"
	@go get -u $(GOTOOLS)

#Run this from CI
get_vendor_deps:
	@rm -rf vendor/
	@echo "--> Running dep"
	@go mod vendor

#Run this locally.
ensure_deps:
	@rm -rf vendor/
	@echo "--> Running dep"
	@go mod vendor

########################################
### Formatting, linting, and vetting

fmt:
	@go fmt ./...

metalinter:
	@echo "--> Running linter"
	@gometalinter.v2 --vendor --deadline=600s --disable-all  \
		--enable=deadcode \
		--enable=gosimple \
	 	--enable=misspell \
		--enable=safesql \
		./...
		#--enable=gas \
		#--enable=maligned \
		#--enable=dupl \
		#--enable=errcheck \
		#--enable=goconst \
		#--enable=gocyclo \
		#--enable=goimports \
		#--enable=golint \ <== comments on anything exported
		#--enable=gotype \
	 	#--enable=ineffassign \
	   	#--enable=interfacer \
	   	#--enable=megacheck \
	   	#--enable=staticcheck \
	   	#--enable=structcheck \
	   	#--enable=unconvert \
	   	#--enable=unparam \
		#--enable=unused \
	   	#--enable=varcheck \
		#--enable=vet \
		#--enable=vetshadow \

metalinter_all:
	@echo "--> Running linter (all)"
	gometalinter.v2 --vendor --deadline=600s --enable-all --disable=lll ./...


###########################################################
### Local testnet using docker

# Build linux binary on other platforms
build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

build-compress:
	upx build/kvant
