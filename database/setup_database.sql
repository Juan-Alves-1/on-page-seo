-- Create the database
CREATE DATABASE friendly_url_results;

-- Use the database
USE friendly_url_results; 

-- Create the table to store URL analysis results
CREATE TABLE results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    keyword VARCHAR(255) NOT NULL,
    result TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
