package main

import (
	"context"
	"net/http"
	"os"

	"google.golang.org/grpc"

	"fmt"
	"io/ioutil"
	"net"

	"github.com/backstopmedia/gRPC-book-example/server/api"
	"github.com/backstopmedia/gRPC-book-example/server/db"
	pb "github.com/backstopmedia/gRPC-book-example/server/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

const (
	defaultDataFile = "data.json"
	defaultPort     = 8080
)

func main() {

	app := cli.NewApp()
	app.Name = "grpc-example"
	app.Commands = []cli.Command{
		grpcServerCmd(),
		gatewayServerCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func grpcServerCmd() cli.Command {
	return cli.Command{
		Name:  "grpc-server",
		Usage: "starts a gRPC server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: defaultPort,
			},
			cli.StringFlag{
				Name:  "file",
				Value: defaultDataFile,
			},
		},
		Action: func(c *cli.Context) error {
			port, file := c.Int("port"), c.String("file")

			logrus.Printf("Starting RPC server on port %d...", port)
			list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return err
			}

			data, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			db, err := db.NewJSONProvider(data)
			if err != nil {
				return err
			}

			interceptorOpt := grpc.UnaryInterceptor(api.Interceptors())
			s := grpc.NewServer(interceptorOpt)
			pb.RegisterStarfriendsServer(s, api.New(db))

			if err := s.Serve(list); err != nil {
				return err
			}

			return nil
		},
	}
}

func gatewayServerCmd() cli.Command {
	return cli.Command{
		Name:  "gw-server",
		Usage: "starts a Gateway server for Starfriends",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 8081,
			},
			cli.StringFlag{
				Name:  "grpc-endpoint",
				Value: "localhost:8080",
				Usage: "the address of the running gRPC server to transcode to",
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithInsecure()}
			err := pb.RegisterStarfriendsHandlerFromEndpoint(ctx, mux, c.String("grpc-endpoint"), opts)
			if err != nil {
				return err
			}

			logrus.Printf("Starting JSON Gateway server on port %d...", c.Int("port"))
			return http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), mux)
		},
	}
}
