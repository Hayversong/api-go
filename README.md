# Todo API em Go

Uma API REST simples e extremamente veloz para gerenciamento de listas de tarefas (Todo List), construída em Go utilizando apenas pacotes nativos da linguagem.

Este projeto foi desenvolvido como objeto de estudo para consolidar conceitos de desenvolvimento backend, manipulação de estruturas de dados dinâmicas na memória e protocolo HTTP.

## Tecnologias e Ferramentas Utilizadas

* **Linguagem:** Go (Golang)
* **Editor/IDE:** Zed IDE
* **Ambiente:** WSL2 (Ubuntu Linux)
* **Controle de Versão:** Git & GitHub (Autenticação via SSH)

## Conceitos Praticados

* **Servidor HTTP Nativo:** Utilização do pacote nativo `net/http` para subir o servidor e rotear requisições.
* **Manipulação de JSON:** Codificação e decodificação de dados utilizando `encoding/json`.
* **Estruturas de Dados:** Uso de `structs` com tags JSON e gerenciamento de memória dinâmica com `slices`.
* **Controle de Fluxo HTTP:** Tratamento de verbos HTTP (`GET` e `POST`) na mesma rota através de estruturas condicionais (`switch/case`).

## Rotas da API

A API responde no endereço padrão `http://localhost:8080` e possui os seguintes endpoints:

| Método | Rota | Descrição | Status de Resposta |
| :--- | :--- | :--- | :--- |
| **GET** | `/tarefas` | Retorna a lista completa de tarefas em formato JSON. | `200 OK` |
| **POST** | `/tarefas` | Cadastra uma nova tarefa na lista. | `201 Created` |

### Exemplo de corpo para o POST:
```json
{
  "titulo": "Finalizar o desafio de persistência",
  "concluida": false
}
