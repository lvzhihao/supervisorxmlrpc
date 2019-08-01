package supervisorxmlrpc

const (
	MethodListMethods     string = "listMethods"
	MethodMethodHelp      string = "methodHelp"
	MethodMethodSignature string = "methodSignature"
	MethodMulticall       string = "multicall"
)

type MulticallParams struct {
	MethodName string        `xmlrpc:"methodName"`
	Params     []interface{} `xmlrpc:"params"`
}

// TODO
type MulticallReturnFailure struct {
	FaultCode   int    `xmlrpc:"faultCode"`
	FaultString string `xmlrpc:"faultString"`
}

type SystemApiInterface interface {
	ListMethods() ([]string, error)
	MethodHelp(string) ([]string, error)
	MethodSignature(string) ([]string, error)
	Multicall([]MulticallParams) ([]interface{}, error)
}

func (c *Client) ListMethods() ([]string, error) {
	return c.CallStringList(GetSystemMethod(MethodListMethods), nil...)
}

func (c *Client) MethodHelp(name string) (string, error) {
	return c.CallString(GetSystemMethod(MethodMethodHelp), name)
}

func (c *Client) MethodSignature(name string) ([]string, error) {
	return c.CallStringList(GetSystemMethod(MethodMethodSignature), name)
}

func (c *Client) Multicall(calls []MulticallParams) (ret []interface{}, err error) {
	ret = make([]interface{}, 0)
	err = c.CallStruct(GetSystemMethod(MethodMulticall), &ret, calls)
	return
}

func ListMethods(client *Client) ([]string, error) {
	return client.ListMethods()
}

func MethodHelp(client *Client, name string) (string, error) {
	return client.MethodHelp(name)
}

func MethodSignature(client *Client, name string) ([]string, error) {
	return client.MethodSignature(name)
}

func Multicall(client *Client, calls []MulticallParams) (ret []interface{}, err error) {
	return client.Multicall(calls)
}
