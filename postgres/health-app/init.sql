CREATE TABLE IF NOT EXISTS hospital (
    id SERIAL PRIMARY KEY,
    hcode VARCHAR(5) UNIQUE NOT NULL,              
    name VARCHAR(255) NOT NULL,
    his_base_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    hcode VARCHAR(5),                             
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_hospital_code 
        FOREIGN KEY (hcode) 
        REFERENCES hospital(hcode) 
        ON DELETE CASCADE
);


INSERT INTO hospital (
    hcode, name, his_base_url
) VALUES
('00001', 'โรงพยาบาลทดสอบ1', 'http://nginx/his'),
('00002', 'โรงพยาบาลทดสอบ2', 'http://nginx/his2');