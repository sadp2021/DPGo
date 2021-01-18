package normal

import "errors"

//对外暴露的类型枚举
const (
	JSONType = iota
	YamlType
)

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// YamlRuleConfigParserFactory YamlRuleConfigParser 的工厂类
type YamlRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (y YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return YamlRuleConfigParser{}
}

// JSONRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type JSONRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (j JSONRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return JSONRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t int) (IRuleConfigParserFactory, error) {
	switch t {
	case JSONType:
		return JSONRuleConfigParserFactory{}, nil
	case YamlType:
		return YamlRuleConfigParserFactory{}, nil
	default:
	}
	return nil, errors.New("wrong type")
}
