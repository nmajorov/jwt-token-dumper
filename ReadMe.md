# JWT token parser dumper 

read JWT token and dump it to stdout

example:

            cat token.txt | xargs ./jwt-token-dumper 

            {
              "acr": "1",
              "allowed-origins": [],
              "aud": "java",
              "auth_time": 0,
              "azp": "java",
              "email": "nmajorov@redhat.com",
              "exp": 1528398881,
              "family_name": "Majorov",
              "given_name": "Nikolaj",
              "iat": 1528398581,
              "iss": "http://ssodemo.niko-cloud.ch:8080/auth/realms/demo",
              "jti": "85c3fba7-d6b5-4fd9-b68d-8cd669963ae1",
              "name": "Nikolaj Majorov",
              "nbf": 0,
              "preferred_username": "niko",
              "realm_access": {
                "roles": [
                  "uma_authorization"
                ]
              },
              "resource_access": {
                "account": {
                  "roles": [
                    "manage-account",
                    "manage-account-links",
                    "view-profile"
                  ]
                }
              },
              "session_state": "03c5c2bb-ce75-4a8f-9e9f-c803b54888b6",
              "sub": "8e399b69-1573-46be-9629-c8dd1680c889",
              "typ": "Bearer"
            }
            Issuer: http://ssodemo.niko-cloud.ch:8080/auth/realms/demo
            Audience: java
            Expiration: 2018-06-07 21:14:41 +0200 CEST
