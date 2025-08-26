package api_service

type ServiceConfig struct {
	Name   string            `yaml:"name"`
	Queue  string            `yaml:"queue"`
	Method map[string]string `yaml:"method"`
}
