# go-api-example

```
GET localhost:8080/api/v1/users
```
```
GET localhost:8080/api/v1/users/{username}
```
```
DEL localhost:8080/api/v1/users/{id}
```
```
POST localhost:8080/api/v1/users

Example Request Body :

{
  "username": "Mustafa" ,
  "surname": "Kocatepe"
}
```
```
PATCH localhost:8080/api/v1/users/{id}

Example Request Body :

{
  "username": "Mustafa" 
}
```
```
PUT localhost:8080/api/v1/users/{id}

Example Request Body :

{
  "username": "Mustafa" ,
  "surname": "Kocatepe"
}
```