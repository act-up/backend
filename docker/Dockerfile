FROM golang:alpine AS build-env
RUN apk --update add git
WORKDIR /go/src

# Get api dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq

# Get api source code
RUN git clone https://github.com/act-up/backend.git

RUN cd /go/src/backend && go build -o main .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /api
COPY --from=build-env /go/src/backend /api

EXPOSE 8080
ENTRYPOINT [ "./main" ]
