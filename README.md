# server-validation-go

## Api Description
Project uses an Api to handle requests for two routes
- serverInfo/{server}
- history

## Api Server

Api is developed using Go 
- To run make sure to download this project below de GOPATH/src or set the GOPATH route to this folder so the imports are found during compilation
- Execute: `go run main.go`
- Api properties can be configured to set desired port in which to run it. ( Default is 8888 )

In addittion it uses local CockroachDb
- Install CockroachDb ( https://www.cockroachlabs.com/docs/stable/install-cockroachdb-mac.html )
- Start a local node by running `cockroach start --insecure --store=node1 --listen-addr=localhost:26666 --http-addr=localhost:8080 --background`
- Connect via terminal by running `cockroach sql --insecure --host=localhost:26666`
- Run the script located on cockroachdb/script.db to create the database

The Api uses http://ip-api.com/json/ and https://api.ssllabs.com/api/v3/analyze?host= to get the information about the servers. 

## Author
Created by MaestroOso
