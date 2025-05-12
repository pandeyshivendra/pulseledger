package db

import (
	"pulseledger/config"
	"pulseledger/entities"
	"pulseledger/repositories"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	once     sync.Once
)

func Init() *gorm.DB {
	once.Do(func() {
		database = InitWithConfig(config.GetConfig())
	})
	return GetDatabase()
}

func InitWithConfig(cfg config.Config) *gorm.DB {
	dbPath := cfg.GetDBPath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database %s: %v", dbPath, err)
	}

	log.Infof("Connected to SQLite database at %s", dbPath)

	if cfg.IsDbResetEnabled() {
		resetAndMigrateDB(db)
	}

	return db
}

func GetDatabase() *gorm.DB {
	return database
}

func resetAndMigrateDB(db *gorm.DB) {
	log.Warn("Starting database reset and migration (dev only)")

	log.Info("Dropping existing tables...")
	err := db.Migrator().DropTable(
		&entities.Transaction{},
		&entities.Account{},
		&entities.OperationType{},
	)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Info("Tables dropped successfully")

	log.Info("Running auto migration...")
	err = db.AutoMigrate(
		&entities.Account{},
		&entities.Transaction{},
		&entities.OperationType{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}
	log.Info("Auto migration completed")

	log.Info("Seeding operation types...")
	if err = repositories.SeedOperationTypes(db); err != nil {
		log.Fatalf("Failed to seed operation types: %v", err)
	}
	log.Info("Successfully seeded operation types")
}
