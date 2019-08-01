package supervisorxmlrpc

const (
	MethodGetAPIVersion        string = "getAPIVersion"
	MethodGetSupervisorVersion string = "getSupervisorVersion"
	MethodGetIdentification    string = "getIdentification"
	MethodGetState             string = "getState"
	MethodGetPID               string = "getPID"
	MethodReadLog              string = "readLog"
	MethodClearLog             string = "clearLog"
	MethodShutdown             string = "shutdown"
	MethodRestart              string = "restart"
)

type StatusApiInterface interface {
	GetAPIVersion() (string, error)
	GetSupervisorVersion() (string, error)
	GetIdentification() (string, error)
	GetState() (GetStateReturn, error)
	GetPID() (int, error)
	ReadLog(int, int) (string, error)
	ClearLog() (bool, error)
	Shutdown() (bool, error)
	Restart() (bool, error)
}

type GetStateReturn struct {
	StateCode int    `xmlrpc:"statecode"`
	StateName string `xmlrpc:"statename"`
}

const (
	StateCodeShutDown   int    = -1 // SHUTDOWN: Supervisor is in the process of shutting down.
	StateCodeRestarting int    = 0  // RESTARTING: Supervisor is in the process of restarting.
	StateCodeRunning    int    = 1  // RUNNING: Supervisor is working normally.
	StateCodeFAtal      int    = 2  // FATAL: Supervisor has experienced a serious error.
	StateNameShowdown   string = "SHUTDOWN"
	StateNameRestarting string = "RESTARTING"
	StateNameRunning    string = "RUNNING"
	StateNameFatal      string = "FATAL"
)

func (c *Client) GetAPIVersion() (string, error) {
	return c.CallString(GetSupervisorMethod(MethodGetAPIVersion), nil...)
}

func (c *Client) GetSupervisorVersion() (string, error) {
	return c.CallString(GetSupervisorMethod(MethodGetSupervisorVersion), nil...)
}

func (c *Client) GetIdentification() (string, error) {
	return c.CallString(GetSupervisorMethod(MethodGetIdentification), nil...)
}

func (c *Client) GetState() (ret GetStateReturn, err error) {
	err = c.CallStruct(GetSupervisorMethod(MethodGetState), &ret, nil...)
	return
}

func (c *Client) GetPID() (int, error) {
	return c.CallInt(GetSupervisorMethod(MethodGetPID), nil...)
}

func (c *Client) ReadLog(offset, length int) (string, error) {
	return c.CallString(GetSupervisorMethod(MethodReadLog), offset, length)
}

func (c *Client) ClearLog() (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodClearLog), nil...)
}

func (c *Client) Shutdown() (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodShutdown), nil...)
}

func (c *Client) Restart() (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodRestart), nil...)
}

func GetAPIVersion(client *Client) (string, error) {
	return client.GetAPIVersion()
}

func GetSupervisorVersion(client *Client) (string, error) {
	return client.GetSupervisorVersion()
}

func GetIdentification(client *Client) (string, error) {
	return client.GetIdentification()
}

func GetState(client *Client) (GetStateReturn, error) {
	return client.GetState()
}

func GetPID(client *Client) (int, error) {
	return client.GetPID()
}

func ReadLog(client *Client, offset, length int) (string, error) {
	//TODO readlog return empty
	return client.ReadLog(offset, length)
}

func ClearLog(client *Client) (bool, error) {
	return client.ClearLog()
}

func Shutdown(client *Client) (bool, error) {
	return client.Shutdown()
}

func Restart(client *Client) (bool, error) {
	return client.Restart()
}
