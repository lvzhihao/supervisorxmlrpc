package supervisorxmlrpc

import (
	"testing"
)

func TestGetProcessInfo(t *testing.T) {
	type args struct {
		client *Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    *ProcessInfoReturn
		wantErr bool
	}{
		{
			name: "getProcessInfo",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
			},
			want: &ProcessInfoReturn{
				Name: "single",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProcessInfo(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProcessInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			if got.Name != tt.want.Name {
				t.Errorf("GetProcessInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllProcessInfo(t *testing.T) {
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
			name: "getAllProcessInfo",
			args: args{
				client: newConnect(t, testServer),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllProcessInfo(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllProcessInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("GetAllProcessInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopProcess(t *testing.T) {
	type args struct {
		client *Client
		name   string
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "stopProcess",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				wait:   true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopProcess(tt.args.client, tt.args.name, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StopProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartProcess(t *testing.T) {
	type args struct {
		client *Client
		name   string
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "startProcess",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				wait:   true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StartProcess(tt.args.client, tt.args.name, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StartProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopAllProcess(t *testing.T) {
	type args struct {
		client *Client
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "stopAllProcess",
			args: args{
				client: newConnect(t, testServer),
				wait:   true,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopAllProcesses(tt.args.client, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopAllProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("StopAllProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartAllProcess(t *testing.T) {
	type args struct {
		client *Client
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "startAllProcess",
			args: args{
				client: newConnect(t, testServer),
				wait:   true,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StartAllProcesses(tt.args.client, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartAllProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("StartAllProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopProcessGroup(t *testing.T) {
	type args struct {
		client *Client
		group  string
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "stopProcessGroup",
			args: args{
				client: newConnect(t, testServer),
				group:  "demo",
				wait:   true,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopProcessGroup(tt.args.client, tt.args.group, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopProcessGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("StopProcessGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartProcessGroup(t *testing.T) {
	type args struct {
		client *Client
		group  string
		wait   bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "startProcessGroup",
			args: args{
				client: newConnect(t, testServer),
				group:  "demo",
				wait:   true,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StartProcessGroup(tt.args.client, tt.args.group, tt.args.wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartProcessGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("StartProcessGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalProcess(t *testing.T) {
	type args struct {
		client *Client
		name   string
		signal string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "signalProcess/signal/1",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				signal: "1",
			},
			want: true,
		},
		{
			name: "signalProcess/signal/error",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				signal: "error",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignalProcess(tt.args.client, tt.args.name, tt.args.signal)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignalProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignalProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalProcessGroup(t *testing.T) {
	type args struct {
		client *Client
		name   string
		signal string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "signalProcessGroup",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo",
				signal: "HUP",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignalProcessGroup(tt.args.client, tt.args.name, tt.args.signal)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignalProcessGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("SignalProcessGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalAllProcesses(t *testing.T) {
	type args struct {
		client *Client
		signal string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "signalAllProcesses",
			args: args{
				client: newConnect(t, testServer),
				signal: "HUP",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignalAllProcesses(tt.args.client, tt.args.signal)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignalAllProcesses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("SignalAllProcesses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReloadConfig(t *testing.T) {
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
			name: "reloadConfig",
			args: args{
				client: newConnect(t, testServer),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReloadConfig(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReloadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("ReloadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveProcessGroup(t *testing.T) {
	type args struct {
		client *Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "removeProcessGroup",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// stop demo group
			StopProcessGroup(tt.args.client, tt.args.name, true)
			got, err := RemoveProcessGroup(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveProcessGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveProcessGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddProcessGroup(t *testing.T) {
	type args struct {
		client *Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "addProcessGroup",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddProcessGroup(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddProcessGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddProcessGroup() = %v, want %v", got, tt.want)
			} else {
				StartProcessGroup(tt.args.client, tt.args.name, true)
			}
		})
	}
}

func TestSendProcessStdin(t *testing.T) {
	type args struct {
		client *Client
		name   string
		chars  string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "sendProcessStdin",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				chars:  "HelloWorld\n",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendProcessStdin(tt.args.client, tt.args.name, tt.args.chars)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendProcessStdin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendProcessStdin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendRemoteCommEvent(t *testing.T) {
	type args struct {
		client *Client
		name   string
		data   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "sendRemoteCommEvent",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo",
				data:   "hi",
			},
			want: true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendRemoteCommEvent(tt.args.client, tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendRemoteCommEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendRemoteCommEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
