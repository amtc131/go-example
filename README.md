# go-example

## Creating the Docker Container
So to start we want to create our bare bone Postgres container:
``
docker run --name my-db -e POSTGRES_PASSWORD=pwd -e POSTGRES_DB=database -d -p 5432:5432 postgres
``
