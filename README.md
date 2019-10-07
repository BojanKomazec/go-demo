# go-demo
Go Demo Application

## Running go-demo in Docker container and making it connect to PostgreSQL DB running in another container


To build the image based on the Dockerfile, use;
```
$ docker build -t go-demo .
```

Before running go-demo container, use `$ docker network ls` to find the name of the network that PostgresDB container is connected to and then `$ docker inspect <network_name>` to find the IP address of the Postgres DB container.

We can now run the container and pass DB network, IP and other configuration. `docker run` command can look like this:
```
$ docker run -e DB_HOST=172.16.239.2 -e DB_PORT=5432 -e DB_NAME=demo -e DB_USER=postgres -e DB_PASSWORD:postgres --rm -it --mount type=bind,src="$(pwd)/data-vol",target=/go/src/github.com/BojanKomazec/go-demo/data-vol --network=postgres-demo-net --name go-demo go-demo
```

To run the app without executing demos that require env variables listed above, we can simply run:
```
$ docker run --rm -it --mount type=bind,src="$(pwd)/data-vol",target=/go/src/github.com/BojanKomazec/go-demo/data-vol --name go-demo go-demo
```
Our app can pick up target directory via env variable: we need to add `OUTPUT_DIR=./data-vol` to `.env` file.

To stop this Docker container, run:
```
$ docker stop go-demo
```

## Running go-demo on the local dev host

`go run` compiles go code and then runs the binary.

```
$ go run cmd/main.go
```
To build an app and run it with command line args:
```
$ make build-linux && ./bin/go-demo -help
$ make build-linux && ./bin/go-demo -postgres
$ make build-linux && ./bin/go-demo -postgres=false
```


## DB and test data

To spin off a Postgres instance, clone https://github.com/BojanKomazec/postgres-demo and launch it with docker-compose.

To create a test table and populate it with some dummy data, you can use an example from e.g. http://www.postgresqltutorial.com/postgresql-array/:
```
CREATE TABLE contacts (
   id serial PRIMARY KEY,
   name VARCHAR (100),
   phones TEXT [],
   magic_numbers INTEGER[]
);

INSERT INTO contacts (name, phones, magic_numbers)
VALUES
   (
      'John Doe',
      ARRAY [ '(408)-589-5846', '(408)-589-5555' ],
      ARRAY [ 1, 11, 111 ]
   ),
   (
      'Lily Bush',
      '{"(408)-589-5841"}',
      '{ 2, 22, 222}'
   ),
   (
      'William Gate',
      '{"(408)-589-5842","(408)-589-58423"}',
      '{ 3, 33, 333}'
   );
```
# Profiling

## CPU Profiling

Run:
```
$ make test-bench-cpuprofile pkg="./internal/pkg/package_name/"
```
Example:
```
$ make test-bench-cpuprofile pkg="./internal/pkg/datatypesdemo/"
```
This will create file `cpu.out` in the current/root directory.

To open `cpu.out` and enter interactive mode in *pprof* tool:
```
$ go tool pprof cpu.out
```
To check how much time it takes to execute parts of function that were called in benchmarks:
```
(pprof) list <function_name>
```
Example:
```
$ go tool pprof cpu.out
File: datatypesdemo.test
Type: cpu
Time: Jul 5, 2019 at 4:50pm (BST)
Duration: 1.70s, Total samples = 1.94s (113.98%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list copyList
Total: 1.94s
ROUTINE ======================== github.com/BojanKomazec/go-demo/internal/pkg/datatypesdemo.copyList in /home/bojan/dev/go/src/github.com/BojanKomazec/go-demo/internal/pkg/datatypesdemo/datatypesdemo.go
     210ms      1.31s (flat, cum) 67.53% of Total
         .          .     61:func copyList(in []string) []string {
         .          .     62:   var out []string
      80ms       80ms     63:   for _, s := range in {
     130ms      1.23s     64:           out = append(out, s)
         .          .     65:   }
         .          .     66:   return out
         .          .     67:}
(pprof)
```
To exit, type `exit`:
```
(pprof) exit
```

*benchcmp* tool automatically calculates differences (in %) in performances before and after code changes:
```
../go/src/github.com/BojanKomazec/go-demo$ make test-bench > old.txt
../go/src/github.com/BojanKomazec/go-demo$ make test-bench > new.txt
../go/src/github.com/BojanKomazec/go-demo$ benchcmp old.txt new.txt
benchmark                       old ns/op     new ns/op     delta
BenchmarkCopyList1_100x16-4     1400          1404          +0.29%
BenchmarkCopyList2_100x16-4     467           468           +0.21%

benchmark                       old allocs     new allocs     delta
BenchmarkCopyList1_100x16-4     8              8              +0.00%
BenchmarkCopyList2_100x16-4     1              1              +0.00%

benchmark                       old bytes     new bytes     delta
BenchmarkCopyList1_100x16-4     4080          4080          +0.00%
BenchmarkCopyList2_100x16-4     1792          1792          +0.00%
```
NOTE: The example above does not show any improvements as in my codebase I left both implementations of `copyList` but added number in the name of each iteration: `copyList1`, `copyList2`,...Repost would show improvement if `old.txt` contained benchmark for 1st implementation of `copyList` and `new.txt` benchmarks of the same function but after modifying/improving it.
