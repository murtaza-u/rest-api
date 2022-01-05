# Rest-API

## Framework / toolkit
Gorilla Mux - [Link](https://github.com/gorilla/mux)

## Functionality
1. The API follows typical RESTful API design patterns.
2. The data is save in the DB

## Documentation

### Setting up
```bash
$ git clone https://github.com/Murtaza-Udaipurwala/rest-api
$ cd rest-api
$ go build
$ ./rest-api
```

### API design
`/get/{id}`
- Fetches user by ID from the database.
- Returns a json response back.
- Method `GET`

`/create`
- Creates a new user and add him/her in the Database.
- Requires a json request body as such,
- Method `POST`

```json
{
    "name": "Murtaza Udaipurwala",
    "dob": "2002-12-09",
    "address": "India",
    "description": "Hey, I am a computer genius"
}
```

`/update/{id}`
- Updates users data and saves it into the database.
- Requires a json request body as mentioned above.
- Method `PUT`

`/delete/{id}`
- Deletes user by it ID.
- All the user's data is deleted from the database as well.
