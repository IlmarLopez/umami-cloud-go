package cfg

// EnvConfig utiliza MixedCaps para funciones y estructuras exportadas
type EnvConfig struct {
	Account      string
	Region       string
	IsProduction bool
}

// GetDevConfig devuelve la configuración para el entorno de desarrollo
func GetDevConfig() EnvConfig {
	return EnvConfig{
		Account:      "123456789012", // Reemplaza con tu ID de AWS si lo tienes
		Region:       "us-east-1",
		IsProduction: false,
	}
}

// GetProdConfig devuelve la configuración para el entorno de producción
func GetProdConfig() EnvConfig {
	return EnvConfig{
		Account:      "123456789012", // Reemplaza con tu ID de AWS si lo tienes
		Region:       "us-east-1",
		IsProduction: true,
	}
}
