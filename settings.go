package boom

//Settings ...
type Settings struct {
	Http *HttpSettings `yaml:"http"`
}

// APISettings ...
type HttpSettings struct {
	Host        string `yaml:"port"`
	MaxBodySize int64  `yaml:"max_body_size"`
}
