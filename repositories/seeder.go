package repositories

import (
	"pulseledger/entities"
	"pulseledger/enums"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedOperationTypes(db *gorm.DB) error {
	for _, enum := range enums.AllOperationTypes() {
		opType := entities.OperationType{
			ID:           uint8(enum),
			Description:  enum.Description(),
			IsDepricated: false,
		}

		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&opType).Error; err != nil {
			return err
		}
	}
	return nil
}
