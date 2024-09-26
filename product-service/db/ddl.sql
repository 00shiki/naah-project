-- Drop and recreate the database
DROP DATABASE IF EXISTS railway;
CREATE DATABASE railway;
USE railway;

-- Table vouchers
CREATE TABLE vouchers (
  voucher_id VARCHAR(255) PRIMARY KEY,
  discount DECIMAL(10, 2) NOT NULL,
  valid_until DATE,
  used BOOLEAN NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
);

-- Table users
CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  birth_date DATE,
  address TEXT,
  contact_no VARCHAR(20),
  role INT NOT NULL, 
  verified BOOLEAN DEFAULT FALSE, 
  created_at DATETIME,
  updated_at DATETIME
);

-- Table shoe_models
CREATE TABLE shoe_models (
  model_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
);

-- Table shoe_details
CREATE TABLE shoe_details (
  shoe_id INT AUTO_INCREMENT PRIMARY KEY,
  sd_model_id INT NOT NULL,
  size INT NOT NULL,
  stock INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  CONSTRAINT fk_shoe_details_model_id FOREIGN KEY (sd_model_id) REFERENCES shoe_models (model_id)
);

-- Table carts
CREATE TABLE carts (
  cart_id INT AUTO_INCREMENT PRIMARY KEY,
  c_user_id INT NOT NULL,
  quantity INT NOT NULL,
  c_shoe_id INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  CONSTRAINT fk_carts_user_id FOREIGN KEY (c_user_id) REFERENCES users (user_id),
  CONSTRAINT fk_carts_shoe_id FOREIGN KEY (c_shoe_id) REFERENCES shoe_details (shoe_id)
);

-- Table orders
CREATE TABLE orders (
  order_id INT AUTO_INCREMENT PRIMARY KEY,
  o_user_id INT,
  o_voucher_id VARCHAR(255),
  status VARCHAR(20),
  price INT,
  fee INT,
  discount INT,
  total_price INT,
  created_at DATETIME,
  updated_at DATETIME,
  metadata TEXT,
  CONSTRAINT fk_orders_user_id FOREIGN KEY (o_user_id) REFERENCES users (user_id),
  CONSTRAINT fk_orders_voucher_id FOREIGN KEY (o_voucher_id) REFERENCES vouchers (voucher_id)
);

-- Table payments
CREATE TABLE payments (
  payment_id INT AUTO_INCREMENT PRIMARY KEY,
  p_order_id INT,
  payment_external_id VARCHAR(36),
  amount INT,
  status VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME,
  metadata TEXT,
  CONSTRAINT fk_payments_order_id FOREIGN KEY (p_order_id) REFERENCES orders (order_id)
);

-- Table deliveries
CREATE TABLE deliveries (
  delivery_id INT AUTO_INCREMENT PRIMARY KEY,
  d_order_id INT,
  track_id VARCHAR(225),
  delivery_date DATETIME,
  arrival_date DATETIME,
  courier_name VARCHAR(50),
  courier_service VARCHAR(100),
  weight_grams INT,
  origin_city_id VARCHAR(50),
  destination_city_id VARCHAR(50),
  delivery_fee INT,
  status VARCHAR(20),
  created_at DATETIME,
  updated_at DATETIME,
  metadata TEXT,
  CONSTRAINT fk_deliveries_order_id FOREIGN KEY (d_order_id) REFERENCES orders (order_id)
);

-- Table order_details
CREATE TABLE order_details (
  order_detail_id INT AUTO_INCREMENT PRIMARY KEY,
  od_order_id INT NOT NULL,
  od_shoe_id INT NOT NULL,
  quantity INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  CONSTRAINT fk_order_details_order_id FOREIGN KEY (od_order_id) REFERENCES orders (order_id),
  CONSTRAINT fk_order_details_shoe_id FOREIGN KEY (od_shoe_id) REFERENCES shoe_details (shoe_id)
);

-- Insert data into users
INSERT INTO users (
  email, password_hash, first_name, last_name, birth_date, address, contact_no, role, verified
) VALUES
  ('john.doe@example.com', 'hashed_password_123', 'John', 'Doe', '1990-01-01', '123 Main St, Cityville', '1234567890', 1, TRUE),
  ('jane.smith@example.com', 'hashed_password_456', 'Jane', 'Smith', '1992-02-02', '456 Oak Ave, Townsville', '0987654321', 2, FALSE);

-- Insert data into shoe_models
INSERT INTO shoe_models (name, price) VALUES
  ('Nike Air Max', 1200000),
  ('Adidas Ultraboost', 1500000),
  ('Puma Suede Classic', 800000);

-- Insert data into shoe_details
INSERT INTO shoe_details (sd_model_id, size, stock) VALUES
  (1, 42, 10), 
  (1, 43, 15), 
  (2, 41, 5), 
  (3, 40, 8);

-- Insert data into carts
INSERT INTO carts (c_user_id, quantity, c_shoe_id) VALUES
  (1, 2, 1), 
  (2, 1, 4);

-- Insert data into vouchers
INSERT INTO vouchers (voucher_id, discount, valid_until, used) VALUES
  ('VCH001', 10.00, '2024-12-31', FALSE),
  ('VCH002', 15.50, '2024-11-30', TRUE),
  ('VCH003', 20.00, '2024-10-31', FALSE),
  ('NOVOUCHER', 0.00, '9999-12-31', FALSE);
