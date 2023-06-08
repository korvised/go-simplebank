# ***Migration Database***

## Create migration file in the target folder:

```shell
    // migate database
    migrate create -ext sql -dir db/migration -seq init_schema
    // push sql
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verose up
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