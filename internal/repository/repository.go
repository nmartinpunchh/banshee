package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nmartinpunchh/testSwagger/configs"
	"github.com/nmartinpunchh/testSwagger/internal/models"

	// empty import for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// IRepository ..
type IRepository interface {
	GetAll([]*models.Hello, error)
}

// HelloRepository ..
type HelloRepository struct {
	Env *configs.Env
	db  *gorm.DB
}

// GetAll gets all
func (h *HelloRepository) GetAll() ([]*models.Hello, error) {
	var hellos []*models.Hello
	if err := h.db.Find(&hellos).Error; err != nil {
		log.Error(err)
	}

}

// Init initializes the db and auto migrates the models
func Init(e *configs.Env) *HelloRepository {
	db, err := gorm.Open("mysql", e.DbConnectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Statement{}, &models.Sequence{}, &models.Parallel{}, &models.ActivityInvocation{}, &models.Workflow{})

	hr := &HelloRepository{
		Env: e,
		db:  db,
	}

	return hr

}
