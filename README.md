## REST API with Golang

**This repo has the very simple REST API developed with Golang.**

**simple_api.go** is the file that handles the API requests. 

I have used the http package for this. And there is another ready to go package **Mux**. But i wanted to just use only http package to get some insights about it. 

This API can handle the basic CRUD operations.

It has no backend database to store the information,I have just used a variable (I knew, its not recommended and its not an production code but just for the sake of simplicity)

We can start the API server to Serve and Listen on the given port **(i have used 8090)** of the localhost.

