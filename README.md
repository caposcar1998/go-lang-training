
# Exercises for go

## Chapter 3

Run on chapter3/src

``` bash
go run order.go
```

## Chapter 5

Run on root

```bash
go run .
```

Run on chapter5/src for testing

```bash
go test ./...
go test ./... -tags=integration
go test -bench=Benchmark
```

# Employees

- Run code

```bash
go run main.go
go run .
```


1. Add main module

```bash
go mod init main
go mod tidy
