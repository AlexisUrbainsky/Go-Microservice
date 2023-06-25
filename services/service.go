package services

import (
	"log"
	customerrors "people/err"
	"people/model"

	"gorm.io/gorm"
)

// Struct
type PersonDb struct {
	db *gorm.DB
}

// Function
func NewPersonDb(db *gorm.DB) *PersonDb {
	return &PersonDb{db: db}
}

//Methods

func (p *PersonDb) FindAll() ([]model.Person, *gorm.DB) {
	var people []model.Person

	result := p.db.Find(&people)

	if result.Error != nil {
		log.Println("Error to get all people " + result.Error.Error())
	}

	return people, result
}

func (p *PersonDb) FindById(id int) (model.Person, *gorm.DB) {
	var person model.Person

	result := p.db.Find(&person, id)

	if result.Error != nil {
		log.Println("Error to get person " + result.Error.Error())
	} else if person.Id == 0 {
		log.Println(customerrors.PersonNotFoundErr(id).Error())
		result.Error = customerrors.PersonNotFoundErr(id)
	}

	return person, result
}

func (p *PersonDb) Create(person *model.Person) *gorm.DB {

	result := p.db.Debug().Create(&person)

	if result.Error != nil {
		log.Println("Error to create a new person " + result.Error.Error())
	}

	return result
}

func (p *PersonDb) Update(person model.Person) (model.Person, *gorm.DB) {

	result := p.db.Debug().Save(&person)

	if result.Error != nil {
		log.Println("Error to update person " + result.Error.Error())
	} else if person.Id == 0 || result.RowsAffected == 0 {
		log.Println(customerrors.PersonNotFoundErr(person.Id).Error())
		result.Error = customerrors.PersonNotFoundErr(person.Id)
	}

	return person, result
}

func (p *PersonDb) Delete(person model.Person) (model.Person, *gorm.DB) {

	result := p.db.Debug().Delete(&person)

	if result.Error != nil {
		log.Println("Error to delete person " + result.Error.Error())
	} else if person.Id == 0 || result.RowsAffected == 0 {
		log.Println(customerrors.PersonNotFoundErr(person.Id).Error())
		result.Error = customerrors.PersonNotFoundErr(person.Id)
	}

	return person, result
}
