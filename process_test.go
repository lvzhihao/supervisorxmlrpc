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
		want    *GetProcessInfoReturn
		wantErr bool
	}{
		{
			name: "getProcessInfo",
			args: args{
				client: newConnect(t, testServer),
				name:   "single",
			},
			want: &GetProcessInfoReturn{
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
