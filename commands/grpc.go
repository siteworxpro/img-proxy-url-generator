package commands

import (
	"fmt"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	proto "github.com/siteworxpro/img-proxy-url-generator/grpc"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GrpcCommand() *cli.Command {
	return &cli.Command{
		Name:  "grpc",
		Usage: "Start a grpc service",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to listen on",
				Required: false,
				Value:    9000,
			},
		},
		Action: func(c *cli.Context) error {
			cfg, err := config.NewConfig(c.String("config"))
			if err != nil {
				return err
			}

			s := grpc.NewServer()
			addr := fmt.Sprintf(":%d", c.Int("port"))
			println("listening on", addr)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			svc, err := proto.NewService(cfg)
			if err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

			proto.RegisterGeneratorServer(s, svc)
			err = s.Serve(lis)
			if err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

			return nil
		},
	}
}
