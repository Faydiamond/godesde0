package modelos

import "time"

type (
	User struct {
		Id        int
		Nombre    string
		CreatedAt time.Time
		Status    bool
	}
)

func (U *User) AddUser(id int, nombre string, created time.Time, status bool) {//siempre puntero
	U.Id = id
	U.Nombre = nombre
	U.CreatedAt = created
	U.Status = status
}
