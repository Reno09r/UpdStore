package service

const(
	ToLogs = "UpdStore.Logs"
	ToErrors = "UpdStore.Errors"
	Info = "info"
	Warning = "warning"
	Error = "error"
	situation = "but the following situation occurred: " 

	UserRegisterFailAttempt = "The user attempted to sign up, " + situation
	UserLoginFailAttempt = "The user attempted to sign in, " + situation
	UserRegisterSuccessful = "Created user with id=%d, fname=%s, lname=%s, username=%s"
	UserLoginSuccessful = "User with username=%s has been logined"


	CatalogCreateFailAttempt = "The admin attempted to create catalog, " + situation
	CatalogCreateSuccessful = "Created catalog with id=%d and name=%s"
	CatalogGetFailAttempt = "The attempted to receive catalog, " + situation
	CatalogGetSuccessful = "The catalog was received"
	CatalogDeleteFailAttempt = "The admin attempted to delete catalog, " + situation
	CatalogDeleteSuccessful = "Deleted catalog with id=%d"
	CatalogUpdateFailAttempt = "The admin attempted to update catalog, " + situation
	CatalogUpdateSuccessful = "Updated catalog with id=%d"

	ManufacturerCreateFailAttempt = "The admin attempted to create manufacturer, " + situation
	ManufacturerCreateSuccessful = "Created manufacturer with id=%d and name=%s"
	ManufacturerGetFailAttempt = "The attempted to receive manufacturer, " + situation
	ManufacturerGetSuccessful = "The manufacturer was received"
	ManufacturerDeleteFailAttempt = "The admin attempted to delete manufacturer, " + situation
	ManufacturerDeleteSuccessful = "Deleted manufacturer with id=%d"
	ManufacturerUpdateFailAttempt = "The admin attempted to update manufacturer, " + situation
	ManufacturerUpdateSuccessful = "Updated manufacturer with id=%d"

	ProductCreateFailAttempt = "The admin attempted to create product, " + situation
	ProductCreateSuccessful = "Created product with id=%d and name=%s"
	ProductGetFailAttempt = "The attempted to receive product, " + situation
	ProductGetSuccessful = "The product was received"
	ProductDeleteFailAttempt = "The admin attempted to delete product, " + situation
	ProductDeleteSuccessful = "Deleted product with id=%d"
	ProductUpdateFailAttempt = "The admin attempted to update product, " + situation
	ProductUpdateSuccessful = "Updated product with id=%d"

	ProductToCartFailAttempt = "The attempted to send product to cart, " + situation
	ProductToCartSuccessful = "User with id=%d sent product to cart"
	GetCartFailAttempt = "The attempted to got products in cart, " + situation
	GetCartSuccessful = "User with id=%d got products in cart"
	DeleteProductInCartFailAttempt = "The attempted to delete product in cart, " + situation
	DeleteProductInCartSuccessful = "User with id=%d deleted product in cart"

	BuyProductsFailAttempt = "The attempted to buy products" + situation
	BuyProductsSuccessful = "User with id=%d bought products"
	GetBoughtProductsFailAttempt = "The attempted to got bought products, " + situation
	GetBoughtProductsSuccessful = "User with id=%d got bought products"

	UserGetProfileFailAttempt = "The user attempted to got data from the profile, " + situation
	UserGetProfileSuccessful = "The user recieved data from the profile"
	UserUpdateProfileFailAttempt = "The user attempted to update data from the profile, " + situation
	UserUpdateProfileSuccessful = "The user updated data from the profile"
	UserDeleteProfileFailAttempt = "The user attempted to delete profile, " + situation
	UserDeleteProfileSuccessful = "The user delete profile"

)