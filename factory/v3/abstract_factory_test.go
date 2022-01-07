package v3

import (
	"reflect"
	"testing"
)

type testCase struct {
	name string
	want IRuleConfigParser
}

func TestCreateRuleConfigParser(t *testing.T) {
	cases := []testCase{
		{
			name: "json",
			want: jsonRuleConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleConfigParser(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("CreateRuleConfigParser() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestCreateSystemConfigParser(t *testing.T) {
	cases := []testCase{
		{
			name: "json",
			want: jsonSystemConfigParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateSystemConfigParser(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("CreateSystemConfigParser() = %v, want %v", got, c.want)
			}
		})
	}
}