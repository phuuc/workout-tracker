	CREATE USER 'wt'@'%%' IDENTIFIED BY 'wt';
    GRANT ALL PRIVILEGES ON * . * TO 'wt'@'%%';
    FLUSH PRIVILEGES;

    CREATE DATABASE IF NOT EXISTS `workout-tracker`;
    USE `workout-tracker`;