package v1

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
	want IRuleConfigParser
}

func TestNewRuleConfigParser(t *testing.T) {
	cases := []testCase{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewRuleConfigParser(c.args.t); !reflect.DeepEqual(got, c.want) {
				t.Errorf("NewRuleConfigParser() = %v, want %v", got, c.want)
			}
		})
	}
}