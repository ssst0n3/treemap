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
ENV NPM_REGISTRY https://mirrors.huaweicloud.com/repository/npm/
ENV NPM_REGISTRY https://registry.npm.taobao.org
COPY treemap_frontend /build
WORKDIR /build
RUN npm config set registry $NPM_REGISTRY && \
npm cache clean -f && \
npm install
RUN npm run-script build

FROM alpine
ENV WAIT_VERSION 2.7.3
# ENV WAIT_RELEASE https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait
ENV WAIT_RELEASE https://st0n3-dev.obs.cn-south-1.myhuaweicloud.com/docker-compose-wait/release/$WAIT_VERSION/wait
ADD $WAIT_RELEASE /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=0 /build/treemap /app/
COPY --from=1 /build/dist /app/dist
#COPY file_server /app/file_server
WORKDIR /app
CMD /wait && ./treemap