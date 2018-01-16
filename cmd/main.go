package main

import (
	"flag"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/heroku/silvia-runtime-university/client"
	"github.com/heroku/silvia-runtime-university/spec"
	"github.com/joeshaw/envdecode"
)

func main() {
	var serverUrlFlag = flag.String("server-url", "", "Url of server to connect to")
	flag.Parse()

	type Config struct {
		ServerUrl string `env:"SERVER_URL,default=grpc-server.herokuapp.com:80"`
	}

	var cfg Config
	err := envdecode.Decode(&cfg)
	serverAddress := cfg.ServerUrl
	if *serverUrlFlag != "" {
		serverAddress = *serverUrlFlag
	}

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer conn.Close()

	c := client.RouteGuide{ Client: spec.NewRouteGuideClient(conn)}

	var inputPoints = []spec.Point{
		{ Latitude: 407838351, Longitude: -746143763 },
		{ Latitude: 408122808, Longitude: -743999179 },
		{ Latitude: 413628156, Longitude: -749015468 },
	}

	features, err := c.GetFeatures(context.Background(), inputPoints)
	if err != nil {
		log.Fatalf("Fetching features failed: %v", err)
	}

	log.Println("Features for the given points: ", features)
}
