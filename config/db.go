package config

import (
	"fmt"

	"github.com/suhailmshaik/pg-crn-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname=pg-crn-test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	// Check if the sequence exists before creating it
	var sequenceExists bool
	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_sequences WHERE schemaname = 'public' AND sequencename = 'crn_seq')").Scan(&sequenceExists)

	if !sequenceExists {
		// Creating the CRN sequence if it doesn't exist
		// 8 quintillion, 223 quadrillion, 372 trillion, 36 billion, 854 million, 775 thousand, 808 - possible number of combinations between 1000000000000000000 and 9223372036854775807
		crnSequence := `CREATE SEQUENCE IF NOT EXISTS crn_seq
						START WITH 1000000000000000000
						INCREMENT BY 1
						NO MINVALUE
						MAXVALUE 9223372036854775807
						CACHE 1;`

		if err := db.Exec(crnSequence).Error; err != nil {
			fmt.Println("Error creating sequence: ", err)
			panic(err)
		}
		fmt.Println("CRN Sequence created")
	} else {
		fmt.Println("CRN Sequence already exists")
	}

	// Automatically migrate the schema, which will only create the table if it doesn't exist
	if err := db.AutoMigrate(&models.Payouts{}); err != nil {
		fmt.Println("Error during migration:", err)
		panic(err)
	}

	fmt.Println("Database connected and migration completed")

	DB = db
}

