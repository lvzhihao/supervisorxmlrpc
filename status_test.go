package supervisorxmlrpc

import (
	"bytes"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestGetSupervisorVersion(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getSuperVisorVersion",
			args: args{
				client: newConnect(t, testServer),
			},
			want: func() string {
				cmd := exec.Command(testSupervisorctl, append(testCmds, "version")...)
				var stdout bytes.Buffer
				cmd.Stdout = &stdout
				err := cmd.Run()
				if err != nil {
					return err.Error()
				} else {
					return strings.TrimSpace(stdout.String())
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSupervisorVersion(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSupervisorVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSupervisorVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAPIVersion(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getAPIVersion",
			args: args{
				client: newConnect(t, testServer),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIVersion(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				// don't check
				//t.Errorf("GetAPIVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIdentification(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getIdentification",
			args: args{
				client: newConnect(t, testServer),
			},
			want: "sup.exapmle.single",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIdentification(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIdentification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetState(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		wantRet *GetStatusReturn
		wantErr bool
	}{
		{
			name: "getState",
			args: args{
				client: newConnect(t, testServer),
			},
			wantRet: &GetStatusReturn{
				StateCode: StateCodeRunning,
				StateName: StateNameRunning,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := GetState(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("GetState() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestGetPID(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "getPID",
			args: args{
				client: newConnect(t, testServer),
			},
			want: func() int {
				cmd := exec.Command(testSupervisorctl, append(testCmds, "pid")...)
				var stdout bytes.Buffer
				cmd.Stdout = &stdout
				err := cmd.Run()
				if err != nil {
					t.Log(err)
					return 0
				} else {
					ret, _ := strconv.Atoi(strings.TrimSpace(stdout.String()))
					return ret
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPID(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLog(t *testing.T) {
	type args struct {
		client *Client
		offset int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "readLog",
			args: args{
				client: newConnect(t, testServer),
				offset: 0,
				length: 100,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLog(tt.args.client, tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearLog(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "clearLog",
			args: args{
				client: newConnect(t, testServer),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClearLog(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClearLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClearLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShutdown(t *testing.T) {
	// manaul test
	return
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "shutdown",
			args: args{
				client: newConnect(t, testServer),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Shutdown(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shutdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestart(t *testing.T) {
	type args struct {
		client *Client
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "restart",
			args: args{
				client: newConnect(t, testServer),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Restart(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Restart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Restart() = %v, want %v", got, tt.want)
			}
		})
	}
}
