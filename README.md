# go-test-rest-api-with-postgres
This is a 4 container project. I'm using `docker-compose` to bring it all up
(AKA orchestration).

* 1 data-only container. This is where Postgres' data will live.
* 1 Postgres container. Because Postgres is awesome!
* 1 Go container. This container will house the Go app.
* 1 Nginx container. This will be used as a reverse proxy between Go and the
wild wild web.

## Setup Docker

docker for mac

## Download, Build, and Run

```
git clone https://github.com/DaisukeHirata/go-test-rest-api-with-postgres.git
cd go-test-rest-api-with-postgres
```

### for production
```
export APP_ENV=production
docker-compose up --build
```

### for development
```
export APP_ENV=dev
docker-compose up --build
```

Then fire up Chrome.

If you see
`Cannot start container 8675309: Cannot link to a non running container` then
just run `docker-compose up -d` again. I'm not sure how to prevent that.

If you can not up go container, 
try `docker-compose up go` to figure out errors

[for more info](https://medium.com/@McMenemy/godorp-docker-compose-for-development-and-production-e37fe0a58d61)

## Testing
With all of the test you’ve been writing for your Go application (right?!), you’ll probably want to run them inside of this environment as well. This is easy to do. We can run the following command in our terminal:
```
docker-compose run go go test -v ./...
```

## Environment Variable
to see environment variable
```
docker-compose run go go env 
```

## Build production image
```
docker build -t go-test-rest-api --build-arg app_env=production . 
```

## Files
This is what each file is doing.

### `main.go`
This run a server that listens on port `:8080` inside its container. When you
first start it up, it will attempt to connect to Postgres and
`create table if not exists`. 

### `Dockerfile`
This file will create a new container image by copying the Go source in your
current directory, building it in the container, and running it in the
container. If you want to see how this is happening check out the
`golang:onbuild` `Dockerfile`
[here].

### `docker-compose.yml`
This file describes the 4 containers that our app will use. You should read the
source of this file. Don't worry, there are comments. To execute this file,
you'll have to say `docker-compose build` to create the images. Then,
`docker-compose up` will start the containers. You can pass the `-d` so that
the containers run in the background, `docker-compose up -d`. Finally, to stop
the running containers, you can use `docker-compose stop`.

### `nginx.conf`
This is a basic configuration file that I grabbed from the stock nginx Docker
image. I just added the server block so that web requests get forwarded to
the Go container. The interesting bit is that I forward requests to
`http://app:8080`, where `app` is an entry in `/etc/hosts` that points to the
Go container's IP address and `:8080` is the port I `EXPOSE` in my
`Dockerfile`.

## More About Linking
When you link two containers together, you get some information about where the
other container is located.

```
go:
    links:
        - postgres:db
```

This is from `docker-compose.yml` with the comments removed and describes the
Go container. It says send location information (ha that rhymes) about the
Postgres container to the Go container.

```
[ Go ] <--loc-info-- [ Postgres ]
```

The Go container will be able to access this information in the form of
environment variables. So you can use `os.Getenv` to access that information in
your code.

```
connInfo := fmt.Sprintf(
	"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
	os.Getenv("DB_ENV_POSTGRES_USER"),
	os.Getenv("DB_ENV_POSTGRES_DBNAME"),
	os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
	os.Getenv("DB_PORT_5432_TCP_ADDR"),
	os.Getenv("DB_ENV_POSTGRES_PORT"),
)
```

You can make sure this information is available if you run the `env` command
in a running container.

```
docker exec gotestrestapiwithpostgres_go_1 env
```

Additionally, Docker will add entries in `/etc/hosts` that have the IP address
of the linked container. You can verify this using this command.

```
docker exec gotestrestapiwithpostgres_go_1 cat /etc/hosts
```

[here]: (https://github.com/docker-library/golang/blob/396f40c6188614c7acd6d8299a0ea71030a056a6/1.4/onbuild/Dockerfile)
