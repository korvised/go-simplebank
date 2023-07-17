# ***Migration Database***

## Create migration file in the target folder:

```shell
    # migate database
    migrate create -ext sql -dir db/migration -seq init_schema
    # push sql
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verose up
    
    # add new table
    migrate create -ext sql -dir db/migration -seq add_users
    
    # migate database add add_sessions
    migrate create -ext sql -dir db/migration -seq add_sessions

```

# *** Makefile ***
```shell
    # run command in Makefile
    make createdb
```

# *** SQLC ***
```shell
    # sql genter crud go function
    # init
    sqlc init
    # push sql
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verose up
```

# *** MOCK ***
```shell
    # after install golang mock.
    # if can run below command run this command first
      go get github.com/golang/mock/    
    # init
     mockgen -package mcokdb -destination db/mock/store.go  github.com/korvised/go-simplebank/db/sqlc Store

```

# *** DBDoc ***
```shell
    # https://dbdocs.io/docs is plugin for build database document
    # install dbdocs
    npm i -g dbdocs
    # login then build
	dbdocs build doc/db.dbml
	# set password
	dbdocs password --set secret --project simple_bank
	
	# DBML generate sql schema	
	# document "https://dbml.dbdiagram.io/cli/" install dbml cli first
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

```