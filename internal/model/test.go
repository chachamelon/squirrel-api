package model

type TestCase struct {
	Name    string  `yaml:"name"`
	Request Request `yaml:"request"`
	Assert  Assert  `yaml:"assert"`
}

type Request struct {
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    interface{}       `yaml:"body"`
}

type Assert struct {
	Status int               `yaml:"status"`
	Json   map[string]string `yaml:"json"`
}
