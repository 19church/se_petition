package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string

	Students []Student `gorm:"foreignKey:StudentID"`
}

type PetitionType struct {
	gorm.Model
	Name string

	PetitionTypes []PetitionType `gorm:"foreignKey:PetitionTypeID"`
}

type PetitionPeriod struct {
	gorm.Model
	Num uint

	PetitionPeriods []PetitionPeriod `gorm:"foreignKey:PetitionPeriodID"`
}

type Petition struct {
	gorm.Model

	//Student FK
	StudentID *uint
	Student   Student `gorm:"references:id"`

	//PetitionType FK
	PetitionTypeID *uint
	PetitionType   PetitionType `gorm:"references:id"`

	//PetitionPeriod FK
	PetitionPeriodID *uint
	PetitionPeriod   PetitionPeriod `gorm:"references:id"`

	Petition_Reason     string
	Petition_Startdate  time.Time
	Petition_Enddate    time.Time
	Added_Time          time.Time
}