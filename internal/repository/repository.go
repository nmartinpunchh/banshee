package repository

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/nmartinpunchh/banshee/configs"
	"github.com/nmartinpunchh/banshee/internal/models"

	// empty import for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// IRepository ..
type IRepository interface {
	GetAll() ([]*models.Workflow, error)
	Create(model *models.Workflow) (*models.Workflow, error)
	GetByID(id int) (*models.Workflow, error)
	Delete(id int) (int, error)
}

// WorkflowRepository ..
type WorkflowRepository struct {
	Env *configs.Env
	Db  *gorm.DB
}

// GetAll gets all
func (h *WorkflowRepository) GetAll() ([]*models.Workflow, error) {
	var workflows []*models.Workflow
	if err := h.Db.Find(&workflows).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return workflows, nil

}

// GetByID gets a workflow by ID
func (h *WorkflowRepository) GetByID(id int) (*models.Workflow, error) {
	workflow := &models.Workflow{}
	log.Println(id)
	if err := h.Db.Set("gorm:auto_preload", true).First(&workflow, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return workflow, nil

}

// Create creates a workflow
func (h *WorkflowRepository) Create(model *models.Workflow) (*models.Workflow, error) {
	if err := h.Db.Create(&model).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return model, nil

}

// Delete deletes a workflow
func (h *WorkflowRepository) Delete(id int) (int, error) {
	model := models.Workflow{}
	if err := h.Db.Delete(&model, id).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil

}

// Init initializes the db and auto migrates the models
func Init(e *configs.Env) *WorkflowRepository {
	// log.Println(e.DbConnectionString)
	db, err := gorm.Open("mysql", e.DbConnectionString)
	db.LogMode(true)

	if err != nil {
		log.Panicf("failed to connect database %s", err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.Statement{}, &models.Sequence{}, &models.Parallel{}, &models.ActivityInvocation{}, &models.Workflow{}, &models.Argument{})

	hr := &WorkflowRepository{
		Env: e,
		Db:  db,
	}

	return hr

}
