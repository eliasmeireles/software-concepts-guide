# Design Pattern: Chain of Responsibility

O **Chain of Responsibility** é um design pattern comportamental que permite passar solicitações por uma cadeia de "
handlers" (manipuladores). Cada "handler" decide se processa a solicitação ou a encaminha para o próximo manipulador na
cadeia. Ele promove a redução do acoplamento entre objetos, delegando responsabilidades para diferentes estágios de
processamento.

---

## **Definição**

1. **Objetivo**:
   Permitir que múltiplos objetos (manipuladores) tenham a chance de processar uma solicitação. Assim, encadeamos os
   manipuladores e passamos a solicitação para cada um deles até que alguém a processe ou a cadeia termine.

2. **Quando Usar**:
    - Quando temos múltiplos handlers que podem processar uma solicitação.
    - Quando é difícil determinar qual manipulador deve lidar com a solicitação.
    - Quando você deseja seguir o princípio da responsabilidade única e evitar blocos grandes de `if-else` ou
      `switch-case`.

---

## **Estrutura**

O padrão é composto pelas seguintes partes principais:

- **`Handler` (Manipulador)**: Interface ou classe abstrata que define um método para processar a solicitação e um ponto
  para passar a solicitação adiante.
- **`ConcreteHandler` (Manipulador Concreto)**: Implementação do `Handler` que processa solicitações específicas ou as
  encaminha para o próximo manipulador.
- **`Client`**: A parte do sistema que inicia a cadeia de solicitações.

---

## **Exemplo em Go**

Aqui está um exemplo implementado em Go para ilustrar o Chain of Responsibility:

```go
package main

import (
	"fmt"
)

// Handler é a interface que define o método HandleRequest e permite o encadeamento de handlers.
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// BaseHandler é uma estrutura abstrata que implementa o "encadeamento" básico.
type BaseHandler struct {
	next Handler
}

// SetNext define o próximo handler na cadeia.
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// HandleNext verifica se existe um próximo handler na cadeia e delega a solicitação.
func (h *BaseHandler) HandleNext(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// ConcreteHandler1 é um handler específico que lida com solicitações "Tipo A".
type ConcreteHandler1 struct {
	BaseHandler
}

// HandleRequest tenta processar uma solicitação do "Tipo A" ou delega para o próximo.
func (h *ConcreteHandler1) HandleRequest(request string) {
	if request == "Tipo A" {
		fmt.Println("ConcreteHandler1 processou a solicitação do Tipo A.")
	} else {
		fmt.Println("ConcreteHandler1 não pôde processar a solicitação. Passando adiante...")
		h.HandleNext(request)
	}
}

// ConcreteHandler2 é um handler específico que lida com solicitações "Tipo B".
type ConcreteHandler2 struct {
	BaseHandler
}

// HandleRequest tenta processar uma solicitação do "Tipo B" ou delega para o próximo.
func (h *ConcreteHandler2) HandleRequest(request string) {
	if request == "Tipo B" {
		fmt.Println("ConcreteHandler2 processou a solicitação do Tipo B.")
	} else {
		fmt.Println("ConcreteHandler2 não pôde processar a solicitação. Passando adiante...")
		h.HandleNext(request)
	}
}

// ConcreteHandler3 é um handler genérico que lida com solicitações desconhecidas.
type ConcreteHandler3 struct {
	BaseHandler
}

// HandleRequest processa qualquer solicitação e encerra a cadeia.
func (h *ConcreteHandler3) HandleRequest(request string) {
	fmt.Println("ConcreteHandler3 processou a solicitação desconhecida:", request)
}

func main() {
	// Criando os handlers
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	// Encadeando os handlers
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Criando diferentes solicitações
	requests := []string{"Tipo A", "Tipo B", "Tipo C"}

	// Processando cada solicitação na cadeia
	for _, req := range requests {
		fmt.Println("Nova solicitação:", req)
		handler1.HandleRequest(req)
		fmt.Println()
	}
}
```

### **Explicação do Código**

1. **Encadeamento de Handlers**:
    - `handler1.SetNext(handler2)` vincula o primeiro handler ao segundo, e assim por diante, formando a cadeia.
2. **Processamento**:
    - Quando a solicitação chega, o primeiro handler tenta processá-la. Se não conseguir, passa ao próximo na cadeia.
3. **Flexibilidade**:
    - Adicionar novos handlers ou reorganizar a cadeia não afeta outros manipuladores.

---

## **Prós e Contras**

### **Prós**

1. **Reduz Acoplamento**:
    - Separa os emissores de solicitações dos objetos que as processam.
2. **Flexibilidade**:
    - A cadeia pode ser estendida, reduzida ou reorganizada facilmente.
3. **Responsabilidade única**:
    - Cada handler é responsável apenas por um tipo específico de solicitação.

### **Contras**

1. **Dificuldade em Depurar**:
    - Como as solicitações passam dinâmicamente, pode ser difícil rastrear onde e como a solicitação foi processada.
2. **Performance**:
    - Em cadeias muito longas, o desempenho pode ser impactado.

---

## **Considerações Finais**

O padrão Chain of Responsibility é uma excelente escolha quando você precisa criar sistemas modulares e flexíveis que
exigem diferentes níveis de processamento para cada tipo de solicitação. No entanto, deve ser usado com cuidado em
cenários de alta complexidade ou quando o rastreamento do fluxo é essencial para o sistema.

---

## Links de Referência

- [Documentação oficial do Go](https://go.dev/doc/)
- [Design Patterns - Refactoring Guru (em inglês)](https://refactoring.guru/design-patterns/chain-of-responsibility)
- [Exemplos de Go Patterns](https://github.com/tmrts/go-patterns)
