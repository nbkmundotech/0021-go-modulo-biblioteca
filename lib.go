package biblioteca

// Livro representa um livro
type Livro struct {
	publicadoEm int
	titulo string
}

// NovoLivro constroi um livro
func NovoLivro(publicadoEm int, titulo string) *Livro {
	return &Livro{
		publicadoEm: publicadoEm,
		titulo: titulo,
	}
}
