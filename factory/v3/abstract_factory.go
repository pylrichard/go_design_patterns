package v3

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type ISystemConfigParser interface {
	Parse(data []byte)
}

type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) Parse(data []byte) {
	panic("implement me")
}

type IConfigParserFactory interface {
	CreateRuleConfigParser() IRuleConfigParser
	CreateSystemConfigParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleConfigParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemConfigParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}