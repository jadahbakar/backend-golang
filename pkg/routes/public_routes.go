package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jadahbakar/backend-golang/internal/me"
	"github.com/jadahbakar/backend-golang/pkg/config"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, conf config.Config, pgxPool *pgxpool.Pool) {
	// Create routes group.
	router := a.Group(fmt.Sprintf("%s%s", conf.ApiURLGroup, conf.ApiURLVersion))
	me.AddRoutes(router)

	//--- Master
	// master.AddRoutes(router)
	// masterRepo := shop_infra_product.NewMemoryRepository()

	// master_interfaces_http.NewMasterHandler(router)
	// masterRepo := master_repository.NewMasterRepository(pgxPool)
	// master_interfaces_http.AddRoutes(router, entities.AssesmentEnvironmentEntity)

	//-----------

}
