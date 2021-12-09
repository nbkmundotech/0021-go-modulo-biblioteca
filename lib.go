package biblioteca

// Livro representa um livro
type Livro struct {
	publicadoEm int
	titulo string
}

// NovoLivro constroi um livro
func NovoLivro(titulo string, publicadoEm int) *Livro {
	return &Livro{
		publicadoEm: publicadoEm,
		titulo: titulo,
	}
}
