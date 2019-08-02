package supervisorxmlrpc

const (
	MethodReadProcessStdoutLog string = "readProcessStdoutLog"
	MethodReadProcessStderrLog string = "readProcessStderrLog"
	MethodTailProcessStdoutLog string = "tailProcessStdoutLog"
	MethodTailProcessStderrLog string = "tailProcessStderrLog"
	MethodClearProcessLogs     string = "clearProcessLogs"
	MethodClearAllProcessLogs  string = "clearAllProcessLogs"
)

type LoggingApiInterface interface {
	ReadProcessStdoutLog(string, int, int) (string, error)
	TailProcessStdoutLog(string, int, int) (string, error)
	ReadProcessStderrLog(string, int, int) ([]interface{}, error)
	TailProcessStderrLog(string, int, int) (string, error)
	ClearProcessLogs(string) (bool, error)
	ClearAllProcessLogs() ([]ProcessInfoReturn, error)
}

func (c *Client) ReadProcessStdoutLog(name string, offset, length int) (string, error) {
	return c.CallString(GetSupervisorMethod(MethodReadProcessStdoutLog), name, offset, length)
}

func (c *Client) ReadProcessStderrLog(name string, offset, length int) (string, error) {
	return c.CallString(GetSupervisorMethod(MethodReadProcessStderrLog), name, offset, length)
}

func (c *Client) TailProcessStdoutLog(name string, offset, length int) (ret []interface{}, err error) {
	ret = make([]interface{}, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodTailProcessStdoutLog), &ret, name, offset, length)
	return
}

func (c *Client) TailProcessStderrLog(name string, offset, length int) (ret []interface{}, err error) {
	ret = make([]interface{}, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodTailProcessStderrLog), &ret, name, offset, length)
	return
}

func (c *Client) ClearProcessLogs(name string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodClearProcessLogs), name)
}

func (c *Client) ClearAllProcessLogs() (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodClearAllProcessLogs), &ret, nil...)
	return
}

func ReadProcessStdoutLog(client *Client, name string, offset, length int) (string, error) {
	return client.ReadProcessStdoutLog(name, offset, length)
}

func ReadProcessStderrLog(client *Client, name string, offset, length int) (string, error) {
	return client.ReadProcessStderrLog(name, offset, length)
}

func TailProcessStdoutLog(client *Client, name string, offset, length int) ([]interface{}, error) {
	return client.TailProcessStdoutLog(name, offset, length)
}

func TailProcessStderrLog(client *Client, name string, offset, length int) ([]interface{}, error) {
	return client.TailProcessStderrLog(name, offset, length)
}

func ClearProcessLogs(client *Client, name string) (bool, error) {
	return client.ClearProcessLogs(name)
}

func ClearAllProcessLogs(client *Client) ([]ProcessInfoReturn, error) {
	return client.ClearAllProcessLogs()
}
