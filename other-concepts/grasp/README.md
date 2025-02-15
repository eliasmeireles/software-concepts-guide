# GRASP: General Responsibility Assignment Software Patterns

## O que é GRASP?

GRASP (General Responsibility Assignment Software Patterns) é um conjunto de padrões e diretrizes para design de
software orientado a objetos. Ele foi introduzido por **Craig Larman** em seu livro _Applying UML and Patterns_ como uma
maneira de ajudar os desenvolvedores a atribuir responsabilidades de forma lógica e eficiente em sistemas de software. O
principal objetivo do GRASP é melhorar a coesão e reduzir o acoplamento, promovendo um design robusto e de fácil
manutenção.

GRASP não é um padrão de implementação, mas sim um guia conceitual para decisão de design. Ele inclui nove princípios:

1. Creator
2. Information Expert
3. Low Coupling
4. High Cohesion
5. Controller
6. Polymorphism
7. Pure Fabrication
8. Indirection
9. Protected Variations

### Princípios em Destaque:

**Creator:** Define quem é responsável por criar um objeto. A ideia é atribuir a responsabilidade de criação de um
objeto a uma classe que mais se relaciona com ele.

**Information Expert:** Foca na atribuição de responsabilidades às classes que possuem as informações necessárias para
realizar a tarefa.

**Low Coupling e High Cohesion:** Trabalham juntos para manter o design modular e reutilizável, reduzindo dependências
desnecessárias entre as classes enquanto garante que uma classe tenha uma única função bem definida.

**Controller:** Especifica quem deve ser responsável por gerenciar eventos e funcionalidades no sistema.

---

## Implementação: Exemplo em Go

Abaixo temos um exemplo simples implementado em **Go** para ilustrar os princípios de GRASP. Vamos simular um sistema de
pedidos.

### Exemplo: Controlador de Pedidos (Principles: Controller, Creator, High Cohesion e Low Coupling)

```go
package main

import "fmt"

// Pedido representa um pedido feito por um cliente.
type Pedido struct {
	ID       int
	Produtos []string
	Total    float64
	Status   string
}

// PedidoController (Controller) gerencia ações do sistema relacionadas a pedidos.
type PedidoController struct {
	// Dependência de baixa acoplagem com o repositório
	repo *PedidoRepository
}

// NovoPedidoController cria uma instância de PedidoController.
func NovoPedidoController(repo *PedidoRepository) *PedidoController {
	return &PedidoController{repo: repo}
}

// CriarPedido adiciona um novo pedido ao sistema.
func (pc *PedidoController) CriarPedido(produtos []string, total float64) *Pedido {
	novoPedido := &Pedido{
		ID:       pc.repo.GetProximoID(), // Information Expert (repo é responsável pela numeração)
		Produtos: produtos,
		Total:    total,
		Status:   "Criado",
	}

	pc.repo.SalvarPedido(novoPedido) // Controller delega a ação à camada repositório
	return novoPedido
}

// AtualizarStatus altera o status de um pedido.
func (pc *PedidoController) AtualizarStatus(id int, status string) error {
	pedido := pc.repo.BuscarPedidoPorID(id)
	if pedido == nil {
		return fmt.Errorf("Pedido com ID %d não encontrado", id)
	}
	pedido.Status = status
	pc.repo.SalvarPedido(pedido)
	return nil
}

// PedidoRepository é responsável por armazenar e recuperar pedidos.
type PedidoRepository struct {
	proximoID int
	pedidos   map[int]*Pedido
}

// NovoPedidoRepository cria uma instância de PedidoRepository.
func NovoPedidoRepository() *PedidoRepository {
	return &PedidoRepository{
		proximoID: 1,
		pedidos:   make(map[int]*Pedido),
	}
}

// GetProximoID retorna o próximo ID único para um pedido.
func (repo *PedidoRepository) GetProximoID() int {
	id := repo.proximoID
	repo.proximoID++
	return id
}

// SalvarPedido armazena ou atualiza o pedido no repositório.
func (repo *PedidoRepository) SalvarPedido(pedido *Pedido) {
	repo.pedidos[pedido.ID] = pedido
}

// BuscarPedidoPorID retorna um pedido pelo ID.
func (repo *PedidoRepository) BuscarPedidoPorID(id int) *Pedido {
	return repo.pedidos[id]
}

func main() {
	// Criando repositório e controlador
	repo := NovoPedidoRepository()
	controller := NovoPedidoController(repo)

	// Criando um pedido
	pedido := controller.CriarPedido([]string{"Produto A", "Produto B"}, 129.90)
	fmt.Printf("Pedido criado: %+v\n", pedido)

	// Atualizando status do pedido
	err := controller.AtualizarStatus(pedido.ID, "Entregue")
	if err != nil {
		fmt.Println("Erro ao atualizar status:", err)
	} else {
		fmt.Printf("Status do pedido atualizado: %+v\n", repo.BuscarPedidoPorID(pedido.ID))
	}
}
```

---

## Prós e Contras do Uso de GRASP

### Prós:

1. **Alta Manutenção do Sistema:** Seguindo GRASP, o código se torna mais fácil de entender e modificar ao longo do
   tempo.
2. **Baixo Acoplamento:** Reduz dependências entre classes, facilitando testes e manutenção.
3. **Alta Coesão:** As classes têm responsabilidades bem definidas, reduzindo complexidade.
4. **Flexibilidade:** As decisões de design podem ser mais facilmente adaptadas quando novos requisitos surgem.
5. **Promove a reutilização:** Controladores e experts podem ser aproveitados em outros módulos ou projetos.

### Contras:

1. **Curva de Aprendizado:** Pode ser difícil para desenvolvedores novatos entender todos os conceitos e aplicá-los
   adequadamente.
2. **Overhead No Design:** Nem todos os projetos requerem o mesmo nível de modularidade ou complexidade oferecido pelo
   GRASP.
3. **Decisões Subjetivas:** Em alguns casos, determinar qual classe é o "expert" ou deve atuar como "controller" pode
   ser ambíguo.

---

## Conclusão

GRASP oferece diretrizes claras para um design orientado a objetos saudável, ajudando a atingir um sistema modular,
flexível e de fácil manutenção. No entanto, o uso de GRASP deve ser ajustado às necessidades do projeto. Projetos
menores podem se beneficiar de abordagens mais simples.

Combinar os princípios de GRASP com boas práticas, como testes automatizados e documentação adequada, pode resultar em
um software mais robusto e eficiente.

---

## Links de Referência

1. Livro: _Applying UML and Patterns_ - Craig Larman
2. [Padrões GRASP (Wikipedia - em inglês)](https://en.wikipedia.org/wiki/GRASP_(object-oriented_design))
3. [Artigo sobre GRASP (em português)](https://www.devmedia.com.br/padroes-de-design-grasp/32215)
