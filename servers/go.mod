module github.com/nohns/semesterprojekt2

go 1.19

replace github.com/nohns/semesterprojekt2/proto => ../proto

require (
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/nohns/semesterprojekt2/proto v0.0.0-20230305122051-9285af5f9baf
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
)
