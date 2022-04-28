package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // Для того чтобы отработала функция init
	"log"
)

//Instance of storage
type Storage struct {
	config *Config
	//DataBase FileDescriptor
	db *sql.DB
	//Subfield for repoInterfacing (modelUser)
	UserRepository *UserRepository
	//Subfield for repoInterfacing (modelArticle)
	ArticleRepository *ArticleRepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	storage.db = db
	log.Println("Database connection created successfully")
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

//Public Repo for User
func (s *Storage) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		storage: s,
	}
	return nil
}

//Public Repo for Article
func (s *Storage) Article() *ArticleRepository {
	if s.ArticleRepository != nil {
		return s.ArticleRepository
	}
	s.ArticleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}
