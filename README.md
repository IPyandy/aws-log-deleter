## AWS Log Group Deleter

Simple `go` script to cleanup `cloudwatch` log groups for dev environments.

**CAUTION** this will delete all `Cloudwatch LogGroups` in the given region. Only use this for `dev` environments where there could be a plethora of leftover LogGroups.

## Install

### Method 1

```shell
go get -u github.com/IPyandy/aws-log-deleter
```

### Method 2

```shell
git clone https://github.com/IPyandy/aws-log-deleter.git

cd aws-log-deleter
go install .
```

## Running

an optional `-region` flag can be passed otherwise it will default to normal `aws sdk` config.

- If `AWS_DEFAULT_REGION` is set that will be used
- else if profile in `$HOME/.aws` is set, then use that
- Otherwise it will be empty and program will panic if flag not set.

```shell
# flag -region is optional if above rules are met
aws-log-deleter -region=us-east-2
```
