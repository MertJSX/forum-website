package types

type ConfigFile struct {
	Port      int    `yaml:"port"`
	JWTSecret string `yaml:"jwt-secret"`
}
