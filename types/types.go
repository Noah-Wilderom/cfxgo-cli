package types

type WebpackMode string

const (
	WebpackModeProduction  = "production"
	WebpackModeDevelopment = "development"
)

type Config struct {
	Webpack struct {
		Mode     WebpackMode `yaml:"mode"`
		Cache    bool        `yaml:"cache"`
		Filename struct {
			Production  string `yaml:"production"`
			Development string `yaml:"development"`
		} `yaml:"filename"`
	} `yaml:"webpack"`
	Server struct {
		Go struct {
			Enable bool `yaml:"enable"`
		} `yaml:"go"`
		Typescript struct {
			Enable bool   `yaml:"enable"`
			Main   string `yaml:"main,omitempty"`
		} `yaml:"typescript"`
	} `yaml:"server"`
	Client struct {
		Go struct {
			Enable bool   `yaml:"enable"`
			Exec   string `yaml:"exec,omitempty"`
		} `yaml:"go"`
		Typescript struct {
			Enable bool   `yaml:"enable"`
			Main   string `yaml:"main,omitempty"`
		} `yaml:"typescript"`
	} `yaml:"client"`
	Shared struct {
		Typescript struct {
			Enable bool   `yaml:"enable"`
			Main   string `yaml:"main,omitempty"`
		} `yaml:"typescript"`
	} `yaml:"shared"`
}