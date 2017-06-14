package rpc

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/negz/practice/tiny/proto"
	"github.com/negz/practice/tiny/url"
	"github.com/pkg/errors"
)

type shortenerServer struct {
	sh url.Shortener
}

func newShortenerServer(sh url.Shortener) (proto.ShortenerServer, error) {
	return &shortenerServer{sh}, nil
}

func (s *shortenerServer) Get(_ context.Context, r *proto.GetRequest) (*proto.GetResponse, error) {
	url, err := s.sh.Get(r.GetPath())
	if IsNotFound(err) {
		return &proto.GetResponse{URL: url, Status: Status(Codify(err, codes.NotFound))}, nil
	}
	return &proto.GetResponse{URL: url, Status: Status(err)}, nil
}

func (s *shortenerServer) Create(_ context.Context, r *proto.CreateRequest) (*proto.CreateResponse, error) {
	path, err := s.sh.Create(r.GetURL())
	return &proto.CreateResponse{Path: path, Status: Status(err)}, nil
}

// A Server serves gRPC requests to get and create short URLs.
type Server struct {
	l  net.Listener
	sh url.Shortener
}

// NewServer returns a new gRPC server.
func NewServer(l net.Listener, sh url.Shortener) (*Server, error) {
	return &Server{l: l, sh: sh}, nil
}

// Serve begins serving gRPC requests.
func (s *Server) Serve() error {
	g := grpc.NewServer()
	shs, err := newShortenerServer(s.sh)
	if err != nil {
		return errors.Wrap(err, "cannot create gRPC shortener server")
	}
	proto.RegisterShortenerServer(g, shs)
	return errors.Wrap(g.Serve(s.l), "cannot serve gRPC requests")
}
