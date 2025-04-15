package models

import "gorm.io/gorm"

type FeatureFlag struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	Description  string
	DefaultValue bool         `gorm:"default:false"`
	Dependencies []Dependency `gorm:"foreignKey:ParentID"`
}

type Dependency struct {
	gorm.Model
	ParentID uint        `gorm:"not null"`
	Parent   FeatureFlag `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChildID  uint        `gorm:"not null"`
	Child    FeatureFlag `gorm:"foreignKey:ChildID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Client struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type ClientFeatureFlag struct {
	gorm.Model
	ClientID      uint        `gorm:"not null"`
	Client        Client      `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeatureFlagID uint        `gorm:"not null"`
	FeatureFlag   FeatureFlag `gorm:"foreignKey:FeatureFlagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsEnabled     bool        `gorm:"default:false"`
}
