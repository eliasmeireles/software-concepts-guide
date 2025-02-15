# Padrão de Projeto: Template Method

## Introdução

O padrão de projeto **Template Method** é um padrão comportamental que define a estrutura de um algoritmo em uma classe
base (ou abstrata), mas permite que as subclasses implementem passos específicos desse algoritmo. Ele é comumente
utilizado para reduzir duplicação de código, promover reutilização e facilitar a manutenção, tudo isso enquanto oferece
flexibilidade para personalizar partes específicas do comportamento.

Esse padrão segue o princípio de "definir o esqueleto de um algoritmo e delegar partes do mesmo a subclasses", mantendo
a lógica geral intacta na superclasse.

---

## Estrutura do Template Method

A estrutura básica do padrão pode ser dividida em três componentes principais:

1. **Método Template**: Este método define a sequência de execução do algoritmo. Ele geralmente é implementado na
   superclasse e pode ser marcado como `final` (em linguagens com suporte a herança) para evitar que as subclasses o
   sobrescrevam.
2. **Métodos Concretos (Opcional)**: Métodos definidos na classe base que implementam comportamentos genéricos e podem
   ser reutilizados pelas subclasses.
3. **Métodos Abstratos ou Ganchos (Hook Methods)**: Métodos declarados na superclasse que precisam ser implementados ou
   sobrescritos pelas subclasses para fornecer a funcionalidade específica.

---

## Implementação em Código

Aqui está um exemplo prático do padrão Template Method, implementado em Go, para ilustrar como funciona:

### Exemplo em Go

```go
package main

import "fmt"

// Interface que define o Template
type Template interface {
	Execute()
	Step1()
	Step2()
}

// Estrutura base que implementa o Template Method
type BaseTemplate struct {
	Template
}

// Implementação do método Template
func (t *BaseTemplate) Execute() {
	fmt.Println("Iniciando o algoritmo no Template Method...")
	t.Step1()
	t.Step2()
	fmt.Println("Finalizando o algoritmo no Template Method...")
}

// Subclasse concreta 1
type ConcreteTemplateA struct {
	BaseTemplate
}

// Implementação do primeiro passo específico para Template A
func (c *ConcreteTemplateA) Step1() {
	fmt.Println("ConcreteTemplateA: Executando Step1.")
}

// Implementação do segundo passo específico para Template A
func (c *ConcreteTemplateA) Step2() {
	fmt.Println("ConcreteTemplateA: Executando Step2.")
}

// Subclasse concreta 2
type ConcreteTemplateB struct {
	BaseTemplate
}

// Implementação do primeiro passo específico para Template B
func (c *ConcreteTemplateB) Step1() {
	fmt.Println("ConcreteTemplateB: Executando Step1.")
}

// Implementação do segundo passo específico para Template B
func (c *ConcreteTemplateB) Step2() {
	fmt.Println("ConcreteTemplateB: Executando Step2.")
}

func main() {
	// Criando instâncias das subclasses
	templateA := &ConcreteTemplateA{}
	templateA.Template = templateA

	templateB := &ConcreteTemplateB{}
	templateB.Template = templateB

	// Executando o Template Method com subclasses diferentes
	fmt.Println("Executando Template Method com ConcreteTemplateA:")
	templateA.Execute()

	fmt.Println("\nExecutando Template Method com ConcreteTemplateB:")
	templateB.Execute()
}
```

---

## Explicação do Código

1. **Interface `Template`**: Define a estrutura básica do Template Method com os passos que devem ser seguidos (
   `Execute`, `Step1` e `Step2`).
2. **Estrutura Base `BaseTemplate`**: Implementa o algoritmo geral no método `Execute`. Esta classe delega os passos
   para métodos que devem ser implementados pelas subclasses.
3. **Subclasses Concretas**: `ConcreteTemplateA` e `ConcreteTemplateB` implementam os passos específicos (`Step1` e
   `Step2`), alterando o comportamento, mas mantendo a sequência geral definida em `BaseTemplate`.

Ao executar o código, você verá como o mesmo método `Execute` pode ter resultados diferentes dependendo da implementação
fornecida pelas subclasses.

---

## Conclusão

O padrão Template Method é uma excelente ferramenta para organizar a lógica de algoritmos que compartilham uma estrutura
comum, mas precisam ser personalizados em etapas específicas. Ele incentiva a reutilização de código e o princípio do
código aberto/fechado (Open/Closed Principle), ao mesmo tempo em que facilita a manutenção e evita duplicação de lógica.

Utilizar esse padrão pode ser especialmente útil em aplicações que possuem workflows complexos e repetitivos, permitindo
uma separação clara entre o "como fazer algo" (definido na superclasse) e o "o que será feito exatamente" (definido nas
subclasses). Ele é amplamente aplicado em frameworks e bibliotecas para oferecer flexibilidade ao usuário, permitindo
que detalhes específicos sejam implementados sem alterar a estrutura geral do algoritmo.
