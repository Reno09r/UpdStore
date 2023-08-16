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
	}
	authAdmin := router.Group("/authAdmin")
	{
		authAdmin.POST("/sign-up", h.signUpAdmin)
		authAdmin.POST("/sign-in", h.signInAdmin)
	}

	api := router.Group("/api")
	{
		catalog := api.Group("/catalog")
		{
			catalog.GET("/", h.getCatalogs)
			catalog.GET("/:id", h.getCatalogById)
			products := catalog.Group(":id/products")
    		{
        		products.GET("/", h.getAllProductsByCatalog)
    		}		
		}
		catalogEdit := api.Group("/catalog", h.adminIdentity)
		{
			catalogEdit.POST("/", h.addCatalog)
			catalogEdit.PUT("/:id", h.updateCatalog)
			catalogEdit.DELETE("/:id", h.deleteCatalog)
		}
		manufacturer := api.Group("/manufacturer")
		{
			manufacturer.GET("/", h.getManufacturers)
			manufacturer.GET("/:id", h.getManufacturerById)
			products := manufacturer.Group(":id/products")
			{
				products.GET("/", h.getAllProductsByManufacturer)
			}
		}
		manufacturerEdit := api.Group("/manufacturer", h.adminIdentity)
		{
			manufacturerEdit.POST("/", h.addManufacturer)
			manufacturerEdit.PUT("/:id", h.updateManufacturer)
			manufacturerEdit.DELETE("/:id", h.deleteManufacturer)
		}
		products := api.Group("/products")
		{
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductById)
		}
		productsEdit := api.Group("/products", h.adminIdentity)
		{
			productsEdit.POST("/", h.аddProduct)
			productsEdit.PUT("/:id", h.updateProduct)
			productsEdit.DELETE("/:id", h.deleteProduct)
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