## Users

Headers

> Authorization: Bearer token*

1. Description: List users
   * **GET** api/v1/users
```
    response: {
      "users": [n]
        0:  {
          "createdAt": date0
          "email": email0
          "emailVerified": true/false
          "id": id0
        }
        ...
        n:  {
          "createdAt": date1
          "email": email1
          "emailVerified": true/false
          "id": idn
        }
    }
```
2. Description: Create new user
   * **POST** api/v1/users
```
    body: {
      "user": {
        "admin" : false,
        "createdAt" : "date",
        "disabled" : false,
        "displayName" : "name",
        "email" : "email",
        "emailVerified" : false
      },
      "redirectURL": "..."
    }
    response: {
      "emailSent": true
        "user": {
        "createdAt": "date"
        "displayName": "name"
        "email": "email"
        "id": id generated
      }
    }
```
3. Description: Get user by id
   * **GET** api/v1/users/:id
```
    response: {
        "user": {
          "createdAt": date
          "displayName": nameUser
          "email": emailUser
          "id": id generated
      }
    }
```
4. Description: Enable/disable a user
   * **POST** api/v1/users/:id/disable
```
    body : {"Disable" : false}
    response: {
      "ok": true
    }
```

## Clients

Headers

> Authorization: Bearer token*

1. Description: List clients
   * **GET** api/v1/clients
```
    response: {
      "clients": [2]
        0:  {
          "id": id0
          "redirectURIs": [1]
            0:  "http://127.0.0.1:5555/callback0"
        }
        1:  {
          "id": id1
          "redirectURIs": [1]
            0:  "http://localhost:5556/callback1"
      }
    }
```
2. Description: Create new client
   * **POST** api/v1/clients
```
    body : {"redirectURIs" : ["http://localhost:5555/sample"]}
    response: {
      "id": id generated
      "redirectURIs": [1]
        0:  "http://localhost:5555/sample"
      "secret": secret generated
    }
```
## Admin

Headers

> Authorization: admin-api-secret

1. Description: Get admin by id
   * **GET** ip-overlord/api/v1/admin/:id
```
    response: {
      "email": email
      "id": id
      "password": pwd
    }
```
2. Description: Create new admin
   * **POST** ip-overlord/api/v1/admin
```
    body: {
      "email" : "email",
      "emailVerified" : false/true
    }
    response: {
      "email": "email"
      "id": id generated
    }
```
3. Description: If exist any admin return true else false
   * **GET** ip-overlord/api/v1/state
```
      response: {
        "AdminUserCreated": true
    }
```