# go-demo
Go Demo Application

## Running go-demo in Docker container and making it connect to PostgreSQL DB running in in another container


To build the image based on the Dockerfile, use;
```
$ docker build -t go-demo .
```

Before running go-demo container, use `$ docker network ls` to find the name of the network that PostgresDB container is conneted to and then `$ docker inspect <network_name>` to find the IP address of the Postgres DB container.

We can now run the container and pass DB network, IP and other configuration. `docker run` command can look like this:
```
$ docker run -e DB_HOST=172.16.239.2 -e DB_PORT=5432 -e DB_NAME=demo -e DB_USER=postgres -e DB_PASSWORD:postgres --rm -it --network=postgres-demo-net --name go-demo go-demo
```
