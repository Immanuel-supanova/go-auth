# go-auth


```
// .env

DB="host=host user=user password=pass dbname=dbname port=port sslmode=disable"
DEBUG=<Either true or false>
```

```
// main.go

func init() {
	settings.Config()
	database.ConnectToDb()
	database.SyncDatabase()
}

...
func main() {
	UserServer.ListenAndServe()
}
```