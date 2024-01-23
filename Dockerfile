FROM golang:latest
WORKDIR /app
COPY twc.go /app
ENV GOFLAGS="-ldflags=-s -w"
RUN go build twc.go
CMD ["./twc"]
