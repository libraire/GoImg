package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

var (
	config *viper.Viper
	logger *log.Logger
	DB     *gorm.DB
)

func Init() {
	// setupLogger()
	// loadConfig()
	// setupDatabase()
}

func loadConfig() {

	config = viper.New()

	logger.Info("Loading config file...")
	// config.SetDefault("sqlite.path", "$HOME/workspace/lensman-server/sqlite/lensman.db")

	config.SetConfigName("config")
	config.SetConfigType("yaml")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "~/.goimg/"
	}
	config.AddConfigPath(configPath)
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	logger.Info("Using config file:", config.ConfigFileUsed())
}

func setupLogger() {

	logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.InfoLevel)
}

func setupDatabase() {

	// pure go version driver instead of CGO"github.com/glebarez/sqlite"
	// path := config.Get("sqlite.path").(string)
	// sqliteDB, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// DB = sqliteDB

}
