#Builder: Build react app
FROM node:latest as nodebuilder

COPY app/package.json app/tsconfig.json /app/
COPY app/public /app/public 
COPY app/src /app/src

WORKDIR /app

RUN npm install --silent \
    && npm install react-scripts -g --silent \
    && npm run build

#Builder
FROM golang:alpine AS gobuilder

LABEL maintainer="Himanshu Shekhar <himanshu.shekhar.in@gmail.com>"

WORKDIR /app

COPY go.mod go.sum server.go /app/
COPY api /app/api
COPY config /app/config
COPY config.yml /app/out/
COPY --from=nodebuilder /app/build /app/out/dist

RUN apk update \
    && apk add --no-cache git \
    && go get -d -v \
    && go build -o goapp \
    && mv goapp /app/out

#Image
FROM alpine

ENV APP_HOME=/app \
    APP_DEVELOPER=imhshekhar47
    
COPY --from=gobuilder /app/out /app
EXPOSE 8080
ENTRYPOINT [ "/app/goapp" ]