CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    ktp VARCHAR(255) NOT NULL UNIQUE,
    status ENUM('active', 'inactive') NOT NULL
);


CREATE TABLE exams (
    id INT PRIMARY KEY AUTO_INCREMENT,
    employee_id INT NOT NULL,
    score INT,
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

-- Insert dummy employees
INSERT INTO employees (name, ktp, status) VALUES
    ('John Doe', '1234567890', 'active'),
    ('Jane Smith', '0987654321', 'active'),
    ('Michael Johnson', '9876543210', 'inactive'),
    ('Emily Williams', '4567890123', 'active'),
    ('David Brown', '7890123456', 'inactive');

-- Insert dummy exams
INSERT INTO exams (employee_id, score) VALUES
    (1, 85),
    (2, 92),
    (3, 78),
    (4, 90),
    (5, 82);
