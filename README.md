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

#### Sign in
**Path**
```
/api/v1/users/sign_in
```

**Method**
```
POST
```

**Request params**

| Params |  Type  | Required | Description | Default | Demo value |
| :----: | :----: | :--: | :--: | :----: | :----: |
| locale | String |  N  | Language | zh-CN  | zh-CN  |

**Request Body**

| Params |  Type  | Required | Description | Default | Demo value |
| :-------: | :----: | :--: | :--------: | :----: | :--------: |
| phone     | String | Y   | Mobile   | -      | 86-1234567890 |
| sms_code  | String | Y   | SMS code | -      | 74616      |

**Response body**

| Params |  Type   | Description |
| :--------------: | :-----: | :----------------: |
|  access_token   | String  | access token |

**Response JSON**

```json
// Status Code: 200
{
    "access_token": "c9BvyhEhPw6ppCGgw8ZY1ktTTx9sktyPFZ8k87AcCao",
    "token_type": "Bearer",
    "expires_in": 604800,
    "refresh_token": "qDsfRzpjzAoksjyxiQw0nVJ0Ozz7DcCYvgew5fLCS_I",
    "created_at": 1597904848
}
```
**Note**
以下与终端用户相关的接口, 都需要带上 access_token，方式：在请求的头部增加 Authorization: Bearer access_token


