package main

import (
	"net"
	"os"
	"path/filepath"

	"github.com/negz/practice/tiny/rpc"
	"github.com/negz/practice/tiny/url/b62generator"
	"github.com/negz/practice/tiny/url/mapstore"
	"github.com/negz/practice/tiny/url/shortener"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app    = kingpin.New(filepath.Base(os.Args[0]), "Create and lookup short URLs!")
		listen = app.Flag("listen", "Address at which to listen for gRPC connections.").Default(":10002").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	g, err := b62generator.New()
	kingpin.FatalIfError(err, "cannot create base62 URL generator")

	st, err := mapstore.New()
	kingpin.FatalIfError(err, "cannot create map-based URL store")

	sh, err := shortener.New(g, st)
	kingpin.FatalIfError(err, "cannot create URL shortener")

	l, err := net.Listen("tcp", *listen)
	kingpin.FatalIfError(err, "cannot listen on requested address")

	srv, err := rpc.NewServer(l, sh)
	kingpin.FatalIfError(err, "cannot create gRPC server")

	kingpin.FatalIfError(srv.Serve(), "gRPC server error")
}
