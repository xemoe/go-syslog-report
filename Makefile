all: clean test build

.PHONY : clean test build
export ROOT_DIR=${PWD}
export DOCKER_GO_PATH="/usr/src/myapp"
export DOCKER_GO_IMAGE="my-golang-app"

build: go_get_main go_build_main_group

test: go_get_parser go_test_parser

clean:
	rm -f bin/group

build_docker_image:
	docker build --rm -t ${GO_BUILD_IMAGE} .

go_get_parser:
	docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH}/parser \
		${DOCKER_GO_IMAGE} \
		go get

go_get_main:
	docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH}/main \
		${DOCKER_GO_IMAGE} \
		go get

go_test_parser:
	docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH}/parser \
		${DOCKER_GO_IMAGE} \
		go test -v

go_build_main_group:
	docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH}/main \
		${DOCKER_GO_IMAGE} \
		go build -o ../bin/group group/group.go \
