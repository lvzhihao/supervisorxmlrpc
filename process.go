package supervisorxmlrpc

const (
	MethodGetProcessInfo      string = "getProcessInfo"
	MethodGetAllProcessInfo   string = "getAllProcessInfo"
	MethodStartProcess        string = "startProcess"
	MethodStartAllProcesses   string = "startAllProcesses"
	MethodStartProcessGroup   string = "startProcessGroup"
	MethodStopProcess         string = "stopProcess"
	MethodStopAllProcesses    string = "stopAllProcesses"
	MethodStopProcessGroup    string = "stopProcessGroup"
	MethodSignalProcess       string = "signalProcess"
	MethodSignalAllProcesses  string = "signalAllProcesses"
	MethodSignalProcessGroup  string = "signalProcessGroup"
	MethodSendProcessStdin    string = "sendProcessStdin"
	MethodSendRemoteCommEvent string = "sendRemoteCommEvent"
	MethodReloadConfig        string = "reloadConfig"
	MethodAddProcessGroup     string = "addProcessGroup"
	MethodRemoveProcessGroup  string = "removeProcessGroup"
)

type ProcessApiInterface interface {
	GetProcessInfo(string) (ProcessInfoReturn, error)
	GetAllProcessInfo() ([]ProcessInfoReturn, error)
	StartProcess(string, bool) (bool, error)
	StartAllProcesses(bool) ([]ProcessInfoReturn, error)
	StartProcessGroup(string, bool) ([]ProcessInfoReturn, error)
	StopProcess(string, bool) (bool, error)
	StopAllProcesses(bool) ([]ProcessInfoReturn, error)
	StopProcessGroup(string, bool) ([]ProcessInfoReturn, error)
	SignalProcess(string, string) (bool, error)
	SignalProcessGroup(string, string) ([]ProcessInfoReturn, error)
	SignalAllProcesses(string) ([]ProcessInfoReturn, error)
	SendProcessStdin(string, string) (bool, error)
	SendRemoteCommEvent(string, string) (bool, error)
	ReloadConfig() ([]interface{}, error)
	AddProcessGroup(string) (bool, error)
	RemoveProcessGroup(string) (bool, error)
}

type ProcessInfoReturn struct {
	Name          string `xmlrpc:"name"`
	Group         string `xmlrpc:"group"`
	Description   string `xmlrpc:"description"`
	Start         int    `xmlrpc:"start"`
	Stop          int    `xmlrpc:"stop"`
	Now           int    `xmlrpc:"now"`
	State         int    `xmlrpc:"state"`
	StateName     string `xmlrpc:"statename"`
	Status        int    `xmlrpc:"status"`
	SpawnErr      string `xmlrpc:"spawnerr"`
	ExitStatus    int    `xmlrpc:"exitstatus"`
	LogFile       string `xmlrpc:"logfile"`
	StdoutLogFile string `xmlrpc:"stdout_logfile"`
	StderrLogFile string `xmlrpc:"stderr_logfile"`
	Pid           int    `xmlrpc:"pid"`
}

func (c *Client) GetProcessInfo(name string) (ret *ProcessInfoReturn, err error) {
	ret = new(ProcessInfoReturn)
	err = c.CallStruct(GetSupervisorMethod(MethodGetProcessInfo), ret, name)
	return
}

func (c *Client) GetAllProcessInfo() (ret []ProcessInfoReturn, err error) {
	err = c.CallStruct(GetSupervisorMethod(MethodGetAllProcessInfo), &ret, nil...)
	return
}

func (c *Client) StartProcess(name string, wait bool) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodStartProcess), name, wait)
}

func (c *Client) StartAllProcesses(wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStartAllProcesses), &ret, wait)
	return
}

func (c *Client) StartProcessGroup(group string, wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStartProcessGroup), &ret, group, wait)
	return
}

func (c *Client) StopProcess(name string, wait bool) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodStopProcess), name, wait)
}

func (c *Client) StopAllProcesses(wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStopAllProcesses), &ret, wait)
	return
}

func (c *Client) StopProcessGroup(group string, wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStopProcessGroup), &ret, group, wait)
	return
}

func (c *Client) SignalProcess(name, signal string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodSignalProcess), name, signal)
}

func (c *Client) SignalProcessGroup(name, signal string) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodSignalProcessGroup), &ret, name, signal)
	return
}

func (c *Client) SignalAllProcesses(signal string) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodSignalAllProcesses), &ret, signal)
	return
}

func (c *Client) SendProcessStdin(name, chars string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodSendProcessStdin), name, chars)
}

func (c *Client) SendRemoteCommEvent(name, data string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodSendRemoteCommEvent), name, data)
}

func (c *Client) ReloadConfig() (ret []interface{}, err error) {
	ret = make([]interface{}, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodReloadConfig), &ret, nil...)
	return
}

func (c *Client) AddProcessGroup(name string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodAddProcessGroup), name)
}

func (c *Client) RemoveProcessGroup(name string) (bool, error) {
	return c.CallBool(GetSupervisorMethod(MethodRemoveProcessGroup), name)
}

func GetProcessInfo(client *Client, name string) (*ProcessInfoReturn, error) {
	return client.GetProcessInfo(name)
}

func GetAllProcessInfo(client *Client) ([]ProcessInfoReturn, error) {
	return client.GetAllProcessInfo()
}

func StartProcess(client *Client, name string, wait bool) (bool, error) {
	return client.StartProcess(name, wait)
}

func StartAllProcesses(client *Client, wait bool) ([]ProcessInfoReturn, error) {
	return client.StartAllProcesses(wait)
}

func StartProcessGroup(client *Client, group string, wait bool) ([]ProcessInfoReturn, error) {
	return client.StartProcessGroup(group, wait)
}

func StopProcess(client *Client, name string, wait bool) (bool, error) {
	return client.StopProcess(name, wait)
}

func StopAllProcesses(client *Client, wait bool) ([]ProcessInfoReturn, error) {
	return client.StopAllProcesses(wait)
}

func StopProcessGroup(client *Client, group string, wait bool) ([]ProcessInfoReturn, error) {
	return client.StopProcessGroup(group, wait)
}

func SignalProcess(client *Client, name, signal string) (bool, error) {
	return client.SignalProcess(name, signal)
}

func SignalProcessGroup(client *Client, name, signal string) ([]ProcessInfoReturn, error) {
	return client.SignalProcessGroup(name, signal)
}

func SignalAllProcesses(client *Client, signal string) ([]ProcessInfoReturn, error) {
	return client.SignalAllProcesses(signal)
}

func SendProcessStdin(client *Client, name, chars string) (bool, error) {
	return client.SendProcessStdin(name, chars)
}

func SendRemoteCommEvent(client *Client, name, data string) (bool, error) {
	return client.SendRemoteCommEvent(name, data)
}

func ReloadConfig(client *Client) ([]interface{}, error) {
	return client.ReloadConfig()
}

func AddProcessGroup(client *Client, name string) (bool, error) {
	return client.AddProcessGroup(name)
}

func RemoveProcessGroup(client *Client, name string) (bool, error) {
	return client.RemoveProcessGroup(name)
}
