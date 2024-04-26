package tdclient

type config struct {
	Env     string
	Version string
	Host    string
}

func defaultCfg() *config {
	return &config{
		Env:     "api",
		Version: "v4",
		Host:    ".tibiadata.com",
	}
}

func (c *config) url() string {
	return "https://" + c.Env + c.Host + "/" + c.Version + "/"
}

func WithEnv(env string) {
	c.Env = env
}
