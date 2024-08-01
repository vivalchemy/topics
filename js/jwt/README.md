# jwt

This project simulates the function of the jwt token in the server. It use the auth server and a data server. The tokens are generated and refreshed on the auth server and the data server check for the authentication and provides the corresponding response

## Local development

1. To install dependencies:

```bash
bun install
```

2. To start the `data` server:

```bash
bun run live # this will start the data server on port :3000
```

3. To start the `auth` server:

```bash
bun run auth # this will start the data server on port :4000
```

> The scripts are written in the ./scripts/ directory for easier testing. Use jq for formatting the response

4. To test the tokens:

```bash
./scripts/login.sh <username> # Replace <username> with your username(use one declared in the posts.ts file)


# Output
# {
#    "accessToken" : <accessToken>,
#    "refreshToken" : <refreshToken>
# }
```

```bash
./scripts/fetchPosts.sh <accessToken> # Replace <accessToken> with the generated accessToken

#Output
# [
#   {
#     username: "<username>",
#     content: "This is my first post",
#   },
#   {
#     username: "<username>",
#     content: "This is my second post",
#   }
# ]
```

```bash
./scripts/refresh.sh <refreshToken> # Replace <refreshToken> with the generated accessToken

#Output
# Output
# {
#    "accessToken" : <accessToken>,
# }
```

```bash
./scripts/logout.sh <refreshToken> # Replace <refreshToken> with the generated accessToken

# This will not generate any output
```



General flow you can follow:

1. Generate the access and the refresh token and store the refresh token in local storage and the db for future accessToken generation
1. You can then access the content using the access token
1. Once the access token is expired use the refresh token to generate a new access token
1. To logout delete the refresh token from the local storage and the db
