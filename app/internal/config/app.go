package config

// AppConfiguration stores application configuration
type AppConfiguration struct {
	// FrontendURL stores frontend url
	FrontendURL string `env:"FRONTEND_URL" envDefault:"http://localhost:4200"`
}
