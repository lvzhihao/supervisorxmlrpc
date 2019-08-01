package supervisorxmlrpc

const (
	MethodGetProcessInfo    string = "getProcessInfo"
	MethodGetAllProcessInfo string = "getAllProcessInfo"
)

type ProcessApiInterface interface {
	GetProcessInfo(string) (GetProcessInfoReturn, error)
	GetAllProcessInfo() ([]GetProcessInfoReturn, error)
}

type GetProcessInfoReturn struct {
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

func (c *Client) GetProcessInfo(name string) (ret *GetProcessInfoReturn, err error) {
	ret = new(GetProcessInfoReturn)
	err = c.CallStruct(GetSupervisorMethod(MethodGetProcessInfo), ret, name)
	return
}

func (c *Client) GetAllProcessInfo() (ret []GetProcessInfoReturn, err error) {
	err = c.CallStruct(GetSupervisorMethod(MethodGetAllProcessInfo), &ret, nil...)
	return
}

func GetProcessInfo(client *Client, name string) (*GetProcessInfoReturn, error) {
	return client.GetProcessInfo(name)
}

func GetAllProcessInfo(client *Client) ([]GetProcessInfoReturn, error) {
	return client.GetAllProcessInfo()
}
