FROM golang:1.19
WORKDIR /gtw
COPY go.mod go.sum /gtw/
RUN go mod download
COPY *.go /gtw/
ADD ./pb /gtw/pb
RUN go build -o ./gtw_exec
