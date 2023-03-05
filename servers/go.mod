module github.com/nohns/semesterprojekt2

go 1.19

replace github.com/nohns/semesterprojekt2/proto => ../proto

require (
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/nohns/semesterprojekt2/proto v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)
