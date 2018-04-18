# coverins

Coverage-instrumented binary creation in Go

## Usage

Execute:

```none
> make instrumented
go test -o ../../bin/covtest.coverage -c -covermode=count -coverpkg ./...

> make run
../../bin/covtest.coverage -test.coverprofile=sub.cov run_instrumented_binary
covtest.coverage
```

In other terminal:

```none
> make check
curl -X GET localhost:8080/two
Here's your number: 2
curl -X GET localhost:8080/text
Here's your string: TEST

> make show-coverage
go tool cover -html=sub.cov
```

Last command should open browser with stanard Go coverage infromation
