package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Server struct {
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

// Setup initialize the configuration instance
func Setup() {
	b, err := readFile("configs/app.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, ServerSetting)
	if err != nil {
		log.Fatal(err)
	}

	b, err = readFile("configs/database.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, DatabaseSetting)
	if err != nil {
		log.Fatal(err)
	}

	b, err = readFile("configs/redis.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, RedisSetting)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open '%s' - %s", filePath, err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read '%s' - %s", filePath, err)
	}

	return byteValue, nil
}
