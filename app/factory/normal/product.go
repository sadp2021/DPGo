package normal



//IRuleConfigParser 配置文件格式化器接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

//JSONRuleConfigParser JSON 格式化
type JSONRuleConfigParser struct {
}

//Parse Parse
func (J JSONRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

//YamlRuleConfigParser yaml 格式化
type YamlRuleConfigParser struct {
}

// Parse Parse
func (Y YamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}
