FROM go:1.18
MAINTAINER laya
COPY . /app
RUN go env -w GOPROXY=https://goproxy.cn,direct \
   && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
   && cd /app \
   && go build -o ./laya-go cmd/main.go

FROM debian:stretch-slim
COPY --from=0 /app/laya-go /var/www/code/app/laya-go
##COPY --from=0 /app/config /var/www/code/app/config

WORKDIR /var/www/code/app
CMD ["./laya-go"]
