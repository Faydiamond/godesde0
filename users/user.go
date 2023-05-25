package users

import (
	"fmt"
	"time"

	"github.com/Faydiamond/godesde0/modelos"
)

func AgregarUsuario() {
	u := new(modelos.User)

	u.AddUser(1, "Fabian", time.Now(), true)

	fmt.Println(u)
}
