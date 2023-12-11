FROM golang:latest
WORKDIR /app
COPY twc.go /app
ENV CFLAGS=-O2
RUN go build -o twc .
CMD ["./twc"]
