package address

import "google.golang.org/grpc"

type Client struct {
	client AddressManagerClient
}

func New(conn *grpc.ClientConn) *Client {
	cli := Client{}
	cli.client = NewAddressManagerClient(conn)
	return &cli
}
