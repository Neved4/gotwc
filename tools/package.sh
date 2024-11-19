#!/bin/sh

version=1.1.0

GOOS=darwin GOARCH=amd64 go build -o twc twc.go
tar -czf twc-$version-x86_64-darwin.tar.gz twc

GOOS=darwin GOARCH=arm64 go build -o twc twc.go
tar -czf twc-$version-arm64-darwin.tar.gz twc

GOOS=linux GOARCH=amd64 go build -o twc twc.go
tar -czf twc-$version-x86_64-linux.tar.gz twc

GOOS=linux GOARCH=arm64 go build -o twc twc.go
tar -czf twc-$version-arm64-linux.tar.gz twc

GOOS=freebsd GOARCH=amd64 go build -o twc twc.go
tar -czf twc-$version-x86_64-freebsd.tar.gz twc

GOOS=freebsd GOARCH=arm64 go build -o twc twc.go
tar -czf twc-$version-arm64-freebsd.tar.gz twc

tar -czvf archive.tar.gz --exclude='.DS_Store' .

rm twc

docker build -t twc .
docker run --name dist twc
docker cp dist:/app/twc .
docker stop dist
docker rm dist
tar -czf twc-$version-arm64-linux-debian.tar.gz twc
