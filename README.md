# A.B.C.D.

Main repository for the Go web API and frontend applications for A.B.C.D.

## Configuration

Before running the project, it is necessary to have a Go workspace and the
$GOPATH environment variable. Read (How to Write Go Code)[https://golang.org/doc/code.html]
to configure the project correctly.

### Downloading the project

```
go get github.com/ab22/abcd
```

### Compiling and Running

To compile and run the project, you can simply run:

```
go run *.go
```

If on Windows, this will create a temp executable, so everytime you run this,
the Windows firewall will ask for permissions to run it. So to avoid that popup
to show everytime, you can instead run:

```
build.sh && run.sh
```

## TODO

### Frontend

* Setup a package.json file.
* Configure bower and a bower.json file with the dependencies.
* Create and configure the jshint files and js code styles (jscsrc).
* Setup grunt and it's tasks.

### Backend

* Configure the backend file for dependencies to run the project.
* Setup the db models and services module.
* Setup the vendors folder for dependencies.
* Configuration variables for databases.

### Others

* Setup a Heroku site to host the application.

