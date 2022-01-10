# Rest-API

## Framework / toolkit
1. Go Fiber
2. Ginkgo
3. Testify/mock
4. Mockery(cli tool to autogenerate mocks)

## Functionality
1. The API follows typical RESTful API design patterns.
2. The data is save in a DB

## Documentation
### Setting up
```bash
$ git clone https://github.com/Murtaza-Udaipurwala/rest-api
$ cd rest-api
$ git switch mongoDB
$ go build
$ ./rest-api
```

### API design
`/user` (Method `GET`)
- Fetches all users from the database
- Returns a json response back.

`/user/{id}` (Method `GET`)
- Fetches user by ID from the database.
- Returns a json response back.

`/user` (Method `POST`)
- Creates a new user and adds him/her in the Database.
- Requires a json request body as such,

```json
{
    "username": "Murtaza Udaipurwala",
    "password": "tryguessingit",
}
```

---
`/update/{id}` (Method `PUT`)
- Updates user data and saves it to the database.
- Requires a json request body as mentioned above.

```json
{
    "password": "nottryguessingthis"
}
```

---
`/delete/{id}` (Method `DELETE`)
- Deletes user by his/her ID.
- User's data is deleted from the database as well.
