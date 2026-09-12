package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/student-service/internal/service"
)

func NewRouter(
	studentSvc *service.StudentService,
	enrollmentSvc *service.EnrollmentService,
	paymentSvc *service.PaymentService,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	studentHandler := NewStudentHandler(studentSvc)
	enrollmentHandler := NewEnrollmentHandler(enrollmentSvc)
	r.Route("/students", func(r chi.Router) {
		r.Post("/", studentHandler.CreateStudent)
		r.Get("/", studentHandler.ListStudents)
		r.Get("/{id}", studentHandler.GetStudent)
		r.Get("/{id}/enrollments", studentHandler.GetStudentEnrollments)
	})
	r.Get("/schools/{school_id}/students", studentHandler.ListStudentsBySchool)
	r.Route("/enrollments", func(r chi.Router) {
		r.Post("/", enrollmentHandler.CreateEnrollment)
		r.Get("/", enrollmentHandler.ListEnrollmentsByStudent)

	})

	paymentHandler := NewPaymentHandler(paymentSvc)
	r.Route("/payments", func(r chi.Router) {
		r.Post("/", paymentHandler.CreatePayment)
		r.Get("/", paymentHandler.ListPaymentsByStudent)
	})

	return r
}
