#!/bin/bash

go mod download
go mod tidy
rm -rf ~/.docker/config.json