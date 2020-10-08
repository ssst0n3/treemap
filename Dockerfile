FROM golang:1.15
COPY . /build
WORKDIR /build
RUN GO111MODULE="on" GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@http://ftp.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
sed -i "s@http://security.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
sed -i "s@http://deb.debian.org@https://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
apt update && \
apt install -y upx
RUN upx treemap

FROM node:14
COPY treemap_frontend /build
WORKDIR /build
RUN npm config set registry https://mirrors.huaweicloud.com/repository/npm/ && \
npm cache clean -f && \
npm install
RUN npx vue-cli-service build

FROM alpine
ENV WAIT_VERSION 2.7.3
#ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
ADD https://st0n3-dev.obs.cn-south-1.myhuaweicloud.com/docker-compose-wait/release/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=0 /build/treemap /app/
COPY --from=1 /build/dist /app/dist
#COPY file_server /app/file_server
WORKDIR /app
CMD /wait && ./treemap