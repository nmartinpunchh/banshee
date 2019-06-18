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
	GetAll() ([]*models.Journey, error)
	Create(model *models.Journey) (*models.Journey, error)
	GetByID(id int) (*models.Journey, error)
	Delete(id int) (int, error)
}

// JourneyRepository ..
type JourneyRepository struct {
	Env *configs.Env
	Db  *gorm.DB
}

// GetAll gets all
func (h *JourneyRepository) GetAll() ([]*models.Journey, error) {
	var journeys []*models.Journey
	if err := h.Db.Find(&journeys).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return journeys, nil

}

// GetByID gets a journey by ID
func (h *JourneyRepository) GetByID(id int) (*models.Journey, error) {
	journey := &models.Journey{}
	log.Println(id)
	if err := h.Db.Set("gorm:auto_preload", true).First(&journey, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return journey, nil

}

// Create creates a journey
func (h *JourneyRepository) Create(model *models.Journey) (*models.Journey, error) {
	if err := h.Db.Create(&model).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return model, nil

}

// Delete deletes a journey
func (h *JourneyRepository) Delete(id int) (int, error) {
	model := models.Journey{}
	if err := h.Db.Delete(&model, id).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil

}

// Init initializes the db and auto migrates the models
func Init(e *configs.Env) *JourneyRepository {
	// log.Println(e.DbConnectionString)
	db, err := gorm.Open("mysql", e.DbConnectionString)
	db.LogMode(true)

	if err != nil {
		log.Panicf("failed to connect database %s", err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.Statement{}, &models.Sequence{}, &models.Parallel{}, &models.ActivityInvocation{}, &models.Journey{}, &models.Argument{}, &models.Journey{})

	hr := &JourneyRepository{
		Env: e,
		Db:  db,
	}

	return hr

}
