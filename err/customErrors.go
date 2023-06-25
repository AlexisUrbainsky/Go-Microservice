package customerrors

import "fmt"

type PersonNotFoundError struct {
	Id int
}

func (p *PersonNotFoundError) Error() string {
	return fmt.Sprintf("The person with Id %v doesn't exist in the database ", p.Id)
}

func PersonNotFoundErr(id int) error {
	return &PersonNotFoundError{
		Id: id,
	}

}
