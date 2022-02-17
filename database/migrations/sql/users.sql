DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id VARCHAR(50) PRIMARY KEY,
	avatar VARCHAR(50) NOT NULL,
	username VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(20) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	name VARCHAR(50) NOT NULL,
	phone VARCHAR(20) UNIQUE NOT NULL,
    role VARCHAR(20) DEFAULT "user" NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	deleted_at TIMESTAMP
);