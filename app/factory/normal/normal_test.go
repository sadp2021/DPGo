package normal_test

import (
	"testing"

	. "github.com/sadp2021/DPGo/app/factory/normal"
	"github.com/stretchr/testify/assert"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	testCase := []struct {
		name string
		args int
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: JSONType,
			want: JSONRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: YamlType,
			want: YamlRuleConfigParserFactory{},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewIRuleConfigParserFactory(tt.args)

			assert.IsType(t, got, tt.want)
		})
	}
}
