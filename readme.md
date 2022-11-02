# wb l0-blagy

## DB

* postgresql database is running locally on port `5432`
* db name is `l0db`
* db user is `l0user` with password `l0pass`
* `orders` table stores order data
* `items` table stores items in an order

## Service

* NATS Streaming client subscribes to a channel called `orders`. The subscription is durable, so when the service goes offline the data will not be lost (hopefully).
* The incoming data is validated and stored in a database and in cache
* On service initialization the data from the database is loaded into cache
* There is an HTTP server on port `8080` that provides an interface for retrieval of order data by order's UID (the data is retrieved from cache).