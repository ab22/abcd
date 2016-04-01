# A.B.C.D.

[![Join the chat at https://gitter.im/ab22/abcd](https://badges.gitter.im/ab22/abcd.svg)](https://gitter.im/ab22/abcd?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

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

Note: It is required to have **MinGW32/64bit** installed on **Windows**!

## Running the application

### Testing the Backend

Since using 'go test ./...' will take the vendor folder as a valid package
path, it will attempt to test all of the vendored packages. Also, since
this project contains a frontend folder with all of it's frontend
npm/bower modules, 'go test ./...' will also scan those folders taking up
more time for the test run to complete.

To avoid all previously mentioned, use the **test.sh** script to tests current
golang packages.

```shell
./test.sh
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
go build && abcd.exe
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

