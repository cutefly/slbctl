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
go run . config login --username=username --password=password
go run . config server --url https://132b1854-53f1-489b-9152-ac9dad68bdcb.mock.pstmn.io
go run . show member kubernetes-dev-32443-gr
go run . add member kubernetes-dev-32443-gr kubernetes-dev-32443-3
go run . remove member kubernetes-dev-32443-gr kubernetes-dev-32443-3 --force=true
```

## Build

## go build

```sh
go mod tidy
go build
./slbctl config login --username=username --password=password
```
