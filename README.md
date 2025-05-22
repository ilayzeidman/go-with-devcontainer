# go-with-devcontainer

run the server:
go run main.go

invoke the api through unix socket with curl:
curl --unix-socket ./ctl.sock http://localhost/ctl/terminate