package entities

type Account struct {
	BaseEntity
	DocumentNumber uint64        `gorm:"not null"`
	Transactions   []Transaction `gorm:"forignKey:AccountID"`
}
