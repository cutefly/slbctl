# SLB cli

## Install

## install go

```

```

## cobra library

```sh
go mod init slbctl

cobra-cli init
cobra-cli add add
cobra-cli add config
cobra-cli add login
...
```

## Run

## go run

```sh
go run . config login --user=username --password=password
go run . config server --url https://132b1854-53f1-489b-9152-ac9dad68bdcb.mock.pstmn.io --skip-verify=true
go run . show member kubernetes-dev-32443-gr
go run . add member kubernetes-dev-32443-gr kubernetes-dev-32443-3
go run . remove member kubernetes-dev-32443-gr kubernetes-dev-32443-3 --force=true
```

## environment variables

```
username : SLBCTL_USERNAME
password : SLBCTL_PASSWORD
url : SLBCTL_URL
skip-verify : SLBCTL_SKIP_VERIFY(동작안함)
debug : SLBCTL_DEBUG
```

## Build

## go build

```sh
go mod tidy
go build
./slbctl config login --user=username --password=password
```

