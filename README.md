# AuthWebApp

## Installation

Run `make deps` OR execute the below commands to install all the dependencies

```bash
    go get -u github.com/gorilla/mux
    go get -u github.com/joho/godotenv
    go get -u github.com/markbates/goth
    go get -u github.com/gorilla/pat
    go get -u github.com/gorilla/sessions
    go get -u github.com/lib/pq
```
Install PostgreSQL

Change the DB Creds in the `.env` file to your instance.

1. Create a Database `auth`
2. Create a schema inside the DB `auth` as `users`
3. Create import the csv file to create the table `auth` inside `users` schema

Note: The social login creds (client id and client secret) are stored in the `.env` file.

## Usage

Run `make build` to build the project
Run `make run` to run the binary after build (located in bin folder)

Open browser and navigate to `localhost:3000`

API to see all the users: `localhost:3000/user/all`
