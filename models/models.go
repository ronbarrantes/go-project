package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Store     []Store `gorm:"many2many:visits;"`
}

type Points struct {
	gorm.Model
	ID          string `gorm:"primaryKey;not null"`
	CustomerID  string
	StoreID     string
	NumberOfPts int
	UpdatedAt   time.Time
}

type RewardItem struct {
	gorm.Model
	ID         string `gorm:"primaryKey;not null"`
	CustomerID string
	StoreID    string
	RewardName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Rewards struct {
	gorm.Model
	ID           string `gorm:"primaryKey;not null"`
	CustomerID   string
	StoreID      string
	PointsCashed int
	RedeemedOn   time.Time
	RewardItemID string
	RewardItem   *RewardItem `gorm:"foreignKey:RewardItemID"`
	// RedeemedBy   string
}

type Store struct {
	gorm.Model
	ID          string `gorm:"primaryKey;not null"`
	Name        string
	Description string
	Phone       string
	Address     string
	Owner       User `gorm:"foreignKey:OwnerID"`
	OwnerID     string
	Customers   []Customer `gorm:"many2many:visits;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	gorm.Model
	ID        string `gorm:"index;unique;not null"`
	FistName  string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Customer  *Customer `gorm:"foreignKey:UserID"`
}

type Visits struct {
	gorm.Model
	ID         string `gorm:"primaryKey;not null"`
	CustomerID string
	StoreID    string
	CreatedAt  time.Time
}
