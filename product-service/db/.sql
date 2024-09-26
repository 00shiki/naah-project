-- Tabel vouchers
CREATE TABLE
  vouchers (
    voucher_id VARCHAR(255) PRIMARY KEY,
    discount DECIMAL(10, 2) NOT NULL,
    valid_until DATE,
    used BOOLEAN NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
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
    contact_no VARCHAR(20),
    role INT NOT NULL, -- Added role column as an integer
    verified BOOLEAN DEFAULT FALSE, -- Added verified column with default value of false
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- Tabel shoe_models
CREATE TABLE
  shoe_models (
    model_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- Tabel shoe_details
CREATE TABLE
  shoe_details (
    shoe_id INT AUTO_INCREMENT PRIMARY KEY,
    model_id INT NOT NULL,
    size INT NOT NULL,
    stock INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (model_id) REFERENCES shoe_models (model_id)
  );

-- Tabel carts
CREATE TABLE
  carts (
    cart_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    quantity INT NOT NULL,
    shoe_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
  );

-- Tabel orders
CREATE TABLE
  orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    voucher_id VARCHAR(255),
    status VARCHAR(20),
    price INT,
    fee INT,
    discount INT,
    total_price INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
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
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    metadata TEXT,
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
  );

-- Tabel deliveries
CREATE TABLE
  deliveries (
    delivery_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
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
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
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
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
  );

-- Insert into users
INSERT INTO
  users (
    email,
    password_hash,
    first_name,
    last_name,
    birth_date,
    address,
    contact_no,
    role, -- Added role field
    verified -- Added verified field
  )
VALUES
  (
    'john.doe@example.com',
    'hashed_password_123',
    'John',
    'Doe',
    '1990-01-01',
    '123 Main St, Cityville',
    '1234567890',
    1, -- Example role value for John
    TRUE -- Example verified status for John
  ),
  (
    'jane.smith@example.com',
    'hashed_password_456',
    'Jane',
    'Smith',
    '1992-02-02',
    '456 Oak Ave, Townsville',
    '0987654321',
    2, -- Example role value for Jane
    FALSE -- Example verified status for Jane
  );

-- Insert into shoe_models
INSERT INTO
  shoe_models (name, price)
VALUES
  ('Nike Air Max', 1200000),
  ('Adidas Ultraboost', 1500000),
  ('Puma Suede Classic', 800000);

-- Insert into shoe_details
INSERT INTO
  shoe_details (model_id, size, stock)
VALUES
  (1, 42, 10), -- Nike Air Max, Size 42
  (1, 43, 15), -- Nike Air Max, Size 43
  (2, 41, 5), -- Adidas Ultraboost, Size 41
  (3, 40, 8);

-- Puma Suede Classic, Size 40
-- Insert into carts
INSERT INTO
  carts (user_id, quantity, shoe_id)
VALUES
  (1, 2, 1), -- John Doe, 2x Nike Air Max, Size 42
  (2, 1, 4);

-- Jane Smith, 1x Puma Suede Classic, Size 40