#!/usr/bin/env bash

cd src
read -p "enter version : " version
go build -o ../release/fcd-${version} .

echo "updating latest build..."
echo fcd-${version} > ../release/resources/latest_build