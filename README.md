# user-service

## About
This microservice controls the creation/management of users in the system. It's only concerned with storing user data and verifying passwords; it does not manage login sessions. This should be done by the oauth-server or by the front end site.

This server shouldn't be exposed to the Internet, although in practise it shouldn't be a problem.

## Setting up

You'll need Node.JS (we recommend NVM for keeping up to date), a PostgreSQL
database already set up, and `yarn`/`npm`.

```bash
$ yarn # install dependencies
$ cp .env{.example,} # copy our example env file
$ vi .env # edit the env file using your favourite text editor

DB_NAME=my-database-name
DB_USER=username
DB_PASS=password
DB_PORT=5432
PORT=3000

$ yarn run migrate # migrate the database tables
$ yarn start # start the fans!
```

## Endpoints

| Path                | Verb   | Description                             |
|---------------------|--------|-----------------------------------------|
| `/`                 | GET    | Gets user or list of users              |
|                     | POST   | Creates a new user                      |
|                     | PATCH  | Update one or more attributes of a user |
|                     | DELETE | Deletes a user                          |
| `/validatePassword` | POST   | Validates that a password is correct    |
| `/changePassword`   | POST   | Changes a user's password               |

## Response format

Responses will look like this:

```json
{
  "content": {},
  "error": ""
}
```

`content`: the main content of the response

`error`: If there was an error, this will show any extra information

## Tests

Coming soon...! Will most likely use mocha + chai.
