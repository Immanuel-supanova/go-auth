# go-auth


```
// .env

DB="host=host user=user password=pass dbname=dbname port=port sslmode=disable"
DEBUG=<Either true or false>
```

```
// main.go

func init() {
	goauth.Config()
	goauth.ConnectToDb()
	goauth.SyncDatabase()
}

...
func main() {
	goauth.UserServer.ListenAndServe()
}
```