package main

import (
    "cmux_practice/internal/pb"
    "fmt"
    "github.com/soheilhy/cmux"
    "golang.org/x/net/context"
    "golang.org/x/sync/errgroup"
    "google.golang.org/grpc"
    "log"
    "net"
    "net/http"
)

type grpcServer struct{}

func main() {
    listener, err := net.Listen("tcp", "localhost:23456")
    if err != nil {
        log.Fatal(err)
    }

    m := cmux.New(listener)
    httpListener := m.Match(cmux.HTTP1Fast())
    //grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
    grpcListener := m.Match(cmux.Any())

    g := new(errgroup.Group)
    g.Go(func() error { return grpcServe(grpcListener) })
    g.Go(func() error { return httpServe(httpListener) })
    g.Go(func() error { return m.Serve() })

    log.Println("run server: ", g.Wait())
}

func httpServe(l net.Listener) error {
    mux := http.NewServeMux()
    mux.HandleFunc("/ping/", func(w http.ResponseWriter, _ *http.Request) {
        w.Write([]byte("pong"))
    })

    s := &http.Server{Handler: mux}
    return s.Serve(l)
}

func (*grpcServer) Say(ctx context.Context, in *pb.Req) (*pb.Res, error) {
    return &pb.Res{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func grpcServe(l net.Listener) error {
    s := grpc.NewServer()
    pb.RegisterHelloWorldServer(s, &grpcServer{})

    return s.Serve(l)
}