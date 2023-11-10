# Unit tests APIs for iWF

The APIs are generated using [mockgen](https://github.com/uber-go/mock) by the below commands:

```shell
  mockgen -source=iwf/persistence.go -package=iwftest -destination=iwftest/persistence.go
  mockgen -source=iwf/communication.go -package=iwftest -destination=iwftest/communication.go
  mockgen -source=iwf/workflow_context.go -package=iwftest -destination=iwftest/workflow_context.go
  mockgen -source=iwf/client.go -package=iwftest -destination=iwftest/client.go
```

or running this on sdk root folder

```shell
go generate ./...
```

## Usage

See the [example](./example) for more details.
