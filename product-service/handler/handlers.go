package handler

import (
	"context"
	"fmt"
	"sync"

	pb "product-service/pb"
)

type server struct {
	pb.UnimplementedProductServiceServer
	products []*pb.Product
	mutex    sync.Mutex
	nextID   uint64
}

// Inisialisasi server
func newServer() *server {
	return &server{
		products: make([]*pb.Product, 0),
		nextID:   1,
	}
}

// Create product
func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Validate input
	if req.Name == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}
	if req.Price <= 0 {
		return nil, fmt.Errorf("product price must be greater than zero")
	}
	if req.Stock < 0 {
		return nil, fmt.Errorf("product stock cannot be negative")
	}
	if req.Category == nil {
		return nil, fmt.Errorf("category cannot be nil")
	}

	// Create the new product
	product := &pb.Product{
		Id:       s.nextID,
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: &pb.Category{Id: req.Category.Id, Name: req.Category.Name},
	}

	// Save the product to the database
	if err := s.db.CreateProduct(product); err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	// Increment the ID for the next product
	s.nextID++

	return &pb.ProductResponse{
		Products: []*pb.Product{product},
	}, nil
}

// Get product by ID
func (s *server) GetProductById(ctx context.Context, req *pb.ProductRequestById) (*pb.ProductResponse, error) {
	// Fetch the product from the database
	product, err := s.db.GetProductByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("product with ID %d not found: %w", req.Id, err)
	}

	return &pb.ProductResponse{
		Products: []*pb.Product{product},
	}, nil
}

// Get all products with pagination
func (s *server) GetAllProducts(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	// Fetch products from the database with pagination
	products, total, err := s.db.GetAllProducts(req.Page, req.PerPage)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}

	return &pb.ProductResponse{
		Products: products,
		Pagination: &pb.Pagination{
			Total:       uint64(total),
			PerPage:     req.PerPage,
			CurrentPage: req.Page,
			LastPage:    uint32((total + int(req.PerPage) - 1) / int(req.PerPage)),
		},
	}, nil
}

// Update product
func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	// Lock to prevent concurrent modifications
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Fetch the product from the database
	product, err := s.db.GetProductByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("product with ID %d not found: %w", req.Id, err)
	}

	// Update the product details
	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock
	product.Category = &pb.Category{
		Id:   req.Category.Id,
		Name: req.Category.Name,
	}

	// Save the updated product back to the database
	if err := s.db.UpdateProduct(product); err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return &pb.ProductResponse{
		Products: []*pb.Product{product},
	}, nil
}

// Delete product by ID
func (s *server) DeleteProduct(ctx context.Context, req *pb.ProductRequestById) (*pb.DeleteProductResponse, error) {
	// Lock to prevent concurrent modifications
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Attempt to delete the product from the database
	err := s.db.DeleteProduct(req.Id)
	if err != nil {
		return &pb.DeleteProductResponse{
			Success: false,
		}, fmt.Errorf("product with ID %d not found: %w", req.Id, err)
	}

	return &pb.DeleteProductResponse{
		Success: true,
	}, nil
}

// Delete all products
func (s *server) DeleteAllProducts(ctx context.Context, req *pb.EmptyRequest) (*pb.DeleteAllProductsResponse, error) {
	// Lock to prevent concurrent modifications
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Clear all products from the database
	if err := s.db.DeleteAllProducts(); err != nil {
		return nil, fmt.Errorf("error deleting all products: %w", err)
	}

	return &pb.DeleteAllProductsResponse{
		Success: true,
	}, nil
}
