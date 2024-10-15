# API Authentication

## Overview
This document outlines the technical details of Docswap API authentication system.

## Bearer tokens
The API uses Bearer tokens to authenticate users when a request is made to a secured endpoint.
This token must be provided in the request `Authorization` header like so:

```
Authorization: "Bearer <jwt>"
```

With our current implementation of the front end authentication system, the tokens are issued
by Azure AD B2C to the user when they login via the front end application. The format of the tokens
follows the OpenID Connect protocol, and is signed by Azure.

When making request to the API via the Docswap front end application, the `Authorization` header
is automatically added to axios requests using an interceptor if the user is logged in via Azure.

When making requests to the API via third party systems, such as Postman or curl, you must
manually specify the header in order to access private resources.

For more information on JSON Web Tokens (JWTs) and Bearer tokens, refer to the following resources:
- [JWT Introduction](https://jwt.io/introduction)
- [Understanding Bearer Tokens](https://medium.com/@mottammal1993/understanding-bearer-tokens-usage-examples-and-differences-from-api-keys-496f9bfb6038)

## Authentication middleware
In order to secure endpoints of the API, a custom authentication middleware function is
implemented to check the bearer token before any requests can reach the api endpoint logic.

The logic for the authentication middleware is implemented in the `middlewares.Auth()` function
in `middlewares/auth_middleware.go`. This function is used in `routes/router.go` to secure selected endpoints.
Any routes that are defined under the `privateRouterGroup` will require a valid bearer token to be accessed.

The middleware function was created with flexibility and scalability in mind. All the logic
that is written directly in the function is completely generic to any bearer token, decoupling
it from our Azure specific token logic. The core features of this function are:
- Fetch the token from the `Authorization` header
- Format the token in preparation for validation
- Use the validated token claims to determine the user that made the request

## Auth handler interface
### Overview
Following the principles of abstraction and polymorphism, the authentication middleware is decoupled
from any token validation logic via the authentication handler interface. This interface has only 
a single method that must be implemented, which accepts a string token and returns a validated and parsed jwt `Token` pointer.

The interface and implementations are located at `utils/auth/`.

```go
type AuthHandlerInterface interface {
    ParseAndValidateToken(tokenStr string) (*jwt.Token, error)
}
```

The responsibility of this method is to:
- Decode the string token into a jwt.Token, which contains the claims (token data)
- Validate the decoded token signature
- Ensure the token is not expired

The idea behind having this interface is to decouple the middleware from the decoding and validation process.
As far as the middleware is concerned, as long as it is provided a valid token from this function, it will assume
the user is allowed to access the resources.

### Current implementation (Azure AD B2C)
Currently, the authentication handler interface is using an Azure specific implementation. 

The method implemented at `utils/auth/azure_auth_handler.go` decodes the token using the RSA Signing method
and base64 decoding, which is necessary for Azure AD B2C tokens.

The method uses the `jwt.ParseWithClaims()` function to decode the token as well as validate the
signature using a callback function that is passed as a parameter. This callback function must
return a key that will be used to validate the signature. 

Azure requires a public key to validate
the token signature, but these keys are not fixed, they are rotated periodically for better security.
In order to know the correct public key, Azure provides each tenant with a custom url that, when 
requested via an HTTP GET request, will return JSON Web Keys (JWKs) that can be parsed into public keys for validation.

Once the signature has been validated, there are 3 claims that are validated, based on the standard
from OpenID Connect protocol and Azure recommendations:
- The `iss` (issuer) claim, which validates that the token was issued by our Azure tenant. The value
of this claim is compared to the `AZURE_B2C_URL` stored in the `.env`. (see [Environment variables for authentication](#environment-variables-for-authentication) for details)
- The `aud` (audience) claim, which validates that the token was issued for the correct application. The
value of the claim is compared to the `AZURE_CLIENT_ID` stored in the `.env` (see [Environment variables for authentication](#environment-variables-for-authentication) for details)
- The `iat` (issued at) and `exp` (expiration) claims, which validate that the token has been activated
and is not yet expired. The current UNIX time is compared to these values, and it must fall between them to be valid.


## Environment variables for authentication
There a several environment variables that must be set in the `.env` file for the authentication
to work properly. These are private values that should not be saved to git or shared publicly.

### `AZURE_JWKS_URL`
This variable indicates the url provided by Azure for fetching the JSON Web Keys, used to validate
the token signature. For more information, see documentation [here](https://learn.microsoft.com/en-us/entra/identity-platform/access-tokens#validate-tokens).
The `tenant` may also be called "Directory name" in the Azure portal, and the `user_flow_name`
is available under "User flows" in the B2C resource portal.

```
AZURE_JWKS_URL="https://{tenant}.b2clogin.com/{tenant}.onmicrosoft.com/{user_flow_name}/discovery/v2.0/keys"
```
<br>

### `AZURE_B2C_URL`
This variable contains the Azure B2C tenant url, used for validating the issuer of the token.
The `tenant` may also be called "Directory name" in the Azure portal, and the `tenant_directory_id` can be found
under "Directory + subscriptions" which is accessible when switching directories in the portal.

```
AZURE_B2C_URL="https://{tenant}.b2clogin.com/{tenant_directory_id}/v2.0/"
```

<br>

### `AZURE_CLIENT_ID`
This variable contains the value of the client ID for the front end application. It is used to 
validate that the token was issued for the proper application.
The `docswap_front_end_azure_client_id` can be found under the Azure AD B2C portal -> Manage -> App registrations.

```
AZURE_CLIENT_ID={docswap_front_end_azure_client_id}
```
