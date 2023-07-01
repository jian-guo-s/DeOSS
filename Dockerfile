FROM golang:1.19 as builder

WORKDIR /app
COPY . /app

RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go build -o deoss cmd/main.go
RUN chmod +x deoss

FROM ubuntu:20.04

RUN apt update &&  apt install git curl wget vim util-linux -y

WORKDIR /app
VOLUME /home/hamster/cess/workspace
COPY --from=builder /app/deoss /app/deoss
COPY ./conf.yaml /app/
EXPOSE 8080 4001
CMD ["/app/deoss","run"]