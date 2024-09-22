-- Tabel vouchers
CREATE TABLE
  vouchers (
    voucher_id VARCHAR(255) PRIMARY KEY,
    discount DECIMAL(10, 2) NOT NULL,
    valid_until DATE,
    used BOOLEAN NOT NULL
  );

-- Tabel users
CREATE TABLE
  users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    birth_date DATE,
    address TEXT,
    contact_no VARCHAR(20)
  );

-- Tabel shoe_models
CREATE TABLE
  shoe_models (
    model_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL
  );

-- Tabel shoe_details
CREATE TABLE
  shoe_details (
    shoe_id INT AUTO_INCREMENT PRIMARY KEY,
    model_id INT NOT NULL,
    size INT NOT NULL,
    stock INT NOT NULL,
    FOREIGN KEY (model_id) REFERENCES shoe_models (model_id)
  );

-- Tabel carts
CREATE TABLE
  carts (
    cart_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    quantity INT NOT NULL,
    shoe_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
  );

-- Tabel orders
CREATE TABLE
  orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    voucher_id VARCHAR(255), -- changed to VARCHAR(255) to match vouchers table
    status VARCHAR(20),
    price INT,
    fee INT,
    discount INT,
    total_price INT,
    created_at DATETIME,
    updated_at DATETIME,
    metadata TEXT,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (voucher_id) REFERENCES vouchers (voucher_id)
  );

-- Tabel payments
CREATE TABLE
  payments (
    payment_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    payment_external_id VARCHAR(36),
    amount INT,
    status VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    metadata TEXT,
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
  );

-- Tabel deliveries
CREATE TABLE
  deliveries (
    delivery_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
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
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
  );

-- Tabel order_details
CREATE TABLE
  order_details (
    order_detail_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    shoe_id INT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
  );

-- INSERT DATA
-- Insert into vouchers
INSERT INTO
  vouchers (voucher_id, discount, valid_until, used)
VALUES
  ('VCH001', 10.00, '2024-12-31', FALSE),
  ('VCH002', 15.50, '2024-11-30', TRUE),
  ('VCH003', 20.00, '2024-10-31', FALSE),
  ('NOVOUCHER', 0.00, '2024-12-31', TRUE);

-- Insert into users
INSERT INTO
  users (
    email,
    password_hash,
    first_name,
    last_name,
    birth_date,
    address,
    contact_no
  )
VALUES
  (
    'john.doe@example.com',
    'hashed_password_123',
    'John',
    'Doe',
    '1990-01-01',
    '123 Main St, Cityville',
    '1234567890'
  ),
  (
    'jane.smith@example.com',
    'hashed_password_456',
    'Jane',
    'Smith',
    '1992-02-02',
    '456 Oak Ave, Townsville',
    '0987654321'
  );

-- Insert into shoe_models
INSERT INTO
  shoe_models (name, price)
VALUES
  ('Nike Air Max', 120),
  ('Adidas Ultraboost', 150),
  ('Puma Suede Classic', 80);

-- Insert into shoe_details
INSERT INTO
  shoe_details (model_id, size, stock)
VALUES
  (1, 42, 10), -- Nike Air Max, Size 42
  (1, 43, 15), -- Nike Air Max, Size 43
  (2, 41, 5), -- Adidas Ultraboost, Size 41
  (3, 40, 8);

-- Insert into carts
INSERT INTO
  carts (user_id, quantity, shoe_id)
VALUES
  (1, 2, 1), -- John Doe, 2x Nike Air Max, Size 42
  (2, 1, 4);

-- Insert into orders
INSERT INTO
  orders (
    user_id,
    voucher_id,
    status,
    price,
    fee,
    discount,
    total_price,
    metadata
  )
VALUES
  (
    1,
    'VCH001',
    'completed',
    120,
    10,
    10,
    120,
    '{"note": "first order"}'
  ),
  (
    2,
    'NOVOUCHER',
    'pending',
    150,
    15,
    0,
    165,
    '{"note": "gift purchase"}'
  );

-- Insert into payments
INSERT INTO
  payments (
    payment_id,
    order_id,
    payment_external_id,
    amount,
    status,
    metadata
  )
VALUES
  (
    1001,
    1,
    'EXT_001',
    120,
    'paid',
    '{"transaction_id": "TID001"}'
  ),
  (
    1002,
    2,
    'EXT_002',
    165,
    'pending',
    '{"transaction_id": "TID002"}'
  );

-- Insert into deliveries
INSERT INTO
  deliveries (
    delivery_id,
    order_id,
    delivery_date,
    arrival_date,
    courier_name,
    courier_service,
    weight_grams,
    origin_city_id,
    destination_city_id,
    delivery_fee,
    status,
    metadata
  )
VALUES
  (
    1,
    1,
    '2024-09-01',
    '2024-09-03',
    'DHL',
    'Express',
    2000,
    'C001',
    'C002',
    20,
    'delivered',
    '{"tracking_id": "TRACK001"}'
  ),
  (
    2,
    2,
    '2024-09-15',
    NULL,
    'FedEx',
    'Standard',
    1500,
    'C003',
    'C004',
    15,
    'in_transit',
    '{"tracking_id": "TRACK002"}'
  );

-- Insert into order_details
INSERT INTO
  order_details (order_id, shoe_id, quantity)
VALUES
  (1, 1, 2),
  (2, 4, 1);