package biblioteca

type Livro struct {
	publicadoEm int
	titulo string
}

func NovoLivro(publicadoEm int, titulo string) *Livro {
	return &Livro{
		publicadoEm: publicadoEm,
		titulo: titulo,
	}
}
