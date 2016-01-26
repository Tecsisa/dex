Local login flow by Dex using Authorization Code flow
===============================

## Diagram
```
                                   (7)
                       +---------------------------+
                       |                           |
                       |                           |
+----------+       +---v----+      (2)      +------+----+
|          |  (1)  |        +--------------->           |
|   End -  +-------> Client |      (6)      |    DEX    +---+
|   User   |       |        +--------------->           |   |
|          |       +----^---+               +-----------+   |
+----+-----+            |                                   | (3)
     |                  |                   +-----------+   |
     |                  |        (5)        |           |   |
     |                  +-------------------+ DEX LOGIN <---+
     |                                      |           |
     |                                      +-----^-----+
     |              (4)                           |
     +--------------------------------------------+

```
__(1)__ End-User connects to registered client in Dex.

__(2)__ Client connects to Dex to authenticate End-User, so the client sends its client_id, redirect_uri, response_type and scopes. 

Example:
`http(s)://host:port/auth?client_id=client_id&redirect_uri=https://hostClientRedirect/callback&response_type=code&scope=openid+email+profile&state=`

__(3)__ Dex redirect End-User to Dex login page to authenticate it. Dex creates a unique session_key and session for this client, which state is `NEW`.

__(4)__ End-User prompts its username and password to authenticate. If the login is succesful then Dex changes session state to `IDENTIFIED`.

__(5)__ Dex redirect End-User to redirect uri of the client with the code associated to session created in step 3.

__(6)__ Client retrieves code that was received in step 5 and it is used by the client to retrieve and access token. Client sends to Dex next data to request an acces token:

__Header__
* `client_id` and `client secret` separated by "`:`" and encoded with base64 encoding in `Authorization` header.

__Post form__
* `grant_type` which value is code
* `code` which value is the code retrieved in step 5

__(7)__ Dex verifies data received and retrieve End-User identifier to generate an Acces Token for client that requests the token associated to user that was logged in step 4. Dex generate a JWT token, change session state to `EXCHANGED` and sends JWT token to client.

When client has an Access token it can use it to request resources from Resource Server
