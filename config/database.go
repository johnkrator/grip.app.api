package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Read database connection parameters from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Backup the database before proceeding
	if err := backupDatabase(host, port, user, password, dbname); err != nil {
		return nil, fmt.Errorf("failed to backup database: %v", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Define the models you want to keep
	modelsToKeep := GetModelsToKeep()

	// Auto Migrate
	err = db.AutoMigrate(modelsToKeep...)
	if err != nil {
		return nil, fmt.Errorf("failed to run auto migrations: %v", err)
	}

	// Check for columns to remove in each table
	for _, model := range modelsToKeep {
		err := removeUnusedColumns(db, model)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func backupDatabase(host, port, user, password, dbname string) error {
	// Create a timestamp for the backup file name
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	// Create the full path for the backups directory
	backupDir := filepath.Join(currentDir, "backups")

	// Ensure the backups directory exists
	err = os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Create the full path for the backup file
	backupPath := filepath.Join(backupDir, fmt.Sprintf("backup_%s.sql", timestamp))

	// Get the absolute path of the backup file
	absBackupPath, err := filepath.Abs(backupPath)
	if err != nil {
		fmt.Printf("Warning: Couldn't get absolute path: %v\n", err)
	} else {
		fmt.Printf("Backup will be saved to: %s\n", absBackupPath)
	}

	// Construct the pg_dump command
	cmd := exec.Command("C:\\Program Files\\PostgreSQL\\15\\bin\\pg_dump",
		"-h", host,
		"-p", port,
		"-U", user,
		"-d", dbname,
		"-f", backupPath,
		"-Z", "9", // Add compression (0-9, 9 is best compression)
		"-Fc") // Use custom format for more flexible restores

	// Set the PGPASSWORD environment variable
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create backup: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Database backup created successfully: %s\n", backupPath)
	return nil
}

// removeUnusedColumns removes any columns that are not used in the model
func removeUnusedColumns(db *gorm.DB, model interface{}) error {
	tableName := db.NamingStrategy.TableName(reflect.TypeOf(model).Elem().Name())

	// Get existing columns
	columns, err := db.Migrator().ColumnTypes(model)
	if err != nil {
		return fmt.Errorf("failed to get column types for %s: %v", tableName, err)
	}

	// Get model fields
	modelType := reflect.TypeOf(model).Elem()
	modelFields := make(map[string]bool)
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		modelFields[db.NamingStrategy.ColumnName("", field.Name)] = true
	}

	// Add GORM's default fields
	defaultFields := []string{"id", "created_at", "updated_at", "deleted_at"}
	for _, field := range defaultFields {
		modelFields[field] = true
	}

	// Check for columns to remove
	for _, column := range columns {
		columnName := column.Name()
		if _, exists := modelFields[columnName]; !exists {
			// Skip dropping protected columns
			if isProtectedColumn(columnName) {
				fmt.Printf("Skipping protected column %s in table %s\n", columnName, tableName)
				continue
			}
			if err := db.Migrator().DropColumn(model, columnName); err != nil {
				return fmt.Errorf("failed to drop column %s from %s: %v", columnName, tableName, err)
			}
			fmt.Printf("Dropped column %s from table %s\n", columnName, tableName)
		}
	}

	return nil
}

// isProtectedColumn checks if a column is a protected column
func isProtectedColumn(columnName string) bool {
	protectedColumns := map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}
	return protectedColumns[columnName]
}

/*
1. These changes do the following:

We've added GORM's default fields (id, created_at, updated_at, deleted_at) to the modelFields map to prevent them from being considered as "unused" columns.
We've introduced a new function isProtectedColumn that checks if a column name is in the list of protected columns that should never be dropped.
Before attempting to drop a column, we now check if it's a protected column. If it is, we skip it and log a message.

This approach should prevent the error you encountered by not attempting to drop essential columns like id. It will only
drop columns that are truly no longer present in your model and are not part of GORM's default fields.

Remember to always back up your database before running migrations, especially when dropping columns is involved.
This ensures you can recover your data if something unexpected occurs during the migration process.

2. Key changes and additions:

In the backupDatabase function:
We now use filepath.Join to create paths, which is more robust across different operating systems.
We get the current working directory and create an absolute path for the backup file.
We log the full absolute path of where the backup will be saved.
We've added the -Z 9 flag for best compression and -Fc for custom format backup.
Error handling has been improved throughout the code.
The code now uses os.MkdirAll to create the backup directory if it doesn't exist.
We're now using absolute paths, which should help avoid any confusion about where files are being saved.

This updated code incorporates better path handling, improved logging, and enhanced backup options.
It should be more robust and provide clearer information about what's happening during the backup process.

Remember to test this thoroughly in your development environment before using it in production.
Also, consider implementing a backup retention policy and possibly moving old backups to secure, off-site storage for added data protection.
*/
