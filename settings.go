package boom

//Settings ...
type Settings struct {
	Http *HttpSettings `yaml:"http"`
	Grpc *GrpcSettings `yaml:"grpc"`
	Log  *LogSettings  `yaml:"log"`
}

// APISettings ...
type HttpSettings struct {
	Host        string `yaml:"host"`
	Swagger     string `yaml:"swagger"`
	MaxBodySize int64  `yaml:"max_body_size"`
}

type GrpcSettings struct {
	Host string `yaml:"host"`
}

// LogSettings ...
type LogSettings struct {
	Restricteds []string `yaml:"restricteds"`
}
