.POSIX:

GOFLAGS = -gcflags='all=-l -B -wb=false' -ldflags='-w -s' -trimpath

all: build

build:
	go build ${GOFLAGS}

test: coverage.out
	go test -v -race -buildvcs -coverprofile=coverage.out
	go tool cover -html=coverage.out
	gotestsum -f pkgname .

clean:
	go clean
	rm -f coverage.out
