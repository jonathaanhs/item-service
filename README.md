# Item Service

## Requirements
- Go 1.19

## Quick Start
```bash
make unit-test # to run unit test
make run # run the app, running on localhost:8089
make build  # build the app, bin/tokenomy-assessment
```

## Call API
```bash
Case 1: Request without parameter id

curl --location --request GET 'localhost:8089/?id='
{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}]}
```

```bash
Case 2: Request with single id

curl --location --request GET 'localhost:8089/?id=2'
{"code":200,"data":[{"id":2,"name":"B"}]}
```

```bash
Case 3: Request with multiple ids

curl --location --request GET 'localhost:8089/?id=1,3,4'
{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}]}
```

```bash
Case 4: Request with invalid ID

curl --location --request GET 'localhost:8089/?id=xxx'
{"StatusCode":400,"Message":"invalid or empty ID: \"xxx\""}
```

```bash
Case 5: Request with ID not found

curl --location --request GET 'localhost:8089/?id=4'
{"StatusCode":404,"Message":"resource with ID 4 not exist"}
```

## Unit Test Coverage

```console
foo@bar:~$ make unit-test
go clean -testcache
go test ./... --cover
?       github.com/tokenomy-assessment/cmd      [no test files]
?       github.com/tokenomy-assessment/internal [no test files]
ok      github.com/tokenomy-assessment/internal/controller      0.192s  coverage: 100.0% of statements
ok      github.com/tokenomy-assessment/internal/service 0.097s  coverage: 100.0% of statements
ok      github.com/tokenomy-assessment/pkg/httpkit      0.140s  coverage: 100.0% of statements
```
