
### Download postgres latest image
```shell
docker pull postgres:latest
```
### Create and run a container with postgres image
```shell
docker run --name [container_name] -e POSTGRES_PASSWORD=[your_password] -d postgres
```
### Connect to Postgres in docker container
```shell
docker exec -it [container_name] psql -U [postgres_user]
```

## reference
https://medium.com/swlh/building-a-restful-api-with-go-and-postgresql-494819f51810