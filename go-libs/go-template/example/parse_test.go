package example

import "testing"

func TestGenerateTelegrafConf(t *testing.T) {
	type args struct {
		t *TelegrafConfStruct
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "telegraf-test",
			args: args{
				t: &TelegrafConfStruct{
					InstanceID: "test",
					Username:   "test",
					Password:   "test",
					Url:        "test",
					Http:       "http",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateTelegrafConf(tt.args.t)
		})
	}
}
