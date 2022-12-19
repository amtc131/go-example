# go-example

## Creating the Docker Container
So to start we want to create our bare bone Postgres container:

``
docker run --name my-db -e POSTGRES_PASSWORD=pwd -e POSTGRES_DB=database -d -p 5432:5432 postgres
``
To confirm it is running we can the run:

``
docker ps
``

## Connecting to our DB

``
docker container exec -it my-db bash
``
we can confirm all is running fine with:

 ``
 psql -u {user} {db}
 ``

 If you prefer to just connect to the container and run queries direct you can modify the above:

 ``
 psql -U {user} {db} -c "{SQL_QUERY}"
 ``

## Generating our migration scripts

 ``
 migrate create -ext sql -dir db/schema -seq schema_init
 ``
we can then add our SQL script to stucture our DB:

```
-- schema_init.up.sql
BEGIN;
create table users 
(
id serial primary key,
username varchar(255) not null,
handle varchar(255) not null
);
commit;
```
