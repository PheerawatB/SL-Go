-- Create mobiles table
CREATE TABLE mobiles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price VARCHAR(100),
    product_number VARCHAR(100),
    details TEXT,
    produce_date TIMESTAMP,
    guarantee_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL  -- Soft delete timestamp
);

-- Create trading_details table
CREATE TABLE trading_details (
    id SERIAL PRIMARY KEY,
    mobile_id INT REFERENCES mobiles(id),
    user_id INT REFERENCES users(id),
    shop_id INT REFERENCES shops(id),
    discount INT,
    total_price FLOAT,
    total_discount FLOAT,
    details TEXT,
    guarantee_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL  -- Soft delete timestamp
);

-- Create table for User
CREATE TABLE users (
    id SERIAL PRIMARY KEY,          -- Automatically incrementing ID
    name VARCHAR(100) NOT NULL,     -- User's name
    username VARCHAR(100) NOT NULL, -- Username
    email VARCHAR(100) NOT NULL,    -- Email
    phone_no VARCHAR(20),           -- Phone number
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Update timestamp
    deleted_at TIMESTAMP DEFAULT NULL  -- Soft delete timestamp
);

-- Create table for UserType
CREATE TABLE user_types (
    id SERIAL PRIMARY KEY,               -- Automatically incrementing ID
    user_id INT REFERENCES users(id),    -- Foreign key referencing users
    name VARCHAR(100) NOT NULL,          -- Name of user type
    description TEXT,                    -- Description of the user type
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Update timestamp
    deleted_at TIMESTAMP DEFAULT NULL  -- Soft delete timestamp
);

-- Create table for Shop
CREATE TABLE shops (
    id SERIAL PRIMARY KEY,                     -- Automatically incrementing ID
    name VARCHAR(100) NOT NULL,                -- Shop name
    user_id INT REFERENCES users(id),          -- Foreign key referencing users
    user_type_id INT REFERENCES user_types(id),-- Foreign key referencing user_types
    description TEXT,                          -- Shop description
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Update timestamp
    deleted_at TIMESTAMP DEFAULT NULL  -- Soft delete timestamp
);
