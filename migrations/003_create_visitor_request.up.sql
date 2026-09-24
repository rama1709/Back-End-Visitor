CREATE TABLE visitor_request (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    visitor_id BIGINT NOT NULL,
    employee_id BIGINT NOT NULL,
    purpose TEXT,
    visit_date DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'Pending',
    check_in TIMESTAMP NULL,
    check_out TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_visitor
        FOREIGN KEY(visitor_id)
        REFERENCES visitor(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_employee
        FOREIGN KEY(employee_id)
        REFERENCES employee(id)
        ON DELETE CASCADE
);