package supervisorxmlrpc

import (
	"testing"
)

func TestReadProcessStdoutLog(t *testing.T) {
	type args struct {
		client *Client
		name   string
		offset int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "readProcessStdoutLog",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				offset: 0,
				length: 1024,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadProcessStdoutLog(tt.args.client, tt.args.name, tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadProcessStdoutLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("ReadProcessStdoutLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadProcessStderrLog(t *testing.T) {
	type args struct {
		client *Client
		name   string
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
			name: "readProcessStderrLog",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				offset: 0,
				length: 1024,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadProcessStderrLog(tt.args.client, tt.args.name, tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadProcessStderrLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadProcessStderrLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTailProcessStdoutLog(t *testing.T) {
	type args struct {
		client *Client
		name   string
		offset int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "tailProcessStdoutLog",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				offset: 0,
				length: 1024,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TailProcessStdoutLog(tt.args.client, tt.args.name, tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("TailProcessStdoutLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("TailProcessStdoutLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTailProcessStderrLog(t *testing.T) {
	type args struct {
		client *Client
		name   string
		offset int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "tailProcessStderrLog",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
				offset: 0,
				length: 1024,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TailProcessStderrLog(tt.args.client, tt.args.name, tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("TailProcessStderrLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("TailProcessStderrLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearProcessLogs(t *testing.T) {
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
			name: "clearProcessLogs",
			args: args{
				client: newConnect(t, testServer),
				name:   "demo:single",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClearProcessLogs(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClearProcessLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClearProcessLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearAllProcessLogs(t *testing.T) {
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
			name: "clearAllProcessLogs",
			args: args{
				client: newConnect(t, testServer),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClearAllProcessLogs(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClearAllProcessLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("ClearAllProcessLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
