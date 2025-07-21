package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitializeDB(config *Config) (*sql.DB, error) {
	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }

	// // Construct connection string
	// connStr := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	config.DBHost,
	// 	config.DBPort,
	// 	config.DBUser,
	// 	config.DBPassword,
	// 	config.DBName,
	// )

	// // Open connection
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	return nil, err
	// }

	// // Verify connection
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	// log.Println("Successfully connected to database")
	return nil, nil
}
