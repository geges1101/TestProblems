package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type Analytics struct {
	ID        uint      `gorm:"primary_key"`
	Time      time.Time `json:"time"`
	UserID    string    `json:"user_id"`
	Data      Data      `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Data struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

type Response struct {
	Status string `json:"status"`
}

func main() {
	// Чтение конфигурации из файла или переменных окружения
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	// Установка уровня логирования
	logLevel := viper.GetString("log.level")
	switch logLevel {
	case "debug":
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	case "info":
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	case "error":
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
		log.SetOutput(os.Stderr)
	default:
		log.Fatalf("Invalid log level: %s", logLevel)
	}

	// Конфигурация подключения к базе данных PostgreSQL
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.dbname"),
		viper.GetString("db.password"),
	)

	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы аналитики, если она не существует
	if !db.HasTable(&Analytics{}) {
		db.AutoMigrate(&Analytics{})
	}

	// Инициализация роутера
	router := gin.Default()

	// Регистрация маршрутов pprof для профилирования и мониторинга
	pprof.Register(router)

	// Обработчик запроса на сохранение аналитических данных
	router.POST("/analytics", func(c *gin.Context) {
		var jsonBody map[string]interface{}
		err = c.BindJSON(&jsonBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		headers := make(map[string]string)
		for key, values := range c.Request.Header {
			headers[key] = values[0]
		}

		userID := c.Request.Header.Get("X-Tantum-Authorization")

		data := Data{
			Headers: headers,
			Body:    jsonBody,
		}

		analytics := Analytics{
			Time:   time.Now(),
			UserID: userID,
			Data:   data,
		}

		// Сохранение аналитических данных в базе данных (асинхронно)
		go func() {
			if err = db.Create(&analytics).Error; err != nil {
				log.Printf("Failed to save data: %v", err)
			}
		}()

		c.JSON(http.StatusAccepted, Response{Status: "OK"})
	})

	// Запуск сервера на указанном порту
	port := viper.GetInt("server.port")
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", addr)
	err = router.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
