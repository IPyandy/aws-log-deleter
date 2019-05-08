## AWS Log Group Deleter

Simple `go` script to cleanup `cloudwatch` log groups for dev environments.

## Install

### Method 1

```shell
go get -u github.com/IPyandy/aws-log-deleter
```

### Method 2

```shell
git clone https://github.com/IPyandy/aws-log-deleter.git
```

## Running

an optional `-region` flag can be passed otherwise it will default to normal `aws sdk` config. 

- If `AWS_DEFAULT_REGION` is set that will be used
- else if profile in `$HOME/.aws` is set, then use that
- Otherwise it will be empty and program will panic if flag not set.

```shell
aws-log-deleter -region=us-east-2 # flag -region flag is optional if above rules are met
```