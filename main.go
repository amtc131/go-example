package main

import(
  "database/sql"
  "fmt"
  "github.com/golang-migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/postgres"
  _"github.com/golang-migrate/migrate/v4/source/file"
  "github.com/joho/godotenv"
  _"github.com/lib/pq"
)
func main(){
  db, err := sql.Open("postgres", "postgres://postres:mysecretpassword@localhost:5432/database?sslmode=disable")
  if err != nil{
    panic(err)
  }

  driver, err := postgres.WithInstance(db, &postgres.Config{})

  if err != nil{
    panic(err)
  }
  //NOTE: relative path to folder
  m, err := migrate.NewWithDatabaseInstance(
    "file://./db/dev","postgres", driver
  )

  if err != nil{
    panic(err)
  }

  err = m.Up()
  if err != nil{
    fmt.Println(err)
  }

}
