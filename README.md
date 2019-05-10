# go-starter-api

This is an attempt at creating a starter project for creating secure APIs in go-lang.

Consider this a work in progress.  Advice and suggestions are welcome, and should be sent via email to mike.cto@securecloudsolutions.io

## Creating a User

```golang
var err error
newUUID := uuid.Must(uuid.NewV4(), err)
if err != nil {
  panic(err)
}
user := User{
  UUID:     newUUID.String(),
  Username: "bsodmike",
  Email:    "bsodmike@gmail.com",
  Password: "password",
  APIToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cThIIoDvwdueQB468K5xDc5633seEFoqwxjF_xSJyQQ",
}

db.gormDB.NewRecord(user)
db.gormDB.Create(&user)
```

## Docker (Dev)

User `password` as the password.

```
docker-compose up
pgcli -h localhost -p 9001 -U dbuser goapi
```