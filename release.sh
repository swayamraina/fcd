#!/usr/bin/env bash

cd src
read -p "enter version : " version
go build -o ../release/fcd-${version} .