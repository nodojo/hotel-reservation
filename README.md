# Hotel Reservation Backend

## Outline

- users -> book room from hotel
- admins -> going to check reservations/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources

### MongoDB Driver

Documentation

- [mongodb.com/docs/drivers/go/current/quick-start](https://mongodb.com/docs/drivers/go/current/quick-start)

Installing the MongoDB Client

```
go get go.mongodb.org/mongo-driver/mongo
```

### Fiber for Go

Documentation

- [docs.gofiber.io](https://docs.gofiber.io/)

Installing gofiber

```
go get github.com/gofiber/fiber/v2
```

## Docker

### Installing MongoDB as a Docker container

```
docker run --name mongodb -d mongo:latest -p 27017:27017
```
