package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	domain "mcd/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	MCDUsecase domain.MCDUsecase
}

func NewMCDHandler(e *echo.Echo, useCase domain.MCDUsecase) {
	handler := &delivery{MCDUsecase: useCase}

	// User routes
	e.POST("/v1/create/user", handler.createUser)
	e.POST("/v1/user/login", handler.login)

	// e.POST("/v1/delete/user", handler.deleteUser)
	// e.POST("/v1/update/user", handler.updateUser)
	e.GET("/v1/user/:userID", handler.getUserById)

	// Product routes
	e.POST("/v1/create/product", handler.createProduct)
	// e.POST("/v1/delete/product", handler.deleteProduct)
	// e.POST("/v1/update/product", handler.updateProduct)
	e.GET("/v1/product/:productID", handler.getProductById)
	e.GET("/v1/hotel/:hotelID/products", handler.getProductsByHotel)

	// Hotel routes
	e.POST("/v1/create/hotel", handler.createHotel)
	// e.POST("/v1/delete/hotel", handler.deleteHotel)
	// e.POST("/v1/update/hotel", handler.updateHotel)
	e.GET("/v1/hotel", handler.getHotels)

	// User Cart routes
	e.POST("/v1/add/user/cart", handler.addProductToCart)
	e.POST("/v1/delete/user/cart", handler.deleteProductFromCart)
	e.POST("/v1/update/user/cart", handler.updateQuantityInCart)
	e.GET("/v1/user/cart", handler.getUserCart)

	// Order routes
	e.POST("/v1/hotel/:hotelID/create/order", handler.CreateOrder)
	// e.POST("/v1/hotel/:hotelID/update/order", handler.updateOrder)
	// e.GET("/v1/order/:orderID/completed", handler)
	e.GET("/v1/user/:userID/orders", handler.getUserOrders)

	// Payment routes
	//e.GET("/v1/zeel/qrcode", handler.createUser) // Placeholder for now
	//e.POST("/v1/upload/user/qrcode", handler.createUser)

	// Health check route
	e.GET("/", handler.healthCheck)
}

func RoleCheckMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if claims["role"] != role {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Insufficient permissions",
				})
			}
			return next(c)
		}
	}
}

// Health check handler
func (delivery *delivery) healthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, "server is up and running")
}

// User-related handlers
func (delivery *delivery) createUser(context echo.Context) error {
	var user domain.User
	err := json.NewDecoder(context.Request().Body).Decode(&user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = delivery.MCDUsecase.CreateUser(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, "User created successfully")
}

func (delivery *delivery) login(context echo.Context) error {
	var userData domain.UserLogin
	err := json.NewDecoder(context.Request().Body).Decode(&userData)
	if err != nil {
		return err
	}
	res, err := delivery.MCDUsecase.UserLogin(userData)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, res)
}

// func (delivery *delivery) deleteUser(context echo.Context) error {
// 	userID := context.QueryParam("userID")
// 	if userID == "" {
// 		return context.JSON(http.StatusBadRequest, "userID is required")
// 	}

// 	err := delivery.MCDUsecase.DeleteUser(userID)
// 	if err != nil {
// 		return context.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return context.JSON(http.StatusOK, "User deleted successfully")
// }

// func (delivery *delivery) updateUser(context echo.Context) error {
// 	var user domain.User
// 	err := json.NewDecoder(context.Request().Body).Decode(&user)
// 	if err != nil {
// 		return context.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	err = delivery.MCDUsecase.UpdateUser(user)
// 	if err != nil {
// 		return context.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return context.JSON(http.StatusOK, "User updated successfully")
// }

func (delivery *delivery) getUserById(context echo.Context) error {
	userID := context.Param("userID")
	if userID == "" {
		return context.JSON(http.StatusBadRequest, "userID is required")
	}

	user, err := delivery.MCDUsecase.GetUserById(userID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, user)
}

// Product-related handlers
func (delivery *delivery) createProduct(context echo.Context) error {
	var product domain.Product
	err := json.NewDecoder(context.Request().Body).Decode(&product)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = delivery.MCDUsecase.CreateProduct(product)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, "Product created successfully")
}

// func (delivery *delivery) deleteProduct(context echo.Context) error {
// 	productID := context.QueryParam("productID")
// 	if productID == "" {
// 		return context.JSON(http.StatusBadRequest, "productID is required")
// 	}

// 	err := delivery.MCDUsecase.DeleteProduct(productID)
// 	if err != nil {
// 		return context.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return context.JSON(http.StatusOK, "Product deleted successfully")
// }

// func (delivery *delivery) updateProduct(context echo.Context) error {
// 	var product domain.Product
// 	err := json.NewDecoder(context.Request().Body).Decode(&product)
// 	if err != nil {
// 		return context.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	err = delivery.MCDUsecase.UpdateProduct(product)
// 	if err != nil {
// 		return context.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return context.JSON(http.StatusOK, "Product updated successfully")
// }

func (delivery *delivery) getProductById(context echo.Context) error {
	productID := context.Param("productID")
	if productID == "" {
		return context.JSON(http.StatusBadRequest, "productID is required")
	}

	product, err := delivery.MCDUsecase.GetProductById(productID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, product)
}

func (delivery *delivery) getProductsByHotel(context echo.Context) error {
	hotelID := context.Param("hotelID")
	if hotelID == "" {
		return context.JSON(http.StatusBadRequest, "hotelID is required")
	}

	products, err := delivery.MCDUsecase.GetProductsByHotel(hotelID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, products)
}

// Hotel-related handlers
func (delivery *delivery) createHotel(context echo.Context) error {
	var hotel domain.Hotel
	err := json.NewDecoder(context.Request().Body).Decode(&hotel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = delivery.MCDUsecase.CreateHotel(hotel)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, "Hotel created successfully")
}

func (delivery *delivery) deleteHotel(context echo.Context) error {
	hotelID := context.QueryParam("hotelID")
	if hotelID == "" {
		return context.JSON(http.StatusBadRequest, "hotelID is required")
	}

	err := delivery.MCDUsecase.DeleteHotel(hotelID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "Hotel deleted successfully")
}

func (delivery *delivery) updateHotel(context echo.Context) error {
	var hotel domain.Hotel
	err := json.NewDecoder(context.Request().Body).Decode(&hotel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = delivery.MCDUsecase.UpdateHotel(hotel)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "Hotel updated successfully")
}

func (delivery *delivery) getHotels(context echo.Context) error {
	hotels, err := delivery.MCDUsecase.GetHotels()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, hotels)
}

// User cart-related handlers
func (delivery *delivery) addProductToCart(context echo.Context) error {
	var cartProduct domain.CartProducts
	err := json.NewDecoder(context.Request().Body).Decode(&cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err = delivery.MCDUsecase.AddProductToCart(cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, "product is added to cart")
}

func (delivery *delivery) deleteProductFromCart(context echo.Context) error {
	var cartProduct domain.CartProducts
	err := json.NewDecoder(context.Request().Body).Decode(&cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err = delivery.MCDUsecase.DeleteProductFromCart(cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, "product is added to cart")
}

func (delivery *delivery) updateQuantityInCart(context echo.Context) error {
	var cartProduct domain.CartProducts
	err := json.NewDecoder(context.Request().Body).Decode(&cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err = delivery.MCDUsecase.UpdateQuantityInCart(cartProduct)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, "update successfull !")
}

func (delivery *delivery) getUserCart(context echo.Context) error {
	userId := context.QueryParam("userID")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	cart, err := delivery.MCDUsecase.GetUserCart(userIdInt)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, cart)
}

// Order-related handlers
func (delivery *delivery) CreateOrder(context echo.Context) error {
	var order domain.CreateOrderRequest
	err := json.NewDecoder(context.Request().Body).Decode(&order)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err = delivery.MCDUsecase.CreateOrder(order)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusCreated, "Order created successfully")
}

func (delivery *delivery) getUserOrders(context echo.Context) error {
	userID := context.Param("userID")
	if userID == "" {
		return context.JSON(http.StatusBadRequest, "userID is required")
	}
	userIdInt, err := strconv.Atoi(userID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	orders, err := delivery.MCDUsecase.GetUserOrders(userIdInt)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, orders)
}
