package main

import (
	_ "github.com/gliderlabs/logspout/adapters/raw"
	_ "github.com/gliderlabs/logspout/adapters/syslog"
	_ "github.com/gliderlabs/logspout/transports/udp"
)
