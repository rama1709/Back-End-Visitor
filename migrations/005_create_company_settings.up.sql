CREATE TABLE company_settings (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,

    company_name VARCHAR(255) NOT NULL,
    company_email VARCHAR(255),
    company_phone VARCHAR(50),
    company_website VARCHAR(255),
    company_address TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO company_settings (
    company_name,
    company_email,
    company_phone,
    company_website,
    company_address
)
VALUES (
    'Visitor Management System',
    'frontdesk@company.com',
    '+62 812 3456 7890',
    'https://company.com',
    'Jl. Raya Teknologi No. 88, Jakarta'
);