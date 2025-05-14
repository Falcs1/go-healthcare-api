package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Appointments []Appointment `gorm:"foreignKey:PatientID"`
}

type Appointment struct {
	gorm.Model
	PatientID   uint      `json:"patient_id" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	Description string    `json:"description"`
	Patient     Patient   `gorm:"foreignKey:PatientID"`
}
