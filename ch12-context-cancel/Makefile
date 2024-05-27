.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	go build

run-ok: build
	./ch12-context-cancel false

run-cancel: build
	./ch12-context-cancel true
