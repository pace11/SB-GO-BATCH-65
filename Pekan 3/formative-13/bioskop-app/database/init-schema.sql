CREATE TABLE IF NOT EXISTS bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    rating REAL
);
