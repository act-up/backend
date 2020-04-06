FROM golang:alpine AS build-env
WORKDIR /go/src
COPY . /go/src/backend
RUN cd /go/src/backend && go build .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /api
COPY --from=build-env /go/src/backend/api /api

EXPOSE 8080
ENTRYPOINT [ "./main" ]
