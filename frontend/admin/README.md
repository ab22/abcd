# Frontend Admin

This is the frontend application for the admin module. It contains all views,
styles and the angular application code to run the application, as well as the
bower dependencies and package dependencies to build the application for
production.

## Requirements

- npm
- bower
- grunt-cli

## Installation

To configure the application for the first time, npm and bower must be
installed first.

To install all dependencies for this project, just run:

```shell
npm install && bower install
```

This will create a **node_modules/** folder and a **static/bower_components/**
folder with all the dependencies for the application.

## Testing

Before commiting and uploading the code, run:

```shell
grunt test
```
