
CREATE TABLE IF NOT EXISTS Users (
    id uuid UNIQUE NOT NULL,
    nome varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    age varchar(15) NOT NULL
);
