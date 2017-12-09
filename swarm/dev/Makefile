.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/ethereum-alt/go-alternative/cmd/swarm
	go build -o bin/geth     github.com/ethereum-alt/go-alternative/cmd/geth
	go build -o bin/bootnode github.com/ethereum-alt/go-alternative/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ethereum-alt/go-alternative/swarm/...
