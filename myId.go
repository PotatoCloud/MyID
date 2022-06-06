package main

import "flag"

var (
	serverAddr = flag.String("addr", "0.0.0.0:4000", "MyID server listen address")
)

func main() {
	if err := Run(*serverAddr); err != nil {
		return
	}
}
