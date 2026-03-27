package api

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/DigitLock/invoice-generator/backend/internal/api/handlers"
	"github.com/DigitLock/invoice-generator/backend/internal/api/middleware"
	"github.com/DigitLock/invoice-generator/backend/internal/auth"
	"github.com/DigitLock/invoice-generator/backend/internal/config"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(cfg *config.Config, pool *pgxpool.Pool, repos *repository.Repositories, jwtService *auth.JWTService) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Server.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	healthHandler := handlers.NewHealthHandler(pool)
	companyHandler := handlers.NewCompanyHandler(repos.Companies)
	clientHandler := handlers.NewClientHandler(repos.Clients)
	bankAccountHandler := handlers.NewBankAccountHandler(repos.BankAccounts, repos.Companies)
	invoiceHandler := handlers.NewInvoiceHandler(repos.Invoices, repos.Companies, repos.Clients)

	r.Get("/health", healthHandler.Health)
	r.Get("/ready", healthHandler.Ready)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(jwtService))

		r.Route("/companies", func(r chi.Router) {
			r.Get("/", companyHandler.List)
			r.Post("/", companyHandler.Create)
			r.Get("/{id}", companyHandler.Get)
			r.Put("/{id}", companyHandler.Update)
			r.Delete("/{id}", companyHandler.Delete)

			r.Route("/{company_id}/bank-accounts", func(r chi.Router) {
				r.Get("/", bankAccountHandler.List)
				r.Post("/", bankAccountHandler.Create)
			})
		})

		r.Route("/bank-accounts", func(r chi.Router) {
			r.Put("/{id}", bankAccountHandler.Update)
			r.Delete("/{id}", bankAccountHandler.Delete)
		})

		r.Route("/clients", func(r chi.Router) {
			r.Get("/", clientHandler.List)
			r.Post("/", clientHandler.Create)
			r.Get("/{id}", clientHandler.Get)
			r.Put("/{id}", clientHandler.Update)
			r.Delete("/{id}", clientHandler.Delete)
		})

		r.Route("/invoices", func(r chi.Router) {
			r.Get("/", invoiceHandler.List)
			r.Post("/", invoiceHandler.Create)
			r.Get("/{id}", invoiceHandler.Get)
			r.Put("/{id}", invoiceHandler.Update)
			r.Delete("/{id}", invoiceHandler.Delete)
			r.Patch("/{id}/status", invoiceHandler.UpdateStatus)
			r.Patch("/{id}/overdue", invoiceHandler.UpdateOverdue)
			r.Get("/{id}/pdf", invoiceHandler.GeneratePDF)
		})
	})

	return r
}
