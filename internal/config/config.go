package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port         string
	DatabaseURL  string
	MQTTBroker   string
	MQTTPort     int
	MQTTUser     string
	MQTTPassword string
	MQTTClientID string
}

func LoadConfig() *Config {
	port := getEnv("PORT", "8080")
	dbURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/smart_irrigation?sslmode=disable")
	mqttBroker := getEnv("MQTT_BROKER", "localhost")
	mqttPortStr := getEnv("MQTT_PORT", "1883")
	mqttPort, err := strconv.Atoi(mqttPortStr)
	if err != nil {
		mqttPort = 1883
	}
	mqttUser := getEnv("MQTT_USER", "")
	mqttPass := getEnv("MQTT_PASSWORD", "")
	mqttClientID := getEnv("MQTT_CLIENT_ID", "smart-irrigation-web-backend")

	return &Config{
		Port:         port,
		DatabaseURL:  dbURL,
		MQTTBroker:   mqttBroker,
		MQTTPort:     mqttPort,
		MQTTUser:     mqttUser,
		MQTTPassword: mqttPass,
		MQTTClientID: mqttClientID,
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
