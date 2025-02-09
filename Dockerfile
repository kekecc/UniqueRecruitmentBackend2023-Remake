FROM golang:1.20 AS builder

ENV GO111MODULE=on 
ENV GOPROXY=http://goproxy.cn,direct

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main .


FROM ubuntu:latest AS prod
WORKDIR /app
ARG PROJECT_NAME=uniquehr

COPY --from=builder /app/main ./${PROJECT_NAME}
COPY --from=builder /app/config.yaml ./config.yaml

EXPOSE 3333

RUN echo "./${PROJECT_NAME} migrate && ./${PROJECT_NAME} server" > ./run.sh &&\
    chmod u+x ./run.sh

CMD ./run.sh
