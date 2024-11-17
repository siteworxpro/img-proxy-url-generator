#!/usr/bin/env bash

go install

for distro in $(go tool dist list)
do
    arrIN=(${distro//\// })

    if [[ ${arrIN[0]} == 'linux' || ${arrIN[0]} == 'darwin'  || ${arrIN[0]} == 'freebsd' || ${arrIN[0]} == 'windows' ]]; then
      echo "Building $distro..."
      GOOS=${arrIN[0]} GOARCH=${arrIN[1]} GO111MODULE=on CGO_ENABLED=0 go build --ldflags="-X 'main.Version=$(git describe --tags --abbrev=0)'" -o "dist/img-proxy-url-generator_${arrIN[0]}_${arrIN[1]}"
      gpg --detach-sign "dist/img-proxy-url-generator_${arrIN[0]}_${arrIN[1]}"
    fi
done