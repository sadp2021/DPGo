package simple_test

import (
	"testing"

	. "github.com/sadp2021/DPGo/app/factory/simple"
	"github.com/stretchr/testify/assert"
)

func TestNewIRuleConfigParser(t *testing.T) {
	testCase := []struct {
		name string
		args int
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: JSONType,
			want: JSONRuleConfigParser{},
		},
		{
			name: "yaml",
			args: YamlType,
			want: YamlRuleConfigParser{},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewIRuleConfigParser(tt.args)
			assert.IsType(t, got, tt.want)
		})
	}
}
