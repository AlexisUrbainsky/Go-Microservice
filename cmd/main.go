package main

import (
	"crypto"
	"crypto/tls"
	"log"
	"net"
	"os"
	"people/config"
	"people/controllers"
	"people/router"
	"people/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/crypto/pkcs12"
)

func init() {
	config.LoadVariables()
	config.LoadLogsConfiguration()
}

func main() {

	app := initFiberApp()

	if os.Getenv("ENV") == "PROD" {
		listener := initConfigHTTPS()
		log.Fatal(app.Listener(listener))
	} else {
		log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	}
}

func initFiberApp() *fiber.App {

	// Fiber Instance
	app := fiber.New()

	// RequestID
	app.Use(requestid.New())

	// Load DB Configuration
	db, errDb := config.ConnectToDB()

	if errDb != nil {
		log.Println("Error to get the DB " + errDb.Error())
	}

	// Load Service
	personService := services.NewPersonDb(db)

	// Load Controller
	personController := controllers.NewPersonController(personService)

	// Load Routes
	router.AddRoutes(app, personController)

	return app
}

func initTLSConfig(path string, password string) (*tls.Certificate, error) {

	pkcs12Data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	key, cert, err := pkcs12.Decode(pkcs12Data, password)

	if err != nil {
		return nil, err
	}

	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key.(crypto.PrivateKey),
		Leaf:        cert,
	}

	return &tlsCert, nil
}

func initConfigHTTPS() net.Listener {

	path := os.Getenv("CERT_PATH")
	password := os.Getenv("CERT_PASSWORD")

	tlsCert, error := initTLSConfig(path, password)

	if error != nil {
		log.Println("Unable to initialize TLS configuration object. Check your configuration and try again. Program will STOP.")
	}

	config := &tls.Config{Certificates: []tls.Certificate{*tlsCert}}

	ln, err := tls.Listen("tcp", ":6000", config)

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return ln
}
