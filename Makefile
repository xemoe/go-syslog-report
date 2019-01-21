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

##############################################################

test: unit
build: bin

clean:
	rm -f bin/*

image:
	docker build --rm -t ${GO_BUILD_IMAGE} .

##############################################################

modules := workers validators input

$(modules): DOCKER_WORKSPACE = $@
$(modules): FORCE
	$(DOCKER_RUN) go get
	$(DOCKER_RUN) go test -v

.PHONY: FORCE

unit: workers validators input

##############################################################

binary := bin/bro-syslog-group bin/bro-count bin/bro-meta

$(binary): DOCKER_WORKSPACE = main
$(binary): BASEBIN=$(notdir $(basename $@))
$(binary): FORCE
	$(DOCKER_RUN) go get
	$(DOCKER_RUN) go build -o ../$@ $(BASEBIN)/$(BASEBIN).go

.PHONY: FORCE

bin: bin/bro-syslog-group bin/bro-count bin/bro-meta
