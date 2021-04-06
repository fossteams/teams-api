# Login Flow

1. `Authorize`
   A `GET` request to
   https://login.microsoftonline.com/TENANT-ID/oauth2/authorize
   is performed. This request includes as parameters:
   ```
   response_type: id_token
   client_id: 5e3ce6c0-2b1f-4285-8d4b-75ee78787346`
   redirect_uri: https://teams.microsoft.com/go
   state: RANDOM_UUID
   client-request-id: RANDOM_UUID
   x-client-SKU: Js
   x-client-Ver: 1.0.9
   nonce: 9RANDOM_UUID
   ```

2. The user types his email address and press "Sign In"
3. `GetCredentialType`  
   A `POST` request to https://login.microsoftonline.com/common/GetCredentialType?mkt=en-US
   is performed. By inputting the credentials the user is redirected to the company's
   ID Broker that in turn returns a SAML via ADFS (Active Directory Federation Services).
   
4. User logs in with his own credentials on the company's ID Broker

5. A `POST` request with the SAML response is submitted to
   https://login.microsoftonline.com/login.srf. In turn, a redirection to
   https://device.login.microsoftonline.com/?request=REQUEST_TOKEN&flowToken=FLOW_TOKEN
   is performed
   
6. A `POST` request to https://login.microsoftonline.com/common/DeviceAuthTls/reprocess
   is performed with a body containing:
   - `ctx`: The previous `REQUEST_TOKEN`
   - `flowtoken`: The `FLOW_TOKEN` from the previous request
    
7. Some cookies are set:
   - `ESTSAUTHPERSISTENT`
   - `ESTSAUTH`
   - `ESTSAUTHLIGHT`
    
5. Multi-Factor Authentication is performed via https://login.microsoftonline.com/common/SAS/BeginAuth.
   A `POST` request is performed with the following fields:
   - `AuthMethodId: "OneWaySMS"`
   - `Method: "BeginAuth"`
   - `ctx: "REQUEST_TOKEN"`
   - `flowToken: "FLOW_TOKEN"`
    
6. The user fulfills the 2FA requirement by making a `POST`
   request to https://login.microsoftonline.com/common/SAS/EndAuth with:
   - `Method: "EndAuth"`
   - `SessionId: "UUID"`
   - `FlowToken: "FLOW_TOKEN`
   - `Ctx: "REQUEST_TOKEN`
   - `AuthMethodId: "OneWaySMS"`
   - `AdditionalAuthData: "USER_INPUT"`
   - `PollCount: 1`
    
   In this case, `USER_INPUT` is the SMS code.

7. The user confirms whether he wants to stay signed in or not

8. A `POST` request to https://login.microsoftonline.com/kmsi is performed

9. A `POST` request to https://company-name-onmicrosoft-com.eu001.access-control.cas.ms/aad_login.  
   This request is performed with an `id_token` in the body, which seems to be
   a [JWE](https://tools.ietf.org/id/draft-miller-jose-jwe-protected-jwk-00.html) token.
   
10. The previous `POST` request redirects the user to `https://teams.microsoft.com/go#id_token=XXXX`
    where `id_token` is a JWT that has the following relevant parts:
    ```json
    {
        "header": {
            "alg": "RS256",
            "kid": "KEY_ID",
            "typ": "JWT",
            "x5t": "KEY_ID_AS_ABOVE"
        },
        "payload": {
            "aud": "5e3ce6c0-2b1f-4285-8d4b-75ee78787346",
            "iss": "https://sts.windows.net/TENANT-ID/",
            "iat": 1617721484,
            "nbf": 1617721484,
            "exp": 1617725384,
            "acct": 0,
            "aio": "B64_ENCODED_STRING",
            "amr": [
                "pwd",
                "mfa"
            ],
            // ...
        },
        // ...
    }
    ```
    
    Note that this token is still not the right one!

14. Another `/authorize` request is performed, but this time the following appears
    in the query string:
    ```
    resource: https://api.spaces.skype.com
    ```
    
    The redirect_uri is still `https://teams.microsoft.com/go`, but this time
    around, after the simplified login-flow (since we're already logged in),
    we get a Skype JWT:
    ```json
    {
        "header": {
            "alg": "RS256",
            "kid": "KEY_ID",
            "nonce": "NONCE",
            "typ": "JWT",
            "x5t": "KEY_ID"
        },
        "payload": {
            "aud": "https://api.spaces.skype.com",
            "iss": "https://sts.windows.net/364e5b87-c1c7-420d-9bee-c35d19b557a1/",
            "iat": 1617721489,
            "nbf": 1617721489,
            "exp": 1617725389,
            "acct": 0,
            "acr": "1",
            "aio": "B64_ENCODED_STRING",
            "amr": [
                "pwd",
                "mfa"
            ],
            "appid": "5e3ce6c0-2b1f-4285-8d4b-75ee78787346",
            // ...
      },
      // ...
    }
    ```
    
15. Finally, the Skype JWT is used to perform a Teams API request, the
    first call seems to be to https://teams.microsoft.com/api/authsvc/v1.0/authz, and
    the token is simply passed as an
    `Authorization: Bearer XXXXX`.
    
    The result of that request returns _another_ token though:
    ```json
    {
        "header": {
            "alg": "RS256",
            "kid": "102",
            "typ": "JWT",
            "x5t": "ANOTHER_KEY_ID?"
        },
        "payload": {
            "iat": 1617721790,
            "exp": 1617808189,
            "skypeid": "orgid:SOME-UUID",
            "scp": 780,
            "csi": "1617721489",
            "tid": "ANOTHER_UUID",
            "rgn": "emea"
        }
        // ...
    }
    ```
    
    This token, as reported in the previous API call and in the token itself,
    is valid for a full day. Yay!  
    
    When performing some requests, this token is used and referenced as
    `Authorization: skypetoken=eyJhb...`.