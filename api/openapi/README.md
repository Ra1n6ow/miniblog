# install swagger
___
```shell
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

# swagger server
___
```shell
swagger serve -F=swagger --no-open --port 65534 ./api/openapi/openapi.yaml
```