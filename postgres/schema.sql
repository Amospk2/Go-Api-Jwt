
CREATE TABLE IF NOT EXISTS Users (
    id uuid UNIQUE NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    age varchar(15) NOT NULL,
    password varchar(255) NOT NULL
);
