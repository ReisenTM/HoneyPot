package config

type Config struct {
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

// Logger 日志输出格式设置
type Logger struct {
	Format  string `yaml:"format"`
	Level   string `yaml:"level"`
	AppName string `yaml:"appName"`
}

type System struct {
	GrpcManagerAddr string `yaml:"grpcManagerAddr"`
	Network         string `yaml:"network"`
	Uid             string `yaml:"uid"`
}
