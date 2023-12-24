package postgres

import (
	"fmt" // Used to format the DSN string.

	"gorm.io/driver/postgres" // Postgres connection driver
	"gorm.io/gorm"            // Postgres ORM
)

/*
DBConnect establishes a connection to a PostgreSQL database
Parameters:
  - host: The host of the database.
  - port: The port of the database.
  - username: The username for authentication.
  - password: The password for authentication.
  - dbname: The name of the database.
  - models: Optional variadic parameter for specifying GORM models to auto-migrate.

Returns:
  - db: A pointer to the gorm.DB object representing the database connection.
  - err: An error, if any.

“

	func main() {
		db, err := DBConnect("127.0.0.0","5432","postgres","postgres","postgres"")
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		defer db.Close()
	}

“
*/
func DBConnect(host string, port string, username string, password string, dbname string, models ...interface{}) (*gorm.DB, error) {
	// create the DSN string for connecting to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)

	// open a connection to the database using the DSN string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// check if any models were passed in, and if so, auto-migrate them
	if len(models) > 0 {
		err = db.AutoMigrate(models...)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
