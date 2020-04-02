# Act.Up Backend

## API

The REST API is written in [Golang](https://golang.org/) with the [Gin](https://github.com/gin-gonic/gin) web framework and the [Gorm](https://github.com/jinzhu/gorm) ORM library for Golang.


## Database

We are using [PostgreSQL](https://www.postgresql.org/) for the database. The `db` folder contains snippets for setting up the database and tables.


## Tests

We are using [Postman](https://www.postman.com/) to create and run API tests. The tests can also be run with NodeJS.


### Install NodeJS Testing Dependencies

```npm install follow-redirects```


```$ node <test_name>.js```
