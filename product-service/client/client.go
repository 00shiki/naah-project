package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "product-service/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create product
	createResp, err := c.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:  "New Product",
		Price: 100.0,
		Stock: 10,
		Category: &pb.Category{
			Id:   1,
			Name: "Category A",
		},
	})
	if err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
	fmt.Printf("Created Product: %v\n", createResp.Products)

	// Get all products
	allProducts, err := c.GetAllProducts(ctx, &pb.ProductRequest{
		Page:    1,
		PerPage: 5,
	})
	if err != nil {
		log.Fatalf("Error getting products: %v", err)
	}
	fmt.Printf("All Products: %v\n", allProducts.Products)

	// Get product by ID
	getResp, err := c.GetProductById(ctx, &pb.ProductRequestById{Id: createResp.Products[0].Id})
	if err != nil {
		log.Fatalf("Error getting product by ID: %v", err)
	}
	fmt.Printf("Get Product By ID: %v\n", getResp.Products)

	// Update product
	updateResp, err := c.UpdateProduct(ctx, &pb.UpdateProductRequest{
		Id:    createResp.Products[0].Id,
		Name:  "Updated Product",
		Price: 150.0,
		Stock: 20,
		Category: &pb.Category{
			Id:   1,
			Name: "Category Updated",
		},
	})
	if err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
	fmt.Printf("Updated Product: %v\n", updateResp.Products)

	// Delete product
	// delResp, err := c.DeleteProduct(ctx, &pb.ProductRequestById{Id: createResp.Products[0].Id})
	// if err != nil {
	// 	log.Fatalf("Error deleting product: %v", err)
	// }
	// fmt.Printf("Delete Product Success: %v\n", delResp.Success)

	// Delete all products
	deleteAllResp, err := c.DeleteAllProducts(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("Error deleting all products: %v", err)
	}
	fmt.Printf("Delete All Products Success: %v\n", deleteAllResp.Success)
}
