package repo

import (
	"fmt"

	"github.com/MXkodo/todo/config"
	"github.com/MXkodo/todo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
