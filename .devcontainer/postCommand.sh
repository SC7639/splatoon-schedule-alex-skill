#!/bin/bash

go mod download
go mod tidy
rm -rf ~/.docker/config.json
echo -e '\n# Auto-Warpify\n[[ "$-" == *i* ]] && printf '\''\eP$f{"hook": "SourcedRcFileForWarp", "value": { "shell": "zsh", "uname": "'$(uname)'" }}\x9c'\'' ' >> ~/.zshrc
zsh