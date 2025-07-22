package student

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"studentCRUD/database"
)

func CreateStudent() {
	var name string
	var surname string
	var grade float64

	fmt.Print("Insira o nome do aluno: ")
	fmt.Scan(&name)

	fmt.Print("Insira o sobrenome do aluno: ")
	fmt.Scan(&surname)

	fmt.Print("Insira a nota do aluno: ")
	fmt.Scan(&grade)

	_, err := database.DB.Exec("INSERT INTO students (name, surname, grade) VALUES ($1, $2, $3)", name, surname, grade)
	if err != nil {
		fmt.Println("Erro ao cadastrar Aluno:", err)
		return
	}
	fmt.Println("Aluno cadastrado com sucesso!")
}

func ListStudents() {
	rows, err := database.DB.Query("SELECT id, name, surname, grade FROM students")
	if err != nil {
		fmt.Println("Erro ao listar alunos", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var surname string
		var grade float64

		err := rows.Scan(&id, &name, &surname, &grade)
		if err != nil {
			fmt.Println("Erro ao ler usuários", err)
			continue
		}
		fmt.Printf("\nID: %d | Nome: %s | Sobrenome: %s | Nota: %.1f", id, name, surname, grade)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Erro ao iterar pelos alunos", err)
	}
}

func UpdateStudent() {
	ListStudents()

	var id int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nSelecione o ID do aluno que deseja atualizar/editar: ")
	fmt.Scan(&id)

	reader.ReadString('\n')

	var atualNome, atualSobrenome string
	var atualNota float64

	err := database.DB.QueryRow("SELECT name, surname, grade FROM students WHERE id = $1", id).
		Scan(&atualNome, &atualSobrenome, &atualNota)
	if err != nil {
		fmt.Println("Aluno não encontrado ou erro:", err)
		return
	}

	novoNome := atualNome
	novoSobrenome := atualSobrenome
	novaNota := atualNota

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Atualizar/Editar Nome")
		fmt.Println("2. Atualizar/Editar Sobrenome")
		fmt.Println("3. Atualizar/Editar Nota")
		fmt.Println("4. Salvar e retornar")
		fmt.Print("Escolha uma opção: ")

		inputOpcao, _ := reader.ReadString('\n')
		inputOpcao = strings.TrimSpace(inputOpcao)
		opcao, err := strconv.Atoi(inputOpcao)
		if err != nil {
			fmt.Println("Opção inválida, digite um número válido.")
			continue
		}

		switch opcao {
		case 1:
			fmt.Printf("Digite o novo nome (ENTER para manter '%s'): ", novoNome)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input != "" {
				novoNome = input
			}

		case 2:
			fmt.Printf("Digite o novo sobrenome (ENTER para manter '%s'): ", novoSobrenome)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input != "" {
				novoSobrenome = input
			}

		case 3:
			fmt.Printf("Digite a nova nota (ENTER para manter '%.2f'): ", novaNota)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input != "" {
				var n float64
				_, err := fmt.Sscanf(input, "%f", &n)
				if err == nil {
					novaNota = n
				} else {
					fmt.Println("Nota inválida, mantendo valor anterior.")
				}
			}

		case 4:
			result, err := database.DB.Exec(
				"UPDATE students SET name=$1, surname=$2, grade=$3 WHERE id=$4",
				novoNome, novoSobrenome, novaNota, id)
			if err != nil {
				fmt.Println("Erro ao atualizar aluno:", err)
				return
			}
			linhasAfetadas, _ := result.RowsAffected()
			if linhasAfetadas == 0 {
				fmt.Println("Aluno não encontrado.")
			} else {
				fmt.Println("Aluno atualizado com sucesso!")
			}
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}
func DeleteStudent() {
	ListStudents()

	var id int
	fmt.Print("\nSelecione o ID do aluno que deseja excluir: ")
	fmt.Scan(&id)

	result, err := database.DB.Exec("DELETE FROM students WHERE id=$1", id)
	if err != nil {
		fmt.Println("Erro ao excluir aluno:", err)
		return
	}

	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		fmt.Println("Aluno não encontrado")
		return
	}
	fmt.Println("Aluno deletado com sucesso!")
}
