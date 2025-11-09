CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    age INT NOT NULL,
    gender VARCHAR(15) NOT NULL,
    location_lat FLOAT NOT NULL,
    location_lon 
);