
# Prototype1

This will be a minimum proof of concept for the back end to test out some stuff and check feasibility. 

I've already deviated from the database i sketched for this but its viewable in `database/caprototype1.png` 

![A sketch of a relational database](https://github.com/comment-anywhere/prototype1/blob/main/database/carprototype.png)

I think Rapid Prototyping a little is going to help me understand the design needs a lot.

# Installing

- Navigate to your %GOPATH% directory.
- By default, on Windows, it is `C:/Users/%Username%/go/`
- Already in that directory will be `bin` and `pkg` if you have installed go with default settings.
- Create the directories `src/github.com/`
- Inside `.../go/src/github.com/comment-anything/` run the `git clone` operation
- This module will be saved to, for example `C:/Users/karlm/go/src/github.com/comment-anything/prototype1`

This is the best practices way to do things in go. 

# Installing Postgres

- Install Postgres for your system and follow the instructions to set up your global password and server password
- Create a database called ...... (probably postgres)
- Play around with pgAdmin so you can see the database; you may have the manually refresh the tables in pgAdmin to see db updates.

# Set up environment variables

Create a file in the root called `.env`. This will hold private information used across the program. Do not upload this file to github.

Example .env file for default database:

```
# Database setup
HOST=localhost
PORT=5432
USER=postgres
PASSWORD=mypassword
DATABASE_NAME=postgres
```


# Running

```
> go get .
> go run .
```