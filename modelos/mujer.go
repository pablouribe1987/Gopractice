package modelos

type Mujer struct {
	Hombre
}

//func (m *Mujer) Respirar() { m.respirando = true }

// solo contiene las mismas funciones de la interfaz hombre se convierte en humano
// func (m *Mujer) Comer()       { m.comiendo = true }
// func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }
