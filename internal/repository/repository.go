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
	Create(model *models.Workflow) (*models.Workflow, error)
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

// Create creates a workflow
func (h *WorkflowRepository) Create(model *models.Workflow) (*models.Workflow, error) {
	if err := h.Db.Create(&model.Root.ActivityInvocation).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return model, nil

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

	// defer func() {
	// 	log.Println("closing")
	// 	err := db.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// Migrate the schema
	db.AutoMigrate(&models.Statement{}, &models.Sequence{}, &models.Parallel{}, &models.ActivityInvocation{}, &models.Workflow{}, &models.Argument{})

	hr := &WorkflowRepository{
		Env: e,
		Db:  db,
	}

	return hr

}
