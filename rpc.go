package supervisorxmlrpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/kolo/xmlrpc"
)

const (
	DefaultSupervisorMethodPrefix string = "supervisor."
	DefaultSystemMethodPrefix     string = "system."
)

type Client struct {
	server    string            // server url http:// or unix://
	connect   *xmlrpc.Client    // third part xmlrpc client
	transport http.RoundTripper // if unxi://
}

func GetSupervisorMethod(method string) string {
	return DefaultSupervisorMethodPrefix + method
}

func GetSystemMethod(method string) string {
	return DefaultSystemMethodPrefix + method
}

func Connect(server string) (*Client, error) {
	u, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http":
		return connectHttp(u, nil)
	case "unix":
		return connectUnix(u)
	default:
		return nil, fmt.Errorf("url error: %s", server)
	}
}

func connectHttp(u *url.URL, tr http.RoundTripper) (*Client, error) {
	c := &Client{
		server:    u.String(),
		transport: tr,
	}
	connect, err := xmlrpc.NewClient(c.GetServer(), c.GetTransport())
	if err != nil {
		return nil, err
	} else {
		c.SetConnect(connect)
		return c, nil
	}
}

func connectUnix(u *url.URL) (*Client, error) {
	path := u.Path
	transport := &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", path)
		},
	}
	// change scheme && urlpath
	u.Scheme = "http"
	u.Path = "unix/RPC2"
	return connectHttp(u, transport)
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

func (c *Client) CallInt(method string, args ...interface{}) (rst int, err error) {
	err = c.connect.Call(method, args, &rst)
	return
}

func (c *Client) CallStringList(method string, args ...interface{}) (rst []string, err error) {
	err = c.connect.Call(method, args, &rst)
	return
}

func (c *Client) CallStruct(method string, v interface{}, args ...interface{}) (err error) {
	err = c.connect.Call(method, args, v)
	return
}
