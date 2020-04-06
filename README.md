# Act.Up Backend

## API

The REST API is written in [Golang](https://golang.org/) with the [Gin](https://github.com/gin-gonic/gin) web framework and the [Gorm](https://github.com/jinzhu/gorm) ORM library.


## Database

We are using [PostgreSQL](https://www.postgresql.org/) for the database, with the [pq](https://github.com/lib/pq) Postgres driver for Go.

The `db` folder contains snippets for setting up the database and tables.


## Tests

We are using [Postman](https://www.postman.com/) to create and run API tests. The tests can also be run with NodeJS.


### Install NodeJS Testing Dependencies

```$ npm install follow-redirects```


```$ node <test_name>.js```


## Installation

Fetch and install golang packages:

```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/lib/pq
```

```
go install github.com/gin-gonic/gin github.com/jinzhu/gorm github.com/lib/pq
```

Build and install this package:

```
go build main.go
go install main.go
```
