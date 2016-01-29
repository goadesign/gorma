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
DIRS=$(shell go list -f {{.Dir}} ./...)
DEPEND=\
	github.com/golang/lint/golint \
	github.com/onsi/ginkgo \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega \
 	golang.org/x/tools/cmd/goimports \

.PHONY: goagen

all: depend lint test

depend:
	@go get $(DEPEND)

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi


test:
	@ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -skipPackage vendor

