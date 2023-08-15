package handler

import (
	"github.com/Reno09r/Store/server/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.PUT("/update", h.update)
	}

	api := router.Group("/api")
	{
		catalog := api.Group("/catalog")
		{
			catalog.POST("/", h.addCatalog)
			catalog.GET("/", h.getCatalogs)
			catalog.GET("/:id", h.getCatalogById)
			catalog.PUT("/:id", h.updateCatalog)
			catalog.DELETE("/:id", h.deleteCatalog)
			products := catalog.Group(":id/products")
    		{
        		products.GET("/", h.getAllProductsByCatalog)
    		}		
		}
		manufacturer := api.Group("/manufacturer")
		{
			manufacturer.POST("/", h.addManufacturer)
			manufacturer.GET("/", h.getManufacturers)
			manufacturer.GET("/:id", h.getManufacturerById)
			manufacturer.PUT("/:id", h.updateManufacturer)
			manufacturer.DELETE("/:id", h.deleteManufacturer)

			products := manufacturer.Group(":id/products")
			{
				products.GET("/", h.getAllProductsByManufacturer)
			}
		}
		products := api.Group("/products")
		{
			products.POST("/", h.аddProduct)
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductById)
			products.PUT("/:id", h.updateProduct)
			products.DELETE("/:id", h.deleteProduct)
		}
		cart := api.Group("/cart", h.userIdentity)
		{
			cart.POST("/", h.insertProductInCart)
			cart.GET("/", h.getProductsFromCart)
			cart.DELETE("/:id", h.deleteProductFromCart)
		}
	}
	return router
}