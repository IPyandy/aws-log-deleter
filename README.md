## AWS Log Group Deleter

Simple `go` script to cleanup `cloudwatch` log groups for dev environments.

## Running

Clone first

```shell
git clone https://github.com/IPyandy/aws-log-deleter.git
```

```shell
cd aws-log-deleter
go run .
```

or

```shell
go build -o aws-log-deleter .
./aws-log-deleter
```

or make sure GOPATH is set

```shell
go install .
aws-log-deleter
```
