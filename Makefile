all: clean test build

.PHONY : clean test build

export ROOT_DIR=${PWD}
export DOCKER_GO_PATH=/usr/src/myapp
export DOCKER_GO_IMAGE="my-golang-app"

##############################################################

DOCKER_RUN= \
	docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH}/$(DOCKER_WORKSPACE) \
		${DOCKER_GO_IMAGE} \
		${DOCKER_CMD}

##############################################################

test: go_test_parser
build: bin

clean:
	rm -f bin/*

build_docker_image:
	docker build --rm -t ${GO_BUILD_IMAGE} .

go_get_parser: DOCKER_WORKSPACE = parser
go_get_parser: DOCKER_CMD = go get
go_get_parser:
	$(DOCKER_RUN)

go_test_parser: DOCKER_WORKSPACE = parser
go_test_parser: DOCKER_CMD = go test -v
go_test_parser: go_get_parser
	$(DOCKER_RUN)

go_get_main: DOCKER_WORKSPACE = main
go_get_main: DOCKER_CMD = go get
go_get_main:
	$(DOCKER_RUN)

##############################################################

binary := bin/group bin/count bin/meta

$(binary): DOCKER_WORKSPACE = main
$(binary): BASEBIN=$(notdir $(basename $@))
$(binary): DOCKER_CMD = go build -o ../$@ $(BASEBIN)/$(BASEBIN).go
$(binary): go_get_main
	$(DOCKER_RUN)

bin: bin/group bin/count bin/meta

##############################################################
