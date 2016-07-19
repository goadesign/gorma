#!/bin/bash

goagen --design=github.com/goadesign/gorma/example/design bootstrap
goagen --design=github.com/goadesign/gorma/example/design gen --pkg-path github.com/goadesign/gorma
