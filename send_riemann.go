package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/amir/raidman"
)

const (
	defaultRiemannServer = "localhost:5555"
	defaultEventTTL      = 10
)

var (
	riemannServer = flag.String("server", defaultRiemannServer, "Riemann server (host:port)")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `send_riemann - Send custom events to Riemann
usage:
  send_riemann [options] <JSON-formatted Riemann event>

Riemann event JSON spec:
  {
    "host": "hostname (default: local hostname)",
    "service": "service",
    "time": <UNIX epoch timestamp (default: now)>,
    "metric": <numeric value>,
    "state": "state (ok, warning, critical)",
    "ttl": <numeric value (seconds, default 10)>,
    "description": "some event description...",
    "tags": [ "tag1", "tag2", "tagN" ],
    "attributes": { "key1": "val1", "key2": "val2" }
  }

options:
`)
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	var (
		err      error
		hostname string
		client   *raidman.Client
		event    = raidman.Event{Ttl: defaultEventTTL,
			Host: hostname,
			Time: time.Now().Unix()}
	)

	if len(flag.Args()) == 0 {
		fmt.Println("error: no events supplied")
		flag.Usage()
	}

	if hostname, err = os.Hostname(); err != nil {
		die(fmt.Errorf("unable to get system hostname: %s\n", err))
	}

	if err := json.Unmarshal([]byte(flag.Arg(0)), &event); err != nil {
		die(fmt.Errorf("unable to parse event JSON data: %s\n", err))
	}

	if client, err = raidman.Dial("tcp", *riemannServer); err != nil {
		die(fmt.Errorf("unable to connect to riemann server: %s\n", err))
	}

	if err = client.Send(&event); err != nil {
		die(fmt.Errorf("unable to send event to Riemann server: %s\n", err))
	}

	client.Close()
}

func die(reason error) {
	fmt.Printf("error: %s", reason)
	os.Exit(1)
}
