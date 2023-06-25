package controllers

import (
	"errors"
	"fmt"
	"log"
	customerrors "people/err"
	"people/model"
	"people/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Struct

type PersonController struct {
	service *services.PersonDb
}

// Function

func NewPersonController(service *services.PersonDb) *PersonController {
	return &PersonController{service: service}
}

// PING

func Ping(c *fiber.Ctx) error {
	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller - Ping - Pong")

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "pong"})
}

// FIND ALL

func (p *PersonController) FindAll(c *fiber.Ctx) error {
	var people []model.Person
	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller Begin - FindAll")

	people, result := p.service.FindAll()

	if result.Error != nil {
		log.Println(requestId + "Controller - FindAll - " + result.Error.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to find all People", "error": result.Error.Error()})
	}

	log.Println(requestId + "Controller End - FindAll - Status : " + fmt.Sprint(fiber.StatusAccepted))

	return c.Status(fiber.StatusAccepted).JSON(people)
}

//FIND BY ID

func (p *PersonController) FindById(c *fiber.Ctx) error {
	var person model.Person

	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller Begin - FindById - Param Id " + c.Params("id"))

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Println(requestId + "Controller - FindById " + "Failed to parse the parameter " + err.Error() + " params: " + c.Params("id"))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to find a person by id", "error": err.Error()})
	}

	person, result := p.service.FindById(id)

	if result.Error != nil {
		log.Println(requestId + "Controller - FindById - Status : Failed to find a person by id " + result.Error.Error())

		var e *customerrors.PersonNotFoundError

		if errors.As(result.Error, &e) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Person with id " + fmt.Sprint(id) + " doesn't exist", "error": result.Error.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to find a person by id", "error": result.Error.Error()})
	}

	log.Println(requestId + "Controller END - FindById - Status : " + fmt.Sprint(fiber.StatusAccepted))

	return c.Status(fiber.StatusAccepted).JSON(person)
}

// CREATE

func (p *PersonController) Create(c *fiber.Ctx) error {
	var person model.Person
	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller Begin - Create ")

	if err := c.BodyParser(&person); err != nil {
		log.Println(requestId + "Controller - Create " + "Failed to parse the json data " + err.Error() + " body: " + string(c.Request().Body()))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse the json data", "error": err.Error})
	}

	if err := p.service.Create(&person).Error; err != nil {
		log.Println(requestId + "Controller - Create - Failed to create a person " + err.Error())
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to create a person", "error": err.Error()})
	}

	log.Println(requestId + "Controller END - Create - Status : " + fmt.Sprint(fiber.StatusAccepted))

	return c.Status(fiber.StatusAccepted).JSON(person)

}

//UPDATE

func (p *PersonController) Update(c *fiber.Ctx) error {
	var person model.Person

	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller Begin - Update ")

	if err := c.BodyParser(&person); err != nil {
		log.Println(requestId + "Controller - Update " + "Failed to parse the json data " + err.Error() + " body: " + string(c.Request().Body()))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse the json data", "error": err.Error})
	}

	person, result := p.service.Update(person)

	if result.Error != nil {
		log.Println(requestId + "Controller - Update - Failed to update a person : " + result.Error.Error())

		var e *customerrors.PersonNotFoundError

		if errors.As(result.Error, &e) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to Update a person", "error": result.Error.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update a person ", "error": result.Error.Error()})
	}

	log.Println(requestId + "Controller - Update - Status : " + fmt.Sprint(fiber.StatusAccepted))

	return c.Status(fiber.StatusAccepted).JSON(person)
}

//Delete

func (p *PersonController) Delete(c *fiber.Ctx) error {
	var person model.Person

	requestId := " - " + fmt.Sprintf("%+v", c.Locals("requestid")) + " - "

	log.Println(requestId + "Controller - Delete - Param Id " + c.Params("id"))

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Println(requestId + "Controller - Delete " + "Failed to parse the parameter " + err.Error() + " params: " + c.Params("id"))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete a person", "error": err.Error()})
	}

	person.Id = id
	person, result := p.service.Delete(person)

	if result.Error != nil {
		log.Println(requestId + "Controller - Delete - Status : Failed to Delete a person " + result.Error.Error())

		var e *customerrors.PersonNotFoundError

		if errors.As(result.Error, &e) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to Delete a person", "error": result.Error.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to Delete a person", "error": result.Error.Error()})
	}

	log.Println(requestId + "Controller - Delete - Status : " + fmt.Sprint(fiber.StatusAccepted))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deleted"})
}
