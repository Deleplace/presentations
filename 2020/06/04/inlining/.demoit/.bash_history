curl -s https://golang.org/ref/spec | grep -i inlin
curl -s https://golang.org/ref/spec | grep -i implementation-dependent
go test -gcflags -S
go test -gcflags="-m -m" |& grep f1
go test -gcflags="-m -m"
go test -gcflags="-m" |& grep f1
go test -gcflags="-m"
go test -bench=.
