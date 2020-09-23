FROM golang:1.14
COPY . /build
WORKDIR /build
RUN GO111MODULE="on" GOPROXY="https://goproxy.io" go build

#FROM node:14
#COPY skilltree_frontend /build
#WORKDIR /build
#RUN npm config set registry https://mirrors.huaweicloud.com/repository/npm/ && \
#npm cache clean -f && \
#npm install
#RUN npm run-script build

FROM ubuntu
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=0 /build/treemap /app/
#COPY --from=1 /build/dist /app/dist
#COPY file_server /app/file_server
WORKDIR /app
CMD /wait && ./treemap