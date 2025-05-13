CREATE TABLE livro (
    idLivro SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    autor VARCHAR(100) NOT NULL,
    anoPublicacao INT NOT NULL
);

CREATE TABLE emprestimo(
	idEmprestimo SERIAL PRIMARY KEY,
	nomeLeitor VARCHAR(100)  NOT NULL,
	dataEmprestimo DATE NOT NULL,
	dataDevolucao DATE,	
	observacao TEXT,
	idLivro INT NOT NULL,
	FOREIGN KEY (idLivro) REFERENCES livro(idLivro)
);

ALTER TABLE livro ADD COLUMN disponibilidade BOOLEAN NOT NULL;
ALTER TABLE livro ALTER COLUMN autor TYPE VARCHAR (200);

ALTER TABLE emprestimo DROP COLUMN observacao;


INSERT INTO livro(titulo, autor, anoPublicacao, disponibilidade)
VALUES
  ('Dom Casmurro', 'Machado de Assis', 1899, true),
  ('1984', 'George Orwell', 1949, false),
  ('A Revolução dos Bichos', 'George Orwell', 1945, true),
  ('Capitães da Areia', 'Jorge Amado', 1937, true);

INSERT INTO emprestimo (idLivro, nomeLeitor, dataEmprestimo, dataDevolucao)
VALUES 
	(1, 'Maria Silva', '2025-05-01', '2025-05-15'),
	(2, 'João Santos', '2025-05-03', '2025-05-17'),
	(3, 'Ana Souza', '2025-05-05', '2025-05-19'),
	(4, 'Carlos Oliveira', '2025-05-07', '2025-05-21');

UPDATE emprestimo SET nomeLeitor = 'Adriam' WHERE idEmprestimo = 1;
DELETE FROM emprestimo WHERE idEmprestimo = 4;

SELECT * FROM livro WHERE disponibilidade = true ORDER BY titulo ASC;

SELECT * FROM livro;
SELECT * FROM emprestimo;