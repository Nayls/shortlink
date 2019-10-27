.PHONY: help
help:
	@echo "Make help starting."
	@echo "  generate"
	@echo "  build"
	@echo "  run"
	@echo "  run-it"
	@echo "  run-detach"
	@echo "  exec"

SHELL := /bin/bash
VERSION := 1.0.0

DOCKER_BUILDKIT?=0
CI_REGISTRY_IMAGE?=batazor/shortlink
CI_COMMIT_SHORT_SHA?=$(shell git rev-parse --short=8 -q HEAD)

IMAGE?=alpine
COMMAND?=/bin/ash
ifeq ($(IMAGE),alpine)
	IMAGE_TAG?=alpine
	COMMAND?=/bin/ash
else ifeq ($(IMAGE),buster)
	IMAGE_TAG?=buster
	COMMAND?=/bin/bash
endif

DOCKERFILE?=${PWD}/docker/${IMAGE}.Dockerfile
DOCKER_CONTEXT?=${PWD}


test:
	echo $$TARGET

.: generate

generate:
	@echo "proto generation link entity"
	@protoc -I/usr/local/include -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=plugins=grpc:. \
	pkg/link/link.proto

	@echo "proto generation gRPC-web"
	@protoc -I/usr/local/include -I. \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--go_out=plugins=grpc:. \
	--swagger_out=logtostderr=true,allow_delete_body=true:. \
	--grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
	pkg/api/grpc-web/api.proto
	@mv pkg/api/grpc-web/api.swagger.json docs/api.swagger.json

	@echo "Generate go static"
	@go generate pkg/api/graphql/schema/schema.go

golint:
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done

run:
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/database/redis.yaml \
		-f ops/docker-compose/gataway/traefik.yaml \
		up -d

down:
	@docker-compose down --remove-orphans
