package config

import (
	"errors"
	"log"
	"os"

	"gorm.io/gorm/logger"
)

func loadLogsConfigFile() *os.File {

	if _, err := os.Stat(os.Getenv("LOGPATH")); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(os.Getenv("LOGPATH"), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	file, err := os.OpenFile(os.Getenv("LOGPATH")+os.Getenv("LOGFILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return file
}

func LoadLogsConfiguration() {
	log.SetOutput(loadLogsConfigFile())
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

}

func CreateLogsConfigGorm() logger.Interface {
	gormLogger := logger.New(log.New(loadLogsConfigFile(), "", log.Ldate|log.Ltime|log.Lmicroseconds), // io writer
		logger.Config{
			LogLevel:                  logger.Error | logger.Info, // Log level
			IgnoreRecordNotFoundError: false,                      // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                      // Don't include params in the SQL log
			Colorful:                  false,                      // Disable color
		},
	)
	return gormLogger
}
