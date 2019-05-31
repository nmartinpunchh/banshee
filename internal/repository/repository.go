package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/nmartinpunchh/banshee/configs"
	"github.com/nmartinpunchh/banshee/internal/models"

	// empty import for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// IRepository ..
type IRepository interface {
	GetAll() ([]*models.Workflow, error)
	Create(*models.Workflow) (*models.Workflow, error)
}

// WorkflowRepository ..
type WorkflowRepository struct {
	Env *configs.Env
	db  *gorm.DB
}

// GetAll gets all
func (h *WorkflowRepository) GetAll() ([]*models.Workflow, error) {
	var workflows []*models.Workflow
	if err := h.db.Find(&workflows).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return workflows, nil

}

// Create creates a workflow
func (h *WorkflowRepository) Create(*models.Workflow) (*models.Workflow, error) {
	var workflow *models.Workflow
	if err := h.db.Create(&workflow).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return workflow, nil

}

// Init initializes the db and auto migrates the models
func Init(e *configs.Env) *WorkflowRepository {
	connStr := "root:password@tcp(localhost)/punchh"
	// log.Println(e.DbConnectionString)
	db, err := gorm.Open("mysql", connStr)
	db.LogMode(true)
	if err != nil {
		log.Panicf("failed to connect database %s", err.Error())
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
