package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

func main() {
	password := "admin1324"
	
	// Generate bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	
	fmt.Printf("Generated hash for password '%s':\n%s\n\n", password, string(hash))
	
	// Open database
	db, err := sql.Open("sqlite", "./data/database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	
	// Update password for testcli user
	result, err := db.Exec("UPDATE users SET password_hash = ? WHERE username = 'testcli'", string(hash))
	if err != nil {
		log.Fatalf("Failed to update password: %v", err)
	}
	
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("✅ Updated %d row(s) for user 'testcli'\n", rowsAffected)
	fmt.Println("You can now login with: testcli / admin1324")
}
