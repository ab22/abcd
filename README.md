# A.B.C.D.

Main repository for the Go web API and frontend applications for A.B.C.D.

![Login](http://i.imgur.com/esdXYyA.png)

## Configuration

Before running the project, it is necessary to have a Go workspace and the
$GOPATH environment variable. Read [How to Write Go Code](https://golang.org/doc/code.html)
to configure the project correctly.

### Downloading the project

```shell
go get github.com/ab22/abcd
```

### Database Migrations

It is required to have installed Postgres on the local computer. All migration
files are saved in the migrations folder. To automatically run these queries,
it is recommended to use the [migrate](https://github.com/mattes/migrate) tool.

```shell
go get github.com/mattes/migrate
```

Console syntax to migrate all queries:

```shell
cd github.com/ab22/abcd/
migrate -url postgres://user:pass@host:port/dbname?sslmode=disable -path ./migrations up
```

### Compiling and Running

To compile and run the project, you can simply run:

```shell
go run *.go
```

If on Windows, this will create a temp executable, so everytime you run this,
the Windows firewall will ask for permissions to run it. So to avoid that popup
to show everytime, you can instead run:

```shell
build.sh && run.sh
```

## TODO

### Frontend

☑ Setup a package.json file.

☑ Configure bower and a bower.json file with the dependencies.

☑ Create and configure the jshint files and js code styles (jscsrc).

☑ Setup grunt and it's tasks.

☐ Setup the test frameworks.

☐ Create tests for the application.


### Backend

☑ Configure the backend file for dependencies to run the project.

☑ Setup the db models and services module.

☑ Setup the Godep folder for dependencies.

☑ Configuration variables for databases.

☑ Configure a database migrator. Currently, we are using the GORM
  migrations but it would be better to have sql scripts to migrate
  and create the data.

☐ Create tests for the application.


### Others

☑ Setup a Heroku site to host the application.

