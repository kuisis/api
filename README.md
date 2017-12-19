README
------

Running
=======

```
$ ./postgrest-oauth-server -h
2017/12/19 15:29:45 Started!
Usage of ./postgrest-oauth-server:
  -accessTokenJWTSecret string
        Secret key for generating JWT access tokens (default "morethan32symbolssecretkey!!!!!!")
  -accessTokenTTL int
        Access token TTL in seconds (default 7200)
  -cookieBlockKey string
        Block key for cookie creation. 16, 24 or 32 random symbols are valid (default "16charssecret!!!")
  -cookieHashKey string
        Hash key for cookie creation. 64 random symbols recommended (default "supersecret")
  -dbConnString string
        Database connection string (default "postgres://user:pass@localhost:5432/test?sslmode=disable")
  -mainTemplate string
        Name of main template html file (default "index.html")
  -refreshTokenJWTSecret string
        Secret key for generating JWT refresh tokens (default "notlesshan32symbolssecretkey!!!!")
  -templatePath string
        Path to template html file. With trailing slash (default "./")
  -validateRedirectURI
        Whether validate redirect URI or not. Handy for development (default true)
  -verifyTemplate string
        Name of verify template html file (default "verify.html")
```

Testing with Newman
===================
```
$ cd ./example
$ docker-compose -f docker-compose.yml up -d --build
$ cd ../
$ newman run --bail --ignore-redirects --global-var host=localhost ./postgrest-oauth-server.postman_collection.json

```
