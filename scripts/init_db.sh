#!/bin/bash

docker run --name postgres-go -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v postgres-data:/var/lib/postgresql/data -p 5433:5432 -d postgres
echo "Postgresql starting..."
sleep 3

docker exec -it postgres-go psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3
echo "Database productapp created"

docker exec -it postgres-go psql -U postgres -d productapp -c "
CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DOUBLE PRECISION NOT NULL,
  discount DOUBLE PRECISION,
  store VARCHAR(255) NOT NULL
);"
sleep 3
echo "Table products created"