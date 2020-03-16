go test -v -race -run=Test memo_test.go
go test -v -race -run=TestConcurrent memo_test.go
go test -v -race -run=TestConcurrentCancel memo_test.go
