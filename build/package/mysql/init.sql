CREATE DATABASE IF NOT EXISTS xyz_multifinance;

DROP USER IF EXISTS 'xyz_multi'@'%';
CREATE USER 'xyz_multi'@'%' IDENTIFIED BY 'mysql123';
GRANT ALL PRIVILEGES ON xyz_multifinance.* TO 'xyz_multi'@'%';

DROP USER IF EXISTS 'xyz_multi'@'localhost';
CREATE USER 'xyz_multi'@'localhost' IDENTIFIED BY 'mysql123';
GRANT ALL PRIVILEGES ON xyz_multifinance.* TO 'xyz_multi'@'localhost';

FLUSH PRIVILEGES;
