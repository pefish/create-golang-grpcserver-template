#!/bin/bash

cat go.mod | sed "s@_template_@${NAME}@g" > temp && rm -rf go.mod && mv temp go.mod
cat main.go | sed "s@_template_@${NAME}@g" > temp && rm -rf main.go && mv temp main.go
cat client/main.go | sed "s@_template_@${NAME}@g" > client/temp && rm -rf client/main.go && mv client/temp client/main.go
cat service/helloworld/helloworld.go | sed "s@_template_@${NAME}@g" > service/helloworld/temp && rm -rf service/helloworld/helloworld.go && mv service/helloworld/temp service/helloworld/helloworld.go
cp config/sample.yaml config/local.yaml
