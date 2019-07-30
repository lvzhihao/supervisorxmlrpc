package supervisorxmlrpc

import (
	"reflect"
	"testing"
)

var (
	testServer string   = "http://demo:demo@127.0.0.1:9015/RPC2"
	testCmds   []string = []string{
		"-c", "examples/single/supervisord.conf",
		"-u", "demo",
		"-p", "demo",
	}
	testSupervisord   string = "/usr/bin/supervisord"
	testSupervisorctl string = "/usr/bin/supervisorctl"
)

func newConnect(t *testing.T, server string) *Client {
	client, err := Connect(server)
	if err != nil {
		t.Fatalf("Connect Error: %v", err)
	}
	return client
}

func TestClientCallString(t *testing.T) {
	type args struct {
		server string
		method string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "single/getIdentificatio",
			args: args{
				server: testServer,
				method: "supervisor.getIdentification",
				args:   nil,
			},
			want: "sup.exapmle.single",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newConnect(t, tt.args.server)
			got, err := client.CallString(tt.args.method, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientCallBool(t *testing.T) {
	type args struct {
		server string
		method string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "single/clearLog",
			args: args{
				server: testServer,
				method: "supervisor.clearLog",
				args:   nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newConnect(t, tt.args.server)
			got, err := client.CallBool(tt.args.method, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientCallStringList(t *testing.T) {
	type args struct {
		server string
		method string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "single/listMethods",
			args: args{
				server: testServer,
				method: "system.listMethods",
				args:   nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newConnect(t, tt.args.server)
			got, err := client.CallStringList(tt.args.method, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallStringList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			if len(got) <= tt.want {
				t.Errorf("CallStringList() = %v, want %v", got, tt.want)
			}
		})
	}
}
