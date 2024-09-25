package main

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

	product := &pb.Product{
		Id:    s.nextID,
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
		Category: &pb.Category{
			Id:   req.Category.Id,
			Name: req.Category.Name,
		},
	}
	s.products = append(s.products, product)
	s.nextID++

	return &pb.ProductResponse{
		Products: []*pb.Product{product},
	}, nil
}

// Get product by ID
func (s *server) GetProductById(ctx context.Context, req *pb.ProductRequestById) (*pb.ProductResponse, error) {
	for _, product := range s.products {
		if product.Id == req.Id {
			return &pb.ProductResponse{
				Products: []*pb.Product{product},
			}, nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", req.Id)
}

// Get all products with pagination
func (s *server) GetAllProducts(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	start := int(req.Page-1) * int(req.PerPage)
	end := start + int(req.PerPage)

	if start >= len(s.products) {
		return &pb.ProductResponse{
			Products: []*pb.Product{},
		}, nil
	}

	if end > len(s.products) {
		end = len(s.products)
	}

	return &pb.ProductResponse{
		Products: s.products[start:end],
		Pagination: &pb.Pagination{
			Total:       uint64(len(s.products)),
			PerPage:     req.PerPage,
			CurrentPage: req.Page,
			LastPage:    uint32((len(s.products) + int(req.PerPage) - 1) / int(req.PerPage)),
		},
	}, nil
}

// Update product
func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	for _, product := range s.products {
		if product.Id == req.Id {
			product.Name = req.Name
			product.Price = req.Price
			product.Stock = req.Stock
			product.Category = &pb.Category{
				Id:   req.Category.Id,
				Name: req.Category.Name,
			}
			return &pb.ProductResponse{
				Products: []*pb.Product{product},
			}, nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", req.Id)
}

// Delete product by ID
func (s *server) DeleteProduct(ctx context.Context, req *pb.ProductRequestById) (*pb.DeleteProductResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, product := range s.products {
		if product.Id == req.Id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			return &pb.DeleteProductResponse{
				Success: true,
			}, nil
		}
	}
	return &pb.DeleteProductResponse{
		Success: false,
	}, fmt.Errorf("product with ID %d not found", req.Id)
}

// Delete all products
func (s *server) DeleteAllProducts(ctx context.Context, req *pb.EmptyRequest) (*pb.DeleteAllProductsResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.products = []*pb.Product{}
	return &pb.DeleteAllProductsResponse{
		Success: true,
	}, nil
}
