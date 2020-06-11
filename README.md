# Golang Api

This an API project designed to expose CRUD endpoints based on a Doctor model using Golang and Gin in a MongoDB database.

To run the project, you must have Go and Mongo installed in your machine. You can install both following these links:

[Golang](https://golang.org/doc/install)

[MongoDB](https://docs.mongodb.com/manual/installation/)

## Running the project and tests

First thing to do is to **git clone** this repository into your machine.

To run the API, use the following command on your terminal:

```
make api
```

To run the tests, use the following command on your terminal:

```
make test
```

The commands above will also download and install the necessary dependencies.

## Making requests

You can use [Postman](https://www.postman.com/downloads/) to make requests to the following endpoints on port **8080**:

```
GET       /v1/doctors - Retrieves all doctor records from database

GET       /v1/doctors/:id - Retrieve a specific doctor by its ID
  
POST      /v1/doctors - Creates a new doctor

PUT       /v1/doctors/:id - Updates an existing doctor by its ID

DELETE    /v1/doctors/:id - Deletes the selected doctor by its ID
```
