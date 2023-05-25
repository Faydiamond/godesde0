package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	genero     string
}

//implementar interfaz
func (H Hombre) Comer() {
	H.comiendo = true
}

func (H Hombre) Respirar() {
	H.respirando = true
}

func (H Hombre) Pensar() {
	H.pensando = true
}

func (H Hombre) Genero() string {
	H.genero = "Hombre"
	return H.genero
}
