package domain

import "time"

// User represents a user in the system.
type User struct {
	ID           int32  `json:"id,omitempty"`
	Email        string `json:"email"`
	PhoneNumber  uint   `json:"phone_number"` // Changed to uint to match repository
	PasswordHash string `json:"password_hash"`
	Name         string `json:"name"`
	Role         string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Name         string `json:"name"`
	Role         string `json:"role"`
}

// Product represents a product in the system.
type Product struct {
	ID        int32  `json:"id,omitempty"`
	Name      string `json:"name"`
	StockLeft int    `json:"stockLeft" gorm:"column:stockLeft"`
	HotelID   int    `json:"hotel_id"` // Hotel_id renamed to HotelID for consistency
	Category  string `json:"category"`
	Price     int    `json:"price"`
}

// Hotel represents a hotel in the system.
type Hotel struct {
	ID      int16  `json:"hotel,omitempty"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
	State   string `json:"state"`
}

// UserCart represents a cart for a user with a product and quantity.
type UserCart struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UserCartProduct struct {
	ID        int32  `json:"id,omitempty"`
	Name      string `json:"name"`
	StockLeft int    `json:"stockLeft" gorm:"column:stockLeft"`
	HotelID   int    `json:"hotel_id"` // Hotel_id renamed to HotelID for consistency
	Category  string `json:"category"`
	Price     int    `json:"price"`
	HotelName string `json:"hotel_name"`
}
type CartResponse struct {
	Products []UserCartProduct `json:"products"`
}

// CartProducts is a simplified version for checking if a user has certain products.
type CartProducts struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// CreateOrderRequest represents the structure for creating an order with product details.
type CreateOrderRequest struct {
	ID            int                   `gorm:"primaryKey" json:"id"`             // Order ID (optional for request, auto-generated in DB)
	UserID        int                   `json:"user_id"`                          // ID of the user placing the order
	PhoneNumber   string                `json:"phone_number"`                     // Phone number for the order
	DriveThruCode string                `json:"drive_thru_code,omitempty"`        // Optional drive-thru code
	OrderStatus   string                `json:"order_status"`                     // Status of the order (e.g., pending)
	Products      []OrderProductRequest `json:"products"`                         // List of products in the order
	IsDelivered   bool                  `json:"is_delivered"`                     // Delivery status
	OrderTotal    float64               `json:"order_total"`                      // Total price of the order
	CreatedAt     time.Time             `gorm:"autoCreateTime" json:"created_at"` // Timestamp when the order was created
	UpdatedAt     time.Time             `gorm:"autoUpdateTime" json:"updated_at"` // Timestamp when the order was last updated
}

// OrderProductRequest represents each product in an order.
type OrderProductRequest struct {
	ProductID       int     `json:"product_id"`        // ID of the product
	Quantity        int     `json:"quantity"`          // Quantity of the product
	PriceAtPurchase float64 `json:"price_at_purchase"` // Price of the product at the time of the order
}

type Order struct {
	ID            int            `gorm:"primaryKey" json:"id"`
	UserID        int            `json:"user_id"`
	PhoneNumber   string         `json:"phone_number"`
	DriveThruCode string         `json:"drive_thru_code,omitempty"`
	OrderStatus   string         `json:"order_status"`
	IsDelivered   bool           `json:"is_delivered"`
	OrderTotal    float64        `json:"order_total"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Products      []OrderProduct `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"products"` // Ensures cascading delete
}

type OrderProduct struct {
	ID              int     `gorm:"primaryKey" json:"id"` // Auto-incremented by MySQL
	OrderID         int     `json:"order_id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

type OrderResponse struct {
	ID            int               `json:"id"`
	UserID        int               `json:"user_id"`
	PhoneNumber   string            `json:"phone_number"`
	DriveThruCode string            `json:"drive_thru_code,omitempty"`
	OrderStatus   string            `json:"order_status"`
	IsDelivered   bool              `json:"is_delivered"`
	OrderTotal    float64           `json:"order_total"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Products      []UserCartProduct `json:"products"` // Ensures cascading delete
}

// CustomResponse represents a generic API response.
type CustomResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

// MCDUsecase defines the use case interface for managing users, orders, etc.
type MCDUsecase interface {
	// User CRUD operations
	CreateUser(user User) error
	UserLogin(userData UserLogin) (LoginResponse, error)
	UpdateUser(user User) error
	DeleteUser(userID string) error
	GetUserById(userID string) (User, error)

	// Product CRUD operations
	CreateProduct(product Product) error
	UpdateProduct(product Product) error
	DeleteProduct(productID string) error
	GetProductById(productID string) (Product, error)
	GetProductsByHotel(hotelID string) ([]Product, error)

	// Hotel CRUD operations
	CreateHotel(hotel Hotel) error
	UpdateHotel(hotel Hotel) error
	DeleteHotel(hotelID string) error
	GetHotels() ([]Hotel, error)

	AddProductToCart(CartProducts) error
	DeleteProductFromCart(CartProducts) error
	UpdateQuantityInCart(CartProducts) error
	GetUserCart(userID int) (CartResponse, error)

	// Order operations
	CreateOrder(order CreateOrderRequest) error
	//CancelOrder(orderID int) error
	//GetTodayOrders() ([]Order, error)
	GetUserOrders(phoneNumber int) ([]OrderResponse, error)
	MarkOrderCompleted(orderID int) error
}

// MCDRepository defines the repository interface for interacting with the database.
type MCDRepository interface {
	// User CRUD operations
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(userID string) error
	GetUserById(userID string) (User, error)
	GetUserByEmail(email string) (User, error)

	// Product CRUD operations
	CreateProduct(product Product) error
	UpdateProduct(product Product) error
	DeleteProduct(productID string) error
	GetProductById(productID string) (Product, error)
	GetProductsByHotel(hotelID string) ([]Product, error)

	// Hotel CRUD operations
	CreateHotel(hotel Hotel) error
	UpdateHotel(hotel Hotel) error
	DeleteHotel(hotelID string) error
	GetHotels() ([]Hotel, error)
	GetHotelByID(int) (*Hotel, error)

	AddProductToCart(CartProducts) error
	DeleteProductFromCart(CartProducts) error
	UpdateQuantityInCart(CartProducts) error
	GetProductDetails([]int) ([]Product, error)
	GetUserCart(userID int) ([]CartProducts, error)

	// Order operations
	CreateOrder(order Order) error
	//CancelOrder(orderID int) error
	//GetTodayOrders() ([]Order, error)
	GetUserOrders(phoneNumber int) ([]Order, error)
	MarkOrderCompleted(orderID int) error
}
