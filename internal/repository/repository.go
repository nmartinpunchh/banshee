package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nmartinpunchh/banshee/configs"
	"github.com/nmartinpunchh/banshee/internal/models"

	// empty import for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// IRepository ..
type IRepository interface {
	GetAll() ([]*models.Workflow, error)
}

// WorkflowRepository ..
type WorkflowRepository struct {
	Env *configs.Env
	db  *gorm.DB
}

// GetAll gets all
func (h *WorkflowRepository) GetAll() ([]*models.Workflow, error) {
	var hellos []*models.Workflow
	if err := h.db.Find(&hellos).Error; err != nil {
		log.Error(err)
	}

}

// Init initializes the db and auto migrates the models
func Init(e *configs.Env) *WorkflowRepository {
	db, err := gorm.Open("mysql", e.DbConnectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Statement{}, &models.Sequence{}, &models.Parallel{}, &models.ActivityInvocation{}, &models.Workflow{})

	hr := &WorkflowRepository{
		Env: e,
		db:  db,
	}

	return hr

}
