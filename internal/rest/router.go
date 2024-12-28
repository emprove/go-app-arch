package rest

import (
	"net/http"

	"go-app-arch/internal/config"
	"go-app-arch/internal/database"
	"go-app-arch/internal/database/postgres"
	"go-app-arch/internal/mapper"
	"go-app-arch/internal/middleware"
	"go-app-arch/internal/service"
	"go-app-arch/internal/usecase"

	"github.com/go-pkgz/routegroup"
)

func NewRouter(cfg *config.Cfg, ds *config.DynamicState, db database.DB) http.Handler {
	productMapper := mapper.NewProductMapper(cfg)
	fileMapper := mapper.NewFileMapper(cfg)

	productRepo := postgres.NewProductRepository(db, productMapper, fileMapper)
	userRepo := postgres.NewUserRepository(db)
	settingsRepo := postgres.NewSettingsRepository(db)

	productSrv := service.NewProductService(ds, productRepo)
	userSrv := service.NewUserService(userRepo)

	infoUsecase := usecase.NewInfo(cfg, ds, settingsRepo)

	rr := routegroup.New(http.NewServeMux())
	limiter := middleware.NewIPRateLimiter(1, 25)
	rr.Use(middleware.PanicRecovery())
	rr.Use(middleware.RateLimiter(limiter))
	rr.Use(middleware.Cors(cfg.GetAllowedOrigins()))
	rr.Use(middleware.Locale(cfg.AvailableLocalesIso(), ds))

	rr.HandleFunc("GET /", NotFound)

	productHandlerReg(rr, productSrv)
	infoHandlerReg(rr, infoUsecase)

	// adm stack
	rrAdm := rr.Mount("/adm")
	rrAdm.Use(middleware.Authenticate(userSrv))

	productHandlerAdmReg(rrAdm, productSrv)

	return rr
}

func productHandlerAdmReg(rr *routegroup.Bundle, sp *service.ProductService) {
	handler := NewProductHandlerAdm(sp)

	rr.HandleFunc("GET /products", handler.FindList)
	rr.HandleFunc("GET /products/{id}", handler.FindOne)
}

func productHandlerReg(rr *routegroup.Bundle, sp *service.ProductService) {
	handler := NewProductHandler(sp)

	rr.HandleFunc("GET /products", handler.FindList)
	rr.HandleFunc("GET /products/one", handler.FindOne)
}

func infoHandlerReg(rr *routegroup.Bundle, u *usecase.Info) {
	handler := NewInfoHandler(u)

	rr.HandleFunc("GET /info/locales", handler.GetLocales)
	rr.HandleFunc("GET /info/config", handler.GetConfig)
}
