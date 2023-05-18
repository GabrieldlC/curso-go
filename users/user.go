package users

import (
	"fmt"
	"time"

	"github.com/curso-go/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Pedro", time.Now(), true)

	fmt.Println(u)
}
