CREATE TABLE IF NOT EXISTS "users" (
    user_id INT PRIMARY KEY,
    name VARCHAR(56) NOT NULL,
    surname VARCHAR(56) NOT NULL,
    patronymic VARCHAR(56) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(56) NOT NULL
);