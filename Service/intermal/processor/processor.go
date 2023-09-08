package processor

import (
	"module/Service/intermal/db"
	"module/Service/intermal/models"
)

type ModelsProcessor struct {
	storage *db.ModuleStorage
}

func NewModelProcessor(storage.db.ModuleStorage) *ModuleProcessor {
	processor := new(ModelsProcessor)
	processor.storage = storage
	return processor
}

func (processor *ModuleProcessor) ListModels(nameFilter string) ([]models.Models) {
	return processor.storage.GetModelsList(nameFilter)
}
