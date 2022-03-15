package example

import "testing"

var (
	namespace       string = "test"
	emqxApiUsername string = "admin"
	emqxApiPassword string = "public"
)

func TestGenerateTelegrafConf(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "telegraf-test",
			args: args{
				t: Task{
					Namespace:       namespace,
					EmqxApiUsername: emqxApiUsername,
					EmqxApiPassword: emqxApiPassword,
					TelegrafConf: TelegrafConf{
						MetricConfig: MetricConfig{
							BCMetricGateway: "http://gateway",
						},
						LogType: "external",
						LogConfig: LogConfig{
							EsUrl:       "http://url",
							EsUsername:  "esUsername",
							EsPassword:  "esPassword",
							EsIndexName: "esIndexName",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateTelegrafConf(tt.args.t)
		})
	}
	// task := Task{
	// 	Namespace:       namespace,
	// 	EmqxApiUsername: emqxApiUsername,
	// 	EmqxApiPassword: emqxApiPassword,
	// 	TelegrafConf: TelegrafConf{
	// 		MetricConfig: MetricConfig{
	// 			BCMetricGateway: "http://gateway",
	// 		},
	// 		LogType: "external",
	// 		LogConfig: LogConfig{
	// 			EsUrl:       "http://url",
	// 			EsUsername:  "esUsername",
	// 			EsPassword:  "esPassword",
	// 			EsIndexName: "esIndexName",
	// 		},
	// 	},
	// }
	// GenerateTelegrafConf(task)
}
