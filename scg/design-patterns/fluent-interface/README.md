# Design Pattern: Fluent Interface

## O que é Fluent Interface?

O Design Pattern Fluent Interface é criado para melhorar a legibilidade e a concisão de um código, permitindo o
encadeamento de chamadas de métodos. Ele é particularmente útil quando você deseja construir objetos complexos ou
configurar dados, pois oferece um estilo expressivo e intuitivo de escrita.

A principal ideia desse padrão é permitir que os métodos retornem a própria instância do objeto (`this` ou similar),
possibilitando o encadeamento fluente de métodos.

---

## Características principais:

1. **Encadeamento de métodos:** Métodos na interface retornam o objeto atual, evitando chamadas separadas.
2. **Legibilidade aumentada:** Facilita a leitura e compreensão do código por se parecer com uma linguagem natural.
3. **Sem estados intermediários:** Não é necessário criar variáveis adicionais para as configurações entre métodos.

---

## Como funciona?

No Fluent Interface, métodos são projetados para chamarem uns aos outros consecutivamente, sendo responsáveis por
configurar diferentes aspectos de um objeto ou realizar uma ação maior.

---

## Implementação em Go: Exemplo do Design Pattern Fluent Interface

Em Go, podemos implementar o Fluent Interface aproveitando structs e métodos que retornam o ponteiro da própria struct.
Aqui está um exemplo:

### Um Builder para construir um objeto `Person`:

```go
package main

import "fmt"

// Definimos a estrutura Person
type Person struct {
	name  string
	age   int
	email string
}

// Adicionando métodos para aplicar o padrão Fluent Interface

// Método para configurar o nome
func (p *Person) SetName(name string) *Person {
	p.name = name
	return p
}

// Método para configurar a idade
func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}

// Método para configurar o email
func (p *Person) SetEmail(email string) *Person {
	p.email = email
	return p
}

// Método para exibir as informações da pessoa
func (p *Person) Build() {
	fmt.Printf("Person(Name: %s, Age: %d, Email: %s)\n", p.name, p.age, p.email)
}

func main() {
	// Exemplo de uso do Fluent Interface
	person := &Person{}
	person.SetName("João").
		SetAge(30).
		SetEmail("joao@email.com").
		Build()
}
```

---

### Explicação do exemplo

1. **Definição da Estrutura:** Temos uma estrutura `Person`, onde há campos básicos como `name`, `age` e `email`.
2. **Encadeamento de Métodos:** Os métodos (`SetName`, `SetAge`, `SetEmail`) retornam o ponteiro para o próprio objeto (
   `*Person`), permitindo a construção fluida.
3. **Configuração em uma linha:** No `main()`, construímos uma instância da struct `Person` e configuramos seus valores
   com uma única linha fluida.

---

## Prós e Contras do Fluent Interface

### **Prós**

1. **Leitura clara:** O código se torna expressivo, especialmente ao configurar objetos com muitos campos ou passos.
2. **Menos variáveis temporárias:** Não é necessário criar variáveis intermediárias para armazenar resultados entre
   operações.
3. **Código modularizado:** Cada método configura apenas uma parte/ação, favorecendo a manutenção.

### **Contras**

1. **Debug mais difícil:** Em casos de encadeamento longo, é mais difícil identificar onde ocorre um erro.
2. **Limitações com valores imutáveis:** Caso sua estrutura deva ser imutável, o Fluent Interface se torna mais difícil
   de implementar.
3. **Potencial para uso abusivo:** Pode levar a mensagens encadeadas confusas ou difíceis de entender.

---

## Conclusão

O Fluent Interface é um padrão poderoso que melhora significativamente a legibilidade e o estilo de alguns códigos,
especialmente para tarefas de configuração e construção de objetos. Em Go, ele se adapta bem ao usar métodos que
retornam o ponteiro para a mesma struct, facilitando o encadeamento fluido de chamadas.

No entanto, apesar de suas vantagens de legibilidade, ele deve ser usado com cuidado quando a simplicidade é prioritária
ou quando o rastreamento de erros seria prejudicado.

---

## Referências:

1. [Martin Fowler - Fluent Interface](https://martinfowler.com/bliki/FluentInterface.html)
2. [Refactoring Guru - Fluent Interface](https://refactoring.guru/design-patterns/fluent-interface)
3. [Go Programming Language Documentation](https://go.dev/doc/)
