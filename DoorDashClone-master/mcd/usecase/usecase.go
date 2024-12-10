package usecase

import (
	"fmt"
	"log"
	"mcd/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repository domain.MCDRepository
}

func NewUseCase(repository domain.MCDRepository) domain.MCDUsecase {
	return &usecase{repository: repository}
}

func generateToken(user_id int, email string, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"user_id":  user_id,
		"role":     role,
		"email":    email,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var signingKey interface{} = []byte("dinesh-bali-secret-key")
	signedToken, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Create User - Adds a new user
func (usecase *usecase) CreateUser(user domain.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(password)
	err = usecase.repository.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// login
func (usecase *usecase) UserLogin(userData domain.UserLogin) (domain.LoginResponse, error) {
	var loginResponse domain.LoginResponse
	user, err := usecase.repository.GetUserByEmail(userData.Email)
	if err != nil {
		return loginResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userData.Password))
	if err != nil {
		return loginResponse, err
	}
	token, err := generateToken(int(user.ID), user.Email, user.Name, user.Role)
	if err != nil {
		return loginResponse, err
	}
	loginResponse.Name = user.Name
	loginResponse.Role = user.Role
	loginResponse.Token = token
	return loginResponse, nil
}

// Delete User - Deletes a user by ID
func (usecase *usecase) DeleteUser(userID string) error {
	err := usecase.repository.DeleteUser(userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

// Update User - Updates an existing user
func (usecase *usecase) UpdateUser(user domain.User) error {
	err := usecase.repository.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

// Get User by ID - Fetches user by their ID
func (usecase *usecase) GetUserById(userID string) (domain.User, error) {
	user, err := usecase.repository.GetUserById(userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by id: %v", err)
	}
	return user, nil
}

// Create Product - Adds a new product
func (usecase *usecase) CreateProduct(product domain.Product) error {
	err := usecase.repository.CreateProduct(product)
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}
	return nil
}

// Delete Product - Deletes a product by ID
func (usecase *usecase) DeleteProduct(productID string) error {
	err := usecase.repository.DeleteProduct(productID)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

// Update Product - Updates an existing product
func (usecase *usecase) UpdateProduct(product domain.Product) error {
	err := usecase.repository.UpdateProduct(product)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}

// Get Product by ID - Fetches product by its ID
func (usecase *usecase) GetProductById(productID string) (domain.Product, error) {
	product, err := usecase.repository.GetProductById(productID)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to get product by id: %v", err)
	}
	return product, nil
}

// Get Products by Hotel - Fetches products for a specific hotel
func (usecase *usecase) GetProductsByHotel(hotelID string) ([]domain.Product, error) {
	products, err := usecase.repository.GetProductsByHotel(hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by hotel: %v", err)
	}
	return products, nil
}

// Create Hotel - Adds a new hotel
func (usecase *usecase) CreateHotel(hotel domain.Hotel) error {
	err := usecase.repository.CreateHotel(hotel)
	if err != nil {
		return fmt.Errorf("failed to create hotel: %v", err)
	}
	return nil
}

// Delete Hotel - Deletes a hotel by ID
func (usecase *usecase) DeleteHotel(hotelID string) error {
	err := usecase.repository.DeleteHotel(hotelID)
	if err != nil {
		return fmt.Errorf("failed to delete hotel: %v", err)
	}
	return nil
}

// Update Hotel - Updates an existing hotel
func (usecase *usecase) UpdateHotel(hotel domain.Hotel) error {
	err := usecase.repository.UpdateHotel(hotel)
	if err != nil {
		return fmt.Errorf("failed to update hotel: %v", err)
	}
	return nil
}

// Get Hotels - Fetches all hotels
func (usecase *usecase) GetHotels() ([]domain.Hotel, error) {
	hotels, err := usecase.repository.GetHotels()
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %v", err)
	}
	return hotels, nil
}

func (usecase *usecase) AddProductToCart(cartProduct domain.CartProducts) error {
	err := usecase.repository.AddProductToCart(cartProduct)
	if err != nil {
		log.Printf("Error adding product to cart: %v", err)
		return err
	}
	return nil
}

// DeleteProductFromCart removes a product from the user's cart.
func (usecase *usecase) DeleteProductFromCart(cartProduct domain.CartProducts) error {
	err := usecase.repository.DeleteProductFromCart(cartProduct)
	if err != nil {
		log.Printf("Error deleting product from cart: %v", err)
		return err
	}
	return nil
}

// UpdateQuantityInCart updates the quantity of a product in the user's cart.
func (usecase *usecase) UpdateQuantityInCart(cartProduct domain.CartProducts) error {
	err := usecase.repository.UpdateQuantityInCart(cartProduct)
	if err != nil {
		log.Printf("Error updating quantity in cart: %v", err)
		return err
	}
	return nil
}

// GetUserCart retrieves all products in the user's cart.
func (usecase *usecase) GetUserCart(userID int) (domain.CartResponse, error) {
	cart, err := usecase.repository.GetUserCart(userID)
	if err != nil {
		log.Printf("Error getting user cart: %v", err)
		return domain.CartResponse{}, err
	}
	var cartResponse domain.CartResponse
	var productIDS []int
	for _, item := range cart {
		productIDS = append(productIDS, item.ProductID)
	}
	//var products []domain.Product
	products, err := usecase.repository.GetProductDetails(productIDS)
	if err != nil {
		return domain.CartResponse{}, err
	}
	var cartProducts []domain.UserCartProduct
	for _, cartItem := range products {
		var cartProduct domain.UserCartProduct
		hotel, err := usecase.repository.GetHotelByID(cartItem.HotelID)
		if err != nil {
			continue
		}
		cartProduct.ID = cartItem.ID
		cartProduct.Category = cartItem.Category
		cartProduct.Name = cartItem.Name
		cartProduct.HotelName = hotel.Name
		cartProduct.Price = cartItem.Price
		cartProduct.StockLeft = cartItem.StockLeft

		cartProducts = append(cartProducts, cartProduct)
	}

	cartResponse.Products = cartProducts

	return cartResponse, nil
}

// Create Order - Adds a new order
func (usecase *usecase) CreateOrder(order domain.CreateOrderRequest) error {
	var db_order domain.Order
	db_order.UserID = order.UserID
	db_order.PhoneNumber = order.PhoneNumber
	db_order.DriveThruCode = order.DriveThruCode
	db_order.OrderStatus = order.OrderStatus
	db_order.IsDelivered = order.IsDelivered
	db_order.OrderTotal = order.OrderTotal

	for i := 0; i < len(order.Products); i++ {
		var orderProduct domain.OrderProduct
		orderProduct.OrderID = order.ID
		orderProduct.PriceAtPurchase = order.Products[i].PriceAtPurchase
		orderProduct.ProductID = order.Products[i].ProductID
		orderProduct.Quantity = order.Products[i].Quantity
		orderProduct.PriceAtPurchase = order.Products[i].PriceAtPurchase
		db_order.Products = append(db_order.Products, orderProduct)
	}
	err := usecase.repository.CreateOrder(db_order)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}
	return nil
}

// Mark Order Completed - Marks an order as completed
func (usecase *usecase) MarkOrderCompleted(orderID int) error {
	err := usecase.repository.MarkOrderCompleted(orderID)
	if err != nil {
		return fmt.Errorf("failed to mark order as completed: %v", err)
	}
	return nil
}

// Get User Orders - Fetches all orders for a user based on user ID
func (usecase *usecase) GetUserOrders(userID int) ([]domain.OrderResponse, error) {
	orders, err := usecase.repository.GetUserOrders(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user orders: %v", err)
	}

	var orderResponses []domain.OrderResponse

	// Iterate over each order to build the OrderResponse
	for _, order := range orders {
		var productIDS []int
		for _, item := range order.Products {
			productIDS = append(productIDS, item.ProductID)
		}

		// Get product details for the current order
		products, err := usecase.repository.GetProductDetails(productIDS)
		if err != nil {
			return []domain.OrderResponse{}, fmt.Errorf("failed to get product details: %v", err)
		}

		var cartProducts []domain.UserCartProduct
		for _, cartItem := range products {
			var cartProduct domain.UserCartProduct
			hotel, err := usecase.repository.GetHotelByID(cartItem.HotelID)
			if err != nil {
				continue
			}
			cartProduct.ID = cartItem.ID
			cartProduct.Category = cartItem.Category
			cartProduct.Name = cartItem.Name
			cartProduct.HotelName = hotel.Name
			cartProduct.Price = cartItem.Price
			cartProduct.StockLeft = cartItem.StockLeft

			cartProducts = append(cartProducts, cartProduct)
		}

		// Build the OrderResponse for the current order
		orderResponse := domain.OrderResponse{
			ID:          order.ID,
			PhoneNumber: order.PhoneNumber,
			OrderTotal:  order.OrderTotal,
			OrderStatus: order.OrderStatus,
			IsDelivered: order.IsDelivered,
			Products:    cartProducts,
		}

		// Append the constructed OrderResponse to the list
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}
