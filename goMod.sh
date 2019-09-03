#!/bin/bash

read -p "Enter dir: " dir


cd "$(dirname "$0")/$dir"

go mod init github.com/lesteven/goServer/"$dir"
