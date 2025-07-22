package main

import (
	"fmt"
	"studentCRUD/database"
	"studentCRUD/student"
)

func main() {
	database.Conectar()
	defer database.DB.Close()

	for {
		var opcao int
		fmt.Println("\n---- Menu do Aluno ----")
		fmt.Print("\n1. Cadastrar Aluno")
		fmt.Print("\n2. Listar alunos existentes")
		fmt.Print("\n3. Atualizar/Editar Alunos")
		fmt.Print("\n4. Deletar Aluno")
		fmt.Print("\n0. Sair")
		fmt.Print("\nEscolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			student.CreateStudent()
		case 2:
			student.ListStudents()
		case 3:
			student.UpdateStudent()
		case 4:
			student.DeleteStudent()
		case 0:
			fmt.Println("Encerrando programa...")
			return
		}
	}
}
