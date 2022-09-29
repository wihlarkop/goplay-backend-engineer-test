package container

import (
	"fmt"
	"log"
	"os"

	"goplay-backend-engineer-test/repositories/file"
	"goplay-backend-engineer-test/repositories/user"
	"goplay-backend-engineer-test/usecase/file/getfile"
	"goplay-backend-engineer-test/usecase/file/getlistfile"
	"goplay-backend-engineer-test/usecase/file/uploadfile"
	"goplay-backend-engineer-test/usecase/user/userlogin"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Container struct {
	GetFileUsecase     getfile.Inport
	GetListFileUsecase getlistfile.Inport
	UploadFileUsecase  uploadfile.Inport

	UserLoginUsecase userlogin.Inport
}

func NewContainer() *Container {
	sqliteClient, err := Init()
	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	fileRepo := file.NewRepo(sqliteClient)
	userRepo := user.NewRepo(sqliteClient)

	return &Container{
		GetFileUsecase:     getfile.NewUsecase(fileRepo),
		GetListFileUsecase: getlistfile.NewUsecase(fileRepo),
		UploadFileUsecase:  uploadfile.NewUsecase(fileRepo),

		UserLoginUsecase: userlogin.NewUsecase(userRepo),
	}
}

func Init() (*sqlx.DB, error) {
	conn, err := NewClientSqlite()
	return conn, err
}

func NewClientSqlite() (*sqlx.DB, error) {
	cwd, _ := os.Getwd()
	dbName := os.Getenv("SQLITE_DBNAME")
	dbFilePath := cwd + "/storage/" + dbName

	// Check Path File
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		fmt.Println("Create new database " + dbName)

		file, err := os.Create(dbFilePath) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		fmt.Println("Database " + dbName + " created")
	}

	SQLiteDB, err := sqlx.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	createTables(SQLiteDB)

	//defer SQLiteDB.Close() // Defer Closing the database

	return SQLiteDB, nil
}

func createTables(db *sqlx.DB) {
	// Create Table Query
	query := `
		CREATE TABLE IF NOT EXISTS "users" (
			"id"	INTEGER,
			"username"	TEXT NOT NULL UNIQUE,
			"password"	TEXT NOT NULL UNIQUE,
			"created_at" Timestamp With Time Zone NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
		
		CREATE TABLE IF NOT EXISTS "upload_file" (
			"id"	INTEGER,
			"path"	TEXT NOT NULL,
			"upload_by" TEXT NOT NULL,
			"created_at" Timestamp With Time Zone NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}
