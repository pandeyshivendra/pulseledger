package repositories

import (
	"testing"

	"pulseledger/entities"
	"pulseledger/enums"
	"pulseledger/tests/testdb"

	"github.com/stretchr/testify/assert"
)

func TestSeedOperationTypes_ShouldInsertAllEnumValues(t *testing.T) {
	db := testdb.InitTestDB(t)

	err := SeedOperationTypes(db)
	assert.NoError(t, err)

	for _, enum := range enums.AllOperationTypes() {
		var opType entities.OperationType
		result := db.First(&opType, "id = ?", uint8(enum))
		assert.NoError(t, result.Error)
		assert.Equal(t, enum.Description(), opType.Description)
		assert.False(t, opType.IsDepricated)
	}
}
