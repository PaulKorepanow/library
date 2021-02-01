package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/postgres"
	"log"
	"main/internal/model"
	"main/internal/store"
	"main/internal/store/sqlbase"
	"net/http"
	"os"
	"path"
)

type Config struct {
	BindAddr    string `json:"bind_addr"`
	LogLevel    string `json:"log_level"`
	DataBaseUrl string `json:"data_base_url"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config *Config
	file, err := os.Open(path.Join(configPath, "config.json"))
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

type Server struct {
	config *Config
	logger *log.Logger
	db     store.Store
}

func NewServer(config *Config, db *sqlbase.SqlStore) Server {
	logger := log.New(os.Stdout, "/n/r", log.LstdFlags)
	return Server{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Start() {
	s.logger.Println("Starting server ...")
	http.ListenAndServe(s.config.BindAddr, s.configureRouter())
}

func (s *Server) configureRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", s.HandleMain()).Methods(http.MethodGet)
	return router
}

func ConnectToDB(databaseURL string) (*sqlbase.SqlStore, error) {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&model.Book{}).Error; err != nil {
		return nil, err
	}
	return sqlbase.NewSqlStore(db), nil
}
