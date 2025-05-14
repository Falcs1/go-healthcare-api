package main

import (
	"log"

	"healthcare_project/database"
	"healthcare_project/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	router := gin.Default()

	h := &handlers.Handler{DB: database.DB}

	patientRoutes := router.Group("/patients")
	{
		patientRoutes.POST("", h.CreatePatient)
		patientRoutes.GET("", h.GetPatients)
		patientRoutes.GET("/:id", h.GetPatientByID)
		patientRoutes.GET("/:id/appointments", h.GetAppointmentsForPatient)
	}

	appointmentRoutes := router.Group("/appointments")
	{
		appointmentRoutes.POST("", h.CreateAppointment)
		appointmentRoutes.GET("", h.GetAppointments)
	}

	log.Println("Server starting on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
