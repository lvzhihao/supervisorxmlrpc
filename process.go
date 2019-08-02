package supervisorxmlrpc

const (
	MethodGetProcessInfo      string = "getProcessInfo"
	MethodGetAllProcessInfo   string = "getAllProcessInfo"
	MethodStartProcess        string = "startProcess"
	MethodStartAllProcess     string = "startAllProcesses"
	MethodStartProcessGroup   string = "startProcessGroup"
	MethodStopProcess         string = "stopProcess"
	MethodStopAllProcess      string = "stopAllProcesses"
	MethodStopProcessGroup    string = "stopProcessGroup"
	MethodSignalProcess       string = "signalProcess"
	MethodSignalAllProcess    string = "signalAllProcesses"
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
	StartAllProcess(bool) ([]ProcessInfoReturn, error)
	StartProcessGroup(string, bool) ([]ProcessInfoReturn, error)
	StopProcess(string, bool) (bool, error)
	StopAllProcesses(bool) ([]ProcessInfoReturn, error)
	StopProcessGroup(string, bool) ([]ProcessInfoReturn, error)
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

func (c *Client) StartAllProcess(wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStartAllProcess), &ret, wait)
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

func (c *Client) StopAllProcess(wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStopAllProcess), &ret, wait)
	return
}

func (c *Client) StopProcessGroup(group string, wait bool) (ret []ProcessInfoReturn, err error) {
	ret = make([]ProcessInfoReturn, 0)
	err = c.CallStruct(GetSupervisorMethod(MethodStopProcessGroup), &ret, group, wait)
	return
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

func StartAllProcess(client *Client, wait bool) ([]ProcessInfoReturn, error) {
	return client.StartAllProcess(wait)
}

func StartProcessGroup(client *Client, group string, wait bool) ([]ProcessInfoReturn, error) {
	return client.StartProcessGroup(group, wait)
}

func StopProcess(client *Client, name string, wait bool) (bool, error) {
	return client.StopProcess(name, wait)
}

func StopAllProcess(client *Client, wait bool) ([]ProcessInfoReturn, error) {
	return client.StopAllProcess(wait)
}

func StopProcessGroup(client *Client, group string, wait bool) ([]ProcessInfoReturn, error) {
	return client.StopProcessGroup(group, wait)
}
