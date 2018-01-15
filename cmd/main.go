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
	flag.String("help", "", "Display usage")
	flag.Parse()

	type Config struct {
		ServerUrl string `env:"SERVER_URL,default=grpc-server.herokuapp.com:80"`
	}

	var cfg Config
	err := envdecode.Decode(&cfg)

	conn, err := grpc.Dial(cfg.ServerUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer conn.Close()

	c := client.RouteGuide{ Client: spec.NewRouteGuideClient(conn)}

	var inputPoints = []spec.Point{
		{ Latitude: 98349834, Longitude: 384738473 },
		{ Latitude: 98349835, Longitude: 384738474 },
		{ Latitude: 98349836, Longitude: 384738475 },
	}

	features, err := c.GetFeatures(context.Background(), inputPoints)
	if err != nil {
		log.Fatalf("Fetching features failed: %v", err)
	}

	log.Println("Features for the given points: ", features)
}
