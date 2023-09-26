package modelos

type Hombre struct {
	Edad       int
	Nombre     string
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
	Vivo       bool
}

func (h *Hombre) Respirar() { h.Respirando = true }

// solo contiene las mismas funciones de la interfaz hombre se convierte en humano
func (h *Hombre) Comer()       { h.Comiendo = true }
func (h *Hombre) Pensar()      { h.Pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }
