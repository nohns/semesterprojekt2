package cloud

/* pb "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1"
pbConnect "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1/phonev1connect" */

//Server responbilbe for communication with the react native phone app

/* type server struct {
    pb.UnimplementedAppServiceHandler
}

func (s *server) GetLockState(ctx context.Context, in *pb.GetLockStateRequest) (*pb.GetLockStateResponse, error) {
}

func (s *server) SetLockState(ctx context.Context, in *pb.SetLockStateRequest) (*pb.SetLockStateResponse, error) {
} */



func Start() {

    //pb.UnimplementedAppServiceHandler()
	/* cloud := &server{}
    mux := http.NewServeMux()
    path, handler := pbConnect.NewAppServiceHandler(cloud)
    mux.Handle(path, handler)
    http.ListenAndServe(
        "localhost:8080",
        // Use h2c so we can serve HTTP/2 without TLS.
        h2c.NewHandler(mux, &http2.Server{}),
    ) */

}