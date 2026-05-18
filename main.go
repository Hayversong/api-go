package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Tarefa representa a estrutura do dado que vamos trafegar na web.
// Os blocos `json:"id"` avisam ao Go como esse campo deve ser escrito em formato JSON.
type Tarefa struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Concluida bool   `json:"concluida"`
}

// Criamos um Slice global para simular o nosso "banco de dados" na memória
var tarefas []Tarefa

func main() {
	// Alimentando nossa lista com algumas tarefas iniciais para teste
	tarefas = append(tarefas, Tarefa{ID: 1, Titulo: "Estudar ponteiros em Go", Concluida: true})
	tarefas = append(tarefas, Tarefa{ID: 2, Titulo: "Criar minha primeira API", Concluida: false})

	// Vinculamos a rota "/tarefas" a uma função que vai processar a requisição
	http.HandleFunc("/tarefas", gerenciarTarefas)

	fmt.Println("Servidor rodando na porta :8080...")
	
	// Liga o servidor na porta 8080. O 'nil' significa que usaremos o roteador padrão do Go.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

// listarTarefas é o nosso "Handler" (manipulador). 
// Toda função que lida com web em Go precisa receber esses dois parâmetros:
// 1. w: Onde se escreve a RESPOSTA para o usuário.
// 2. r: Onde se lê a REQUISIÇÃO que veio da internet (dados enviados, cabeçalhos, etc).
func gerenciarTarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet: // Equivalente a "GET"
		// Código que você já fez: apenas devolve a lista atual
		json.NewEncoder(w).Encode(tarefas)

	case http.MethodPost: // Equivalente a "POST"
		var novaTarefa Tarefa

		// json.NewDecoder(r.Body).Decode(&novaTarefa) faz o inverso:
		// Pega o texto JSON que veio no "corpo" da requisição da internet,
		// converte de volta para uma Struct do Go e joga no endereço de '&novaTarefa'
		err := json.NewDecoder(r.Body).Decode(&novaTarefa)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		// Regra de negócio simples: gera um ID baseado no tamanho da lista + 1
		novaTarefa.ID = len(tarefas) + 1

		// Adiciona a nova tarefa no nosso Slice global
		tarefas = append(tarefas, novaTarefa)

		// Responde para o usuário com o status 201 (Created) e mostra a tarefa criada
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novaTarefa)

	default:
		// Se tentarem usar PUT, DELETE, etc., nós avisamos que não suportamos ainda
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}