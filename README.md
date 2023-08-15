# Go MVC Webapp Template

This repository is a Go MVC template to start developing a web app.

## Features

- **Minimal dependencies:** 15 dependencies (including indirect ones). This template uses two strong libraries (Gorm and some gorilla packages).

- **Standalone binary:** Everything is embed into the builded binary: just do `go build` and get a standalone binary that you should execute in your server. (if you use sqlite, the database will be created in the directory where you run the binary)

- **No hidden magic:** What you see is what you have

- **Classic MVC:** Controllers, views, models, and a `config.go` file to setup new routes

- **Testing enabled:** Apart from using the Go's native testing capabilites, the template is ready to use [Rod](https://github.com/go-rod/rod) to create integration tests. Thanks to rod these tests are incredibly fast, and have [no dependencies at all](https://go-rod.github.io/#/compatibility).



## Structure

- **/controller:** Controllers

- **/model:** Go structs representing the database thanks to Gorm

- **/views:** Go's template views

- **/main_test.go:** Integration test example

- **/config.go:** Routes