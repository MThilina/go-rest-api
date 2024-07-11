-- Create the user if it doesn't already exist
CREATE USER IF NOT EXISTS 'app_user'@'%' IDENTIFIED BY 'root';

-- Grant all privileges to the user on all databases with grant option
GRANT ALL PRIVILEGES ON *.* TO 'app_user'@'%' WITH GRANT OPTION;

-- Flush the privileges to ensure changes are applied
FLUSH PRIVILEGES;
