package mysql

import (
	"context"
	"fmt"
	"log"
	"mcd/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.MCDRepository {
	return &repository{db: db}
}

// CreateUser - Adds a new user to the database
func (r *repository) CreateUser(user domain.User) error {
	err := r.db.WithContext(context.Background()).Table("users").Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// DeleteUser - Deletes a user from the database by ID
func (r *repository) DeleteUser(userID string) error {
	err := r.db.WithContext(context.Background()).Table("users").Where("id = ?", userID).Delete(&domain.User{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// UpdateUser - Updates an existing user's information
func (r *repository) UpdateUser(user domain.User) error {
	err := r.db.WithContext(context.Background()).Table("users").Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// GetUserById - Fetches a user from the database by their ID
func (r *repository) GetUserById(userID string) (domain.User, error) {
	var user domain.User
	err := r.db.WithContext(context.Background()).Table("users").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.WithContext(context.Background()).Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

// CreateProduct - Adds a new product to the database
func (r *repository) CreateProduct(product domain.Product) error {
	err := r.db.WithContext(context.Background()).Table("products").Create(&product).Error
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

// DeleteProduct - Deletes a product from the database by ID
func (r *repository) DeleteProduct(productID string) error {
	err := r.db.WithContext(context.Background()).Table("products").Where("id = ?", productID).Delete(&domain.Product{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

// UpdateProduct - Updates an existing product's information
func (r *repository) UpdateProduct(product domain.Product) error {
	err := r.db.WithContext(context.Background()).Table("products").Where("id = ?", product.ID).Updates(product).Error
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

// GetProductById - Fetches a product by its ID
func (r *repository) GetProductById(productID string) (domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(context.Background()).Table("products").Where("id = ?", productID).First(&product).Error
	if err != nil {
		return product, fmt.Errorf("failed to get product by id: %w", err)
	}
	return product, nil
}

// GetProductsByHotel - Fetches all products for a specific hotel
func (r *repository) GetProductsByHotel(hotelID string) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(context.Background()).Table("products").Where("hotel_id = ?", hotelID).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get products by hotel: %w", err)
	}
	return products, nil
}

// CreateHotel - Adds a new hotel to the database
func (r *repository) CreateHotel(hotel domain.Hotel) error {
	err := r.db.WithContext(context.Background()).Table("hotels").Create(&hotel).Error
	if err != nil {
		return fmt.Errorf("failed to create hotel: %w", err)
	}
	return nil
}

// DeleteHotel - Deletes a hotel from the database by ID
func (r *repository) DeleteHotel(hotelID string) error {
	err := r.db.WithContext(context.Background()).Table("hotels").Where("id = ?", hotelID).Delete(&domain.Hotel{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete hotel: %w", err)
	}
	return nil
}

// UpdateHotel - Updates an existing hotel's information
func (r *repository) UpdateHotel(hotel domain.Hotel) error {
	err := r.db.WithContext(context.Background()).Table("hotels").Where("id = ?", hotel.ID).Updates(hotel).Error
	if err != nil {
		return fmt.Errorf("failed to update hotel: %w", err)
	}
	return nil
}

// GetHotels - Fetches all hotels from the database
func (r *repository) GetHotels() ([]domain.Hotel, error) {
	var hotels []domain.Hotel
	err := r.db.WithContext(context.Background()).Table("hotels").Find(&hotels).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %w", err)
	}
	return hotels, nil
}

// GetHotelByID - Fetches a hotel by its ID from the database
func (r *repository) GetHotelByID(hotelID int) (*domain.Hotel, error) {
	var hotel domain.Hotel
	err := r.db.WithContext(context.Background()).
		Table("hotels").
		Where("id = ?", hotelID).
		First(&hotel).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel with ID %d: %w", hotelID, err)
	}
	return &hotel, nil
}

func (r *repository) AddProductToCart(cartProduct domain.CartProducts) error {
	query := `INSERT INTO user_carts (user_id, product_id, quantity) 
              VALUES (?, ?, ?);`

	tx := r.db.Exec(query, cartProduct.UserID, cartProduct.ProductID, cartProduct.Quantity)
	if tx.Error != nil {
		log.Printf("Error adding product to cart: %v", tx.Error)
		return tx.Error
	}
	return nil
}

// DeleteProductFromCart removes a product from the user's cart.
func (r *repository) DeleteProductFromCart(cartProduct domain.CartProducts) error {
	query := `DELETE FROM user_carts WHERE user_id = ? AND product_id = ?;`

	tx := r.db.Exec(query, cartProduct.UserID, cartProduct.ProductID)
	if tx.Error != nil {
		log.Printf("Error deleting product from cart: %v", tx.Error)
		return tx.Error
	}

	return nil
}

func (r *repository) UpdateQuantityInCart(cartProduct domain.CartProducts) error {
	query := `UPDATE user_carts 
              SET quantity = ? 
              WHERE user_id = ? AND product_id = ?;`

	tx := r.db.Exec(query, cartProduct.Quantity, cartProduct.UserID, cartProduct.ProductID)
	if tx.Error != nil {
		log.Printf("Error updating quantity in cart: %v", tx.Error)
		return tx.Error
	}

	return nil
}

// GetUserCart retrieves all products in the user's cart.
func (r *repository) GetUserCart(userID int) ([]domain.CartProducts, error) {
	// Query to select all products in the user's cart
	query := `SELECT user_id, product_id, quantity 
              FROM user_carts WHERE user_id = ?`

	// Use the GORM Query method to execute the SQL query
	rows, err := r.db.Raw(query, userID).Rows()
	if err != nil {
		log.Printf("Error getting user cart: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Prepare a slice to store cart products
	var cart []domain.CartProducts

	// Iterate over the rows and scan the data into CartProduct structs
	for rows.Next() {
		var cartProduct domain.CartProducts
		if err := rows.Scan(&cartProduct.UserID, &cartProduct.ProductID, &cartProduct.Quantity); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		cart = append(cart, cartProduct)
	}

	// Return the cart
	return cart, nil
}

// GetProductDetails retrieves details for a list of product IDs.
func (r *repository) GetProductDetails(productIDs []int) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(context.Background()).
		Table("products").
		Where("id IN ?", productIDs).
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product details: %w", err)
	}
	return products, nil
}

func (r *repository) CreateOrder(order domain.Order) error {
	ctx := context.Background()

	// Start a database transaction to ensure atomicity
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	// Insert the main order record into the user_orders table
	if err := tx.Table("user_orders").Create(&order).Error; err != nil {
		fmt.Println("ERRRRR:", err)
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("failed to create order: %w", err)
	}

	// Log the order and its products for debugging
	fmt.Println("ORDER::", order.Products)

	// Assign the generated OrderID to each product in the order
	for i := range order.Products {
		// Ensure ID is reset for auto-generation
		order.Products[i].ID = 0
		order.Products[i].OrderID = order.ID
	}

	// Log the products being inserted for debugging
	fmt.Printf("Products to Insert: %+v\n", order.Products)
	fmt.Printf("Products Count: %d\n", len(order.Products))

	// Insert the products into the order_products table
	// if len(order.Products) > 0 {
	// 	if err := tx.Table("order_products").Create(&order.Products).Error; err != nil {
	// 		fmt.Println("QQQQQQQQQQQQQQ", err)
	// 		tx.Rollback() // Roll back the transaction on error
	// 		return fmt.Errorf("failed to create order products: %w", err)
	// 	}
	// }

	// Commit the transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// MarkOrderCompleted - Marks an order as completed
func (r *repository) MarkOrderCompleted(orderID int) error {
	err := r.db.WithContext(context.Background()).Table("orders").Where("id = ?", orderID).Update("order_completed", true).Error
	if err != nil {
		return fmt.Errorf("failed to mark order as completed: %w", err)
	}
	return nil
}

// GetUserOrders - Fetches all orders for a user based on user ID
func (r *repository) GetUserOrders(userID int) ([]domain.Order, error) {
	var orders []domain.Order

	// Use Preload to load the associated products for each order
	err := r.db.WithContext(context.Background()).Table("user_orders").
		Preload("Products"). // Match the field name in the Order struct
		Where("user_id = ?", userID).
		Find(&orders).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}

	return orders, nil
}
