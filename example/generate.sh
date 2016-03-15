#!/bin/bash

goagen --design=github.com/goadesign/gorma/example/design app
goagen --design=github.com/goadesign/gorma/example/design client
goagen --design=github.com/goadesign/gorma/example/design js
goagen --design=github.com/goadesign/gorma/example/design schema
goagen --design=github.com/goadesign/gorma/example/design swagger
goagen --design=github.com/goadesign/gorma/example/design gen --pkg-path github.com/goadesign/gorma
