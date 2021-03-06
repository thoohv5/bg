package boot

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/thoohv5/gf/gin"
	"github.com/thoohv5/gf/grpc/server"
	"google.golang.org/grpc"

	"github.com/thoohv5/bg/route"
)

func Http() error {

	cer, err := tls.LoadX509KeyPair("/var/www/ca/5307072_www.thooh.com.pem", "/var/www/ca/5307072_www.thooh.com.key")
	if err != nil {
		return err
	}

	if err := server.NewServer().Server(&server.Config{
		Network:   "tcp",
		Address:   "0.0.0.0:80",
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cer}},
	}, func(gServer *grpc.Server) error {
		return nil
	}, func(ctx context.Context, mux *runtime.ServeMux, addr string, dialOption []grpc.DialOption) error {
		return nil
	}, func(mux *http.ServeMux) error {

		router, _ := gin.New().Handle(route.RegisterRoute)
		mux.Handle("/blog/", router)
		return nil
	}); nil != err {
		return err
	}

	return nil
}
