#! /usr/bin/make
#
# Makefile for gorma
#
# Targets:
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./... | grep -v /example)
DEPEND=\
	bitbucket.org/pkg/inflect \
	github.com/goadesign/goa.design/tools/mdc \
	github.com/goadesign/goa.design/tools/godoc2md \
	github.com/golang/lint/golint \
	github.com/onsi/ginkgo \
	github.com/spf13/hugo \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega \
	golang.org/x/tools/cmd/goimports \

.PHONY: goagen

all: depend lint test

docs:
	@git clone https://github.com/goadesign/goa.design
	@rm -rf goa.design/content/reference goa.design/public
	@mdc github.com/goadesign/gorma goa.design/content/reference --exclude goa.design --exclude example
	@cd goa.design && hugo
	@rm -rf public
	@mv goa.design/public public
	@rm -rf goa.design

depend:
	@go get -u $(DEPEND)
	@go install $(DEPEND)

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi


test:
	@ginkgo -r  --failOnPending  --race -skipPackage vendor
