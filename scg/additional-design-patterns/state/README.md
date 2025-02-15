# Padrão de Projeto: State

## Descrição

O padrão de projeto **State** é um **padrão comportamental** que permite que um objeto altere seu comportamento quando
seu estado interno muda. De certa forma, é como se o objeto mudasse de classe durante a execução.
Esse padrão é útil para situações em que um objeto deve se comportar de maneira diferente dependendo de seu estado
atual, eliminando a necessidade de longas cadeias de condicionais ou switches no código.

## Intenção

- Permitir que um objeto altere seu comportamento quando o seu estado interno muda.
- Encapsular os diferentes comportamentos em classes separadas, facilitando a separação de responsabilidades e a
  manutenção do código.

## Aplicação

- Quando um objeto apresenta vários estados diferentes e o comportamento muda dependendo do estado.
- Quando o código possui muitas estruturas condicionais (como `if/else` ou `switch`) para lidar com os estados de um
  objeto.

## Estrutura

O padrão State geralmente envolve quatro componentes principais:

1. **Context**: A classe que mantém uma referência para o estado atual e delega o comportamento para o mesmo.
2. **State**: Interface que define o comportamento que será implementado pelos estados concretos.
3. **States Concretos**: Implementações específicas de diferentes estados que alteram o comportamento do objeto quando
   ele está naquele estado.
4. **Transição de Estados**: A lógica para trocar de estado está encapsulada dentro dos estados concretos ou no
   contexto.

## Exemplo em Código (Go)

Abaixo está um exemplo de implementação do padrão State em Go.

```go
package main

import "fmt"

// State interface define o comportamento esperado de cada estado
type State interface {
	Handle(context *Context)
}

// ConcreteStateA representa um estado específico
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Estado: A. Transição para o Estado B.")
	context.SetState(&ConcreteStateB{})
}

// ConcreteStateB representa outro estado específico
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Estado: B. Transição para o Estado A.")
	context.SetState(&ConcreteStateA{})
}

// Context mantém uma referência ao estado atual
type Context struct {
	state State
}

func NewContext() *Context {
	return &Context{
		state: &ConcreteStateA{}, // Estado inicial
	}
}
func (c *Context) SetState(state State) {
	c.state = state
}
func (c *Context) Request() {
	c.state.Handle(c)
}
func main() {
	// Inicializa o contexto
	context := NewContext()
	// Faz algumas alterações de estado
	context.Request() // Estado A -> Estado B
	context.Request() // Estado B -> Estado A
	context.Request() // Estado A -> Estado B
}
```

## Explicação do Código

1. **Context**: A estrutura `Context` mantém uma referência ao estado atual (`state`) e fornece métodos para alterar
   esse estado ou delegar ações.
2. **State**: Uma interface que define o método `Handle(context *Context)`, garantindo que todos os estados respeitem o
   mesmo comportamento.
3. **ConcreteStateA e ConcreteStateB**: Classes que implementam a interface `State` e possuem comportamentos
   específicos. Elas determinam as transições para outros estados.
4. **Main Function**: Demonstra o funcionamento do padrão, com transições entre os estados e chamadas para ações
   desempenhadas por cada estado.

## Vantagens

- Elimina o uso de muitas estruturas `if/else` ou `switch`.
- Torna o código mais flexível e fácil de modificar.
- Facilita a adição de novos estados sem modificar código existente.

## Desvantagens

- Pode aumentar o número de classes no sistema.
- A lógica de transição de estados pode se tornar complexa dependendo do número de estados.

## Quando Usar?

- Quando o comportamento de um objeto precisa mudar dependendo de seu estado.
- Quando você deseja evitar grandes cadeias de condicionais que verificam o estado do objeto.

## Conclusão

O padrão de projeto **State** é uma solução elegante para problemas em que o comportamento de um objeto muda
dinamicamente com base em seu estado interno. Ele facilita a organização e manutenção do código, promovendo o princípio
da responsabilidade única. No entanto, sua aplicação exige cautela, pois pode aumentar a complexidade do sistema devido
ao número de classes envolvidas.
Ao utilizá-lo, é importante avaliar a necessidade de flexibilizar transições de estado, sempre equilibrando clareza de
projeto e escalabilidade, tornando o sistema mais robusto e eficaz em contextos que requerem mudanças frequentes de
comportamento.
