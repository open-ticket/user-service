# Openticket User Service
This is a non-public facing microservice that stores user data, and provides a
JSON API with which to create, read, update and delete users.

This service will be very unlikely to be dealing with user authentication, that
will almost certainly be delegated to a separate auth server, or some admin
panel, if it's using a system like Rails/Django.

## Endpoints
| Endpoint | HTTP Verb | Description         |
|:--------:|-----------|---------------------|
| /        | GET       | Lists all users     |
| /        | POST      | Creates new user    |
| /$userid | GET       | Returns user by id  |
| /$userid | PATCH     | Updates user by id  |
| /$userid | DELETE    | Deletes user by id  |

## User Schema v1
```json
{
  "name": String,
  "email": String,
  "digest": String,
  "created_at": Date
}
```
