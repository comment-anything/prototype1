
You do not have to understand how every aspect of this project works to be able to help on it.

In fact you don't have to know much go code if you know .sql.


## First, Check out the database schema

It is easy to view the schema [At this shared link on dbdiagram.io](https://dbdiagram.io/d/632dbad97b3d2034ff9edeb0).

As of 9/23 there are two easy ways to help.

## Five Ways to Help a Lot

## Sorted From Easiest to Hardest

### 1. Write a very short sql query and use it to generate Go code.

> good if you know just a bit of sql

Right now, some of the queries we need for Users are written in `database/query`. They look like this:

```sql
-- name: GetUserByUserName :one
SELECT * FROM "Users"
WHERE "username" = $1 LIMIT 1;
```
Sqlc uses this query to generate the Go function "GetUserByUserName' which populates a generated User struct. 

Some of the other queries needed to CRUD (Create, Read, Update, Destroy) User are also defined in `users.sql`.


However if you look at the [schema](https://dbdiagram.io/d/632dbad97b3d2034ff9edeb0) there are many more queries which need to be generated; several for each table, and some for combinations of tables!

If you can figure out some of those queries, it would build out the program a lot. 

Check out the [sqlc documentation](https://docs.sqlc.dev/en/stable/); it's very straight forward and it should all be set up.

All you need to do is run:

`make sqlc` and `database/sqlc` will be regenerated from the files in `database/query`.


### 2. Write Unit Tests

Go testing works like this.

        1. Let f1l3 be a file you want to test without an extension
        2. Let p4ck be f1l3s package
        3. Work in a new file named `f1l3_test.go` in the same directory
        4. Add the same package directive : `package p4ck`
        5. Let F0nc be the name of a function you want to test
        6. Create a function called TestF0nc
        7.  with the signature `TestF0nc(t *testing.T)`

Run `go test ./path/to/package` to test that package.

Look at `sqlc/users_test.go` for an example.

I do not believe `TestMain` needs to be defined more than once. I believe it is executed once when package loads and `testQueries` will be available across the entire `database` package. 

### 3. Improve the Schema

There may be ways to improve the schema. You can test out using the `migrate` commands, some of which are defined in the makefile. 

The current migration is setup with [golang-migrate](https://github.com/golang-migrate/migrate) as a cli tool. It could also be used as a library.

Migrations are for if/when you want to change the schema. You write a query that can change all the records and one that can reverse it. You store them in an efficient way so it is easy to step through the database 1 change at a time. Like a massive undo / redo list.

[More info on what migrations are.](https://www.prisma.io/dataguide/types/relational/what-are-database-migrations)

### 4. Make a temporary login/register form template in ./templates 

Test out some html styling and object population and serve it to a random API endpoint to test it out.

Add a handler for it in main.
```
func main() {
    http.HandleFunc("/register", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
Serve up a [template](https://go.dev/doc/articles/wiki/#tmp_6).

### 5. Begin the back end with User Auth

I am still thinking about User Authentication. Here are some thoughts.

 - A register api end point needs to be defined. For now, we can serve a very simple golang templated form.
 - That should connect to the database and create a user.
 - It should have a unit test.
 - We need to use cookies and tokens. 
 - User needs to be able to log in and have proper tokens to maintain logged-in status. CreateUser and so forth should all be designed.
 - Maybe we just store sessions in RAM not in the database. Or we add it to the schema and migrate.

