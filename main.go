package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	// 1. Tenta carregar o que já estava salvo no arquivo
	carregarDados()

	// Se a lista estiver vazia (primeira vez rodando), podemos colocar uma tarefa padrão
	if len(tarefas) == 0 {
		tarefas = append(tarefas, Tarefa{ID: 1, Titulo: "Minha primeira tarefa persistente", Concluida: false})
		salvarDados() // Já cria o arquivo pela primeira vez
	}

	http.HandleFunc("/tarefas", gerenciarTarefas)

	fmt.Println("Servidor rodando na porta :8080...")
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

		salvarDados()

		// Responde para o usuário com o status 201 (Created) e mostra a tarefa criada
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novaTarefa)

	default:
		// Se tentarem usar PUT, DELETE, etc., nós avisamos que não suportamos ainda
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
func salvarDados() {
		// 1. Transforma o slice 'tarefas' completo em um texto JSON identado (bonito de ler)
		// O segundo e terceiro parâmetros servem para formatar o espaçamento do arquivo
		dadosJSON, err := json.MarshalIndent(tarefas, "", "  ")
		if err != nil {
			fmt.Println("❌ Erro ao converter tarefas para JSON:", err)
			return
		}

		// 2. Grava os bytes no arquivo 'tarefas.json'.
		// O número 0644 é a permissão padrão de leitura e escrita do Linux
		err = os.WriteFile("tarefas.json", dadosJSON, 0644)
		if err != nil {
			fmt.Println("❌ Erro ao salvar o arquivo no disco:", err)
		}
}
func carregarDados() {
	// 1. Tenta ler o arquivo do disco
	dados, err := os.ReadFile("tarefas.json")
	if err != nil {
		// Se o arquivo não existir (primeira vez rodando), apenas ignoramos o erro
		fmt.Println("Nenhum arquivo 'tarefas.json' encontrado. Começando com lista padrão.")
		return
	}

	// 2. Converte o texto JSON lido do arquivo de volta para o nosso slice global '&tarefas'
	err = json.Unmarshal(dados, &tarefas) // Note o '&' porque estamos alterando o slice original!
	if err != nil {
		fmt.Println("❌ Erro ao decodificar o arquivo JSON:", err)
	}
}