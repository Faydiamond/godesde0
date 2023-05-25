package modelos

type Mujer struct {
	Hombre
}

//implementar interfaz
func (M Mujer) Comer() {
	M.comiendo = true
}

func (M Mujer) Respirar() {
	M.respirando = true
}

func (M Mujer) Pensar() {
	M.pensando = true
}

func (M Mujer) Genero() string {
	M.genero = "Mujer"
	return M.genero
}
