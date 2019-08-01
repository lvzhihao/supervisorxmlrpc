package supervisorxmlrpc

import (
	"reflect"
	"testing"
)

func TestListMethods(t *testing.T) {
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
			name: "listMethods",
			args: args{
				client: newConnect(t, testServer),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListMethods(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMethods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("ListMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethodHelp(t *testing.T) {
	type args struct {
		client *Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "methodHelp",
			args: args{
				client: newConnect(t, testServer),
				name:   "supervisor.getProcessInfo",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodHelp(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodHelp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("MethodHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethodSignature(t *testing.T) {
	type args struct {
		client *Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "methodSignature",
			args: args{
				client: newConnect(t, testServer),
				name:   "supervisor.getProcessInfo",
			},
			want: []string{"struct", "string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodSignature(tt.args.client, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MethodSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulticall(t *testing.T) {
	type args struct {
		client *Client
		calls  []MulticallParams
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
		wantErr bool
	}{
		{
			name: "multicall",
			args: args{
				client: newConnect(t, testServer),
				calls: []MulticallParams{
					{
						MethodName: "supervisor.getAPIVersion",
						Params:     nil,
					},
					{
						MethodName: "system.methodHelp",
						Params: []interface{}{
							"supervisor.getAPIVersion",
						},
					},
					{
						MethodName: "supervisor.getState",
						Params:     nil,
					},
				},
			},
			wantRet: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := Multicall(tt.args.client, tt.args.calls)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multicall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotRet) < tt.wantRet {
				t.Errorf("Multicall() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
