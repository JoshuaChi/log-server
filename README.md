# log-server

## Background
Aim to setup a restful server with high performance to accept business log activities from both frontend and backend.

1. Log all the request information to log file
2. Will maintain the log file and sync with s3 as needed
3. Will provide authentication for any request


## APIs

### Overview

| Interfaces                            | Reqeusts                                                     |
| :--------------------------------- | ---------------------------------------------------------- |
| **Auth**  |                                                                                |
| Login                               |  /api/v1/login                       |
| **Log**  |                                                                                |
| Log                               |  /api/v1/log                       |

#### Login
**Path**
```
/api/v1/login
```

**Method**
```
POST
```

**Request Body**

| Params |  Type  | Required | Description | Default | Demo value |
| :-------: | :----: | :--: | :--------: | :----: | :--------: |
| Username     | String | Y   | User name   | -      | Username |
| Password  | String | Y   | Password | -      | Password      |


``` curl
curl --location --request POST 'localhost:8080/api/v1/login' \
--header 'content-type: application/json' \
--data-raw '{"Username":"username", "Password":"password"}'
```


**Response body**

| Params |  Type   | Description |
| :--------------: | :-----: | :----------------: |
|  access_token   | String  | access token |
|  refresh_token   | String  | refresh token |

**Response JSON**

```json
// Status Code: 200
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjQ4ZWVjZjI4LTQxZTUtNDk3Zi1iNmU4LWI0OTk1OTlkOWY3ZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzOTM4ODIzNSwidXNlcl9pZCI6MX0.9DyIZgZSmQNOEnxbQhUMI2Xp7RlSzfWU-3EsCgLtTLs",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzk5OTIxMzUsInJlZnJlc2hfdXVpZCI6IjEzMDk3NzhkLWJlNGEtNGNlZC1iYWMyLWI4ZGJlNGIzMDYwNiIsInVzZXJfaWQiOjF9.tvqySJG1Mo6Wfr63Sii3KoUySGMq6n9AOVLZ2k8WjeM"
}
```




#### Login
**Path**
```
/api/v1/log
```

**Method**
```
POST
```

**Request Header**

| Params |  Type  | Required | Description | Default | Demo value |
| :-------: | :----: | :--: | :--------: | :----: | :--------: |
| Authorization     | String | Y   | Bear token   | -      | Bearer {{access_token}} |

**Request Body**

| Params |  Type  | Required | Description | Default | Demo value |
| :-------: | :----: | :--: | :--------: | :----: | :--------: |
| Uid     | String | Y   | User Id   | -      | "1E43" |
| Action     | String | Y   | User action   | -      | "shop" |
| Category     | String | Y   | action category  | -      | "product" |
| Subcategory     | String | N   | action sub category  | -      | "shoes" |


``` curl
curl --location --request POST 'localhost:8080/api/v1/log' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzk5OTIxMzUsInJlZnJlc2hfdXVpZCI6IjEzMDk3NzhkLWJlNGEtNGNlZC1iYWMyLWI4ZGJlNGIzMDYwNiIsInVzZXJfaWQiOjF9.tvqySJG1Mo6Wfr63Sii3KoUySGMq6n9AOVLZ2k8WjeM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Uid":"1E43", 
    "Action":"Shop", 
    "Category": "product", 
    "Subcategory": "test"
}'
```

**Response JSON**

```json
// Status Code: 200
```


**Note**
Go Version > 1.5