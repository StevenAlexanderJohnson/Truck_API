CREATE TABLE IF NOT EXISTS trucks.trucks (
    truck_id INT NOT NULL UNIQUE,
    truck_number INT AUTO_INCREMENT,
    truck_type int NOT NULL,
    truck_price int not NULL,
    PRIMARY KEY (truck_number)
);
