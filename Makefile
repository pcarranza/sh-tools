
all: build

build:
	go build github.com/pcarranza/sh-tools/cmd/gitclone

install:
	go install github.com/pcarranza/sh-tools/cmd/gitclone
