package v2

import (
	"reflect"
	"testing"
)

type args struct {
	t string
}

type testCase struct {
	name string
	args args
	want IRuleConfigParserFactory
}

func TestNewRuleConfigParserFactory(t *testing.T) {
	cases := []testCase{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(c.args.t); !reflect.DeepEqual(got, c.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, c.want)
			}
		})
	}
}