DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS messages;

CREATE TABLE users (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`  VARCHAR(100) NOT NULL,
    `password` VARCHAR(100) NOT NULL
);

CREATE TABLE sessions (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid` VARCHAR(64) NOT NULL UNIQUE,
    `name`  VARCHAR(100) NOT NULL
);

CREATE TABLE rooms (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `room_name`  VARCHAR(100) NOT NULL,
    `update_time` datetime
);

CREATE TABLE messages (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100),
    `room_id` INTEGER REFERENCES rooms(id),
    `text` VARCHAR(255),
    `date` date,
    `time` time,
    `reply_id` int
);