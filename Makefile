VERSION?=0.0.1
APPLICATION=alfred-jb
PACKAGE_NAME=JetBrains.alfredworkflow

# User definded
ARCH?=amd64
OS?=darwin
CGO?=0

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: install-dependencies
install-dependencies:
	go mod tidy

.PHONY: linters
linters:
	golangci-lint run ./...

.PHONY: build-app
build-app: install-dependencies
build-app:
	env GOARCH=${ARCH} GOOS=${GOOS} CGO_ENABLED=${CGO} go build -i -a \
    	-o ./build/app/${APPLICATION} ./cmd/main.go

.PHONY: build-osx
build-osx: ARCH = amd64
build-osx: GOOS = darwin
build-osx: CGO = 0
build-osx: build-app

.PHONY: prepare-workflow
prepare-workflow:
	mkdir -p ./build
	cp -R ./workflow ./build
	sed -i -e 's/REPLACE_VERSION/${VERSION}/g' ./build/workflow/info.plist

.PHONY: build-workflow
build-workflow: clean prepare-workflow linters build-osx
build-workflow:
	cp ./build/app/${APPLICATION} ./build/workflow/
	cd ./build && \
	zip -j ${PACKAGE_NAME} ./workflow/alfred-jb ./workflow/icon.png ./workflow/info.plist