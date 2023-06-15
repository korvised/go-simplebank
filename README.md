# ***Migration Database***

## Create migration file in the target folder:

```shell
    // migate database
    migrate create -ext sql -dir db/migration -seq init_schema
    // push sql
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verose up
    
    // add new table
    migrate create -ext sql -dir db/migration -seq add_users

```

# *** Makefile ***
```shell
    // run command in Makefile
    make createdb
```

# *** SQLC ***
```shell
    // sql genter crud go function
    // init
    sqlc init
    // push sql
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verose up
```

# *** MOCK ***
```shell
    // after install golang mock.
    // if can run below command run this command first
      go get github.com/golang/mock/    
    // init
     mockgen -package mcokdb -destination db/mock/store.go  github.com/korvised/go-simplebank/db/sqlc Store

```