NAME := "go-clean-arch"
TAG := "latest"
BUILD_DIR := "./out"
ENV := ""

default: build

clean:
	rm -rf ${BUILD_DIR}
	mkdir -p ${BUILD_DIR}/internal/configs
	echo "*" >> ${BUILD_DIR}/.gitignore
	echo "!.gitignore" >> ${BUILD_DIR}/.gitignore

build: clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BUILD_DIR}/cmd/server/app ./src/cmd/server/main.go

image: build config
	docker build --pull --rm -f "Dockerfile" -t ${NAME}:${TAG} "."

config:
	cat ./src/internal/configs/config${ENV}.yaml > ${BUILD_DIR}/internal/configs/config.yaml

test-prepare:
	cat ./src/internal/configs/config${ENV}.yaml > ./src/tests/config.yaml

test: test-prepare
	go test ./src/tests/
