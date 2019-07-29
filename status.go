package supervisorxmlrpc

const (
	MethodGetApiVersion        string = "getAPIVersion"
	MethodGetSupervisorVersion string = "getSupervisorVersion"
)

func GetAPIVersion(client *Client) (string, error) {
	return client.CallString(GetRealMethod(MethodGetApiVersion), nil)
}

func GetSupervisorVersion(client *Client) (string, error) {

	return client.CallString(GetRealMethod(MethodGetSupervisorVersion), nil)
}
