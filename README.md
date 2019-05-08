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

an optional `-region` flag can be passed otherwise it will default to normal `aws sdk` config. If `AWS_DEFAULT_REGION` is set that will be used, if profile in `$HOME/.aws` is set, then use that. Otherwise it will be empty and program will panic if flag not set.

```shell
aws-log-deleter -region=us-east-2
```