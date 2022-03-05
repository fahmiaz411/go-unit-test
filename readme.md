go test -v -bench=.
go test -v -run=NotMathUnitTest -bench=.
go test -v -run=NotMathUnitTest -bench=BenchmarkXxx
go test -v -bench=B ./...
go test -v -bench=BenchmarkXxx/sub
