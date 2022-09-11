# GraphQL with couchbase database


## Why we choose gqlgen package ?
* gqlgen GraphQL Golang library is its schema-first approach.
* With gqlgen, you get automatic code generation for query and mutation resolvers based on your GraphQL schema.

## Step-1
> `go mod init graphql`

### Execute the following command to install the gqlgen library.
> `go install github.com/99designs/gqlgen@latest`

![tidy command](./screens/goglgen_package_installation.png)

### Create a file called tools.go at the root of your project, where you will list all your dependencies
> `touch tools.go`

### Paste the following code snippet to tools.go file.
```go
package tools 

import _ "github.com/99designs/gqlgen"
```

> Windows user can use following code to do above steps at once.
`printf 'package tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go`

![tidy command](./screens/tools_go_file_creation.png)

### Open your terminal/command line program and execute the following command to add all missing dependencies.
> `go mod tidy`

![tidy command](./screens/3_tidy_command_add_dependencies.png)

### Create the Golang GraphQL project structure using the init command.
> `go run github.com/99designs/gqlgen init`

![graphql project structure init](./screens/4_graphql_init.png)

This command creates the project layout and files to run the GraphQL API. The `server.go` file is for running the GraphQL API, and the `schema.graphqls` contains some boilerplate schema definitions for the GraphQL API.

After successful execution of the above command you will have the following resource added into the project.
* directory -> graph
* file -> gqlgen.yml
* file -> server.go

![project structure](./screens/5_project_structure.png)

### Run the project using the following command
> `go run server.go`

![server](./screens/6_server.png)

Open http://localhost:8080/ in your browser to test the GraphQL playground.

## Step-2 Define Your Graphql Schema
A GraphQL schema defines the data requirements that clients can request from the GraphQL API. In this next step, we will describe the GraphQL schema for our Movie API by modifying the schema. graphqls file.



## Resource
* [Tutorial](https://hasura.io/blog/building-a-graphql-api-with-golang-postgres-and-hasura)
* [gqlgen](https://github.com/99designs/gqlgen)