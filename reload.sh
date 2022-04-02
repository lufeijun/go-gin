#!/bin/bash

go build index.go &&  kill -SIGHUP `pgrep index` && echo "ok"
