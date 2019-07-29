package supervisorxmlrpc

import (
	"net/http"

	"github.com/kolo/xmlrpc"
)

const (
	DefaultMethodPrefix string = "supervisor."
)

type Client struct {
	server    string            // server url http:// or unix:///
	connect   *xmlrpc.Client    // third part xmlrpc client
	transport http.RoundTripper // TODO if unxi:///
}

func GetRealMethod(method string) string {
	return DefaultMethodPrefix + method
}

func Connect(server string) (*Client, error) {
	c := &Client{
		server:    server,
		transport: nil,
	}
	connect, err := xmlrpc.NewClient(c.GetServer(), c.GetTransport())
	if err != nil {
		return nil, err
	} else {
		c.SetConnect(connect)
		return c, nil
	}
}

func (c *Client) SetConnect(connect *xmlrpc.Client) {
	c.connect = connect
}

func (c *Client) GetConnect() *xmlrpc.Client {
	return c.connect
}

func (c *Client) SetServer(server string) {
	c.server = server
}

func (c *Client) GetServer() string {
	return c.server
}

func (c *Client) SetTransport(transport http.RoundTripper) {
	c.transport = transport
}

func (c *Client) GetTransport() http.RoundTripper {
	return c.transport
}

func (c *Client) CallString(method string, args ...interface{}) (rst string, err error) {
	err = c.connect.Call(method, args, &rst)
	return
}

func (c *Client) CallBool(method string, args ...interface{}) (rst bool, err error) {
	err = c.connect.Call(method, args, &rst)
	return
}
