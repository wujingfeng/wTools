FROM golang:1.17

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY . .
RUN go env && go build -o w_tools .

FROM alpine:latest

WORKDIR /app
COPY --from=0 /app/w_tools ./
COPY --from=0 /app/config.yaml ./
COPY --from=0 /app/resource ./resource

EXPOSE 8888

ENTRYPOINT ./w_tools