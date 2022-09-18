
[x] Design a simple db
[x] Convert the loading/creation of the database into a golang function
[x] CREATE OR ALTER TABLE probably would be useful
[ ] Remove Enum from Postgres, add enum in Golang (despite all that work!)
[ ] Keep the Database Connection open; opening it over and over for every query is expensive and unecessary. (Add ref to it in a struct)
[ ] (maybe) create an admin screen to practice sending a view
[ ] Then, create another module called models
[ ] In models, compose go structs around objects queried from the database
[ ] consider how to map models to api endpoints 