package cli

import (
	"flag"
	"fmt"
)

type ConnectionUrl struct {
	host string
	port int32
}

func newConnectionUrl(host string, port int) *ConnectionUrl {
	return &ConnectionUrl{host: host, port: int32(port)}
}

func (cu *ConnectionUrl) Url() string {
	return fmt.Sprintf("%s:%d", cu.host, cu.port)
}

func GetConnectionData() *ConnectionUrl {
	port := flag.Int("port", 8000, "-p PORT")
	host := flag.String("host", "localhost", "-h HOST")
	flag.Parse()

	conn := newConnectionUrl(*host, *port)
	return conn
}
