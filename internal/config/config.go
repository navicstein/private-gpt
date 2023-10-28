package config

import (
	"fmt"
	env "github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

type EnvironmentConfig struct {
	Name    string `env:"ENVIRONMENT" envDefault:"development"`
	BaseURL string `env:"BASE_URL" envDefault:"https://localhost:1337"`
	Port    string `env:"PORT" envDefault:"1337"`
}

type CorsConfig struct {
	Origins     string `env:"CORS_ORIGINS" envDefault:"*"`
	Methods     string `env:"CORS_METHODS" envDefault:"GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS"`
	Headers     string `env:"CORS_HEADERS" envDefault:"Access-Control-Allow-Origin,Authorization,Origin,Content-Type,Accept,Accept-Language,Content-Length,Cache-Control,X-Requested-With,V-Api-Key"`
	Credentials bool   `env:"CORS_CREDENTIALS" envDefault:"true"`
}

type OpenAIConfig struct {
	APIKey  string `env:"OPENAI_API_KEY"`
	BaseURL string `env:"OPENAI_BASE_URL"`
}

type DeepgramConfig struct {
	APIKey string `env:"DEEPGRAM_API_KEY"`
}

type WeaviateConfig struct {
	Local  bool   `env:"WEAVIATE_LOCAL"`
	Scheme string `env:"WEAVIATE_SCHEME"`
	APIKey string `env:"WEAVIATE_API_KEY"`
	Host   string `env:"WEAVIATE_HOST"`
}

type SwitchersConfig struct {
	DBDriver string `env:"DB_DRIVER"`
	Embedder string `env:"EMBEDDER"`
}

type HuggingFaceConfig struct {
	APIToken string `env:"HUGGINGFACEHUB_API_TOKEN"`
}

type AppConfig struct {
	OpenAI      OpenAIConfig
	Deepgram    DeepgramConfig
	Weaviate    WeaviateConfig
	Switchers   SwitchersConfig
	HuggingFace HuggingFaceConfig
	Cors        CorsConfig
	Environment EnvironmentConfig
	VectorStore string `env:"VECTOR_STORE"`
}

var cfg AppConfig

func GetConfig() AppConfig {
	return cfg
}

func Setup() error {
	mode := os.Getenv("ENV")

	// production must explicitly send ENV in docker
	if mode == "" {
		mode = "development"
	}

	cwd, _ := os.Getwd()
	envName := fmt.Sprintf(".env.%s", mode)
	envPath := filepath.Join(cwd, envName)

	if err := godotenv.Load(envPath); err != nil {
		log.Warn().Str("path", envPath).Msg("No .env.development file found in filesystem, skipping..")
	}

	if err := env.Parse(&cfg); err != nil {
		return err
	}
	return nil
}
