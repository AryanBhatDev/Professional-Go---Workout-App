package routes

import (
	"github.com/AryanBhatDev/Professional-Go---Workout-App/internal/app"
	"github.com/go-chi/chi/v5"
)

func HandleRoutes(app *app.Application) *chi.Mux {

	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.CreateWorkout)
	return r
}
