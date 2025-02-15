# Padrão de Projeto: Builder (Construtor)

## **O que é o padrão de projeto Builder?**

O padrão *Builder* é um padrão de projeto criacional que tem como objetivo separar a construção de um objeto complexo da
sua representação, permitindo que o mesmo processo de construção possa criar diferentes representações. Este padrão é
especialmente útil quando se trabalha com objetos que possuem muitos atributos opcionais ou que dependem de uma
configuração específica.

O *Builder* facilita a criação de objetos passo a passo, oferecendo uma interface fluida ou através de métodos
específicos. Dessa forma, evita-se o uso de construtores longos e de objetos com vários parâmetros, que podem levar a
confusões e erros.

---

## **Quando usar o Builder?**

- Quando você tem um objeto complexo que requer inicialização ou configurações passo a passo;
- Quando o construtor do objeto possui muitos parâmetros opcionais;
- Quando você deseja criar diferentes representações ou configurações de um objeto sem alterar sua lógica de construção.

---

## **Estrutura do Padrão Builder**

O padrão geralmente envolve os seguintes participantes:

1. **Builder**: Uma interface que especifica os métodos para construir diferentes partes de um objeto.
2. **ConcreteBuilder**: Implementação concreta do Builder, responsável por construir e montar partes específicas do
   objeto.
3. **Director**: Opcional. Um objeto que controla a sequência de chamadas ao *Builder* para criar uma configuração
   específica do objeto.
4. **Produto (Product)**: O objeto complexo que está sendo construído.

---

## **Vantagens (Prós)**

- Facilita a construção de objetos passo a passo.
- Promove a reutilização de código ao permitir a criação de diferentes configurações ou representações do mesmo objeto.
- Melhora a legibilidade do código ao substituir construtores com muitos parâmetros por chamadas mais intuitivas.
- Fornece maior controle sobre o processo de construção do objeto.

---

## **Desvantagens (Contras)**

- Introduz complexidade adicional no projeto devido à criação de uma nova interface e classes associadas.
- Pode ser uma solução desnecessária para objetos que não possuem muitas configurações ou complexidades.
- O uso do padrão pode ser confuso se o número de *builders* e representações se tornar muito grande.

---

## **Exemplo Detalhado em Go**

Abaixo, um exemplo detalhado construído em Go para exemplificar o padrão Builder. O exemplo se refere à criação de um
pedido de um hambúrguer (*Burger*), com diferentes configurações.

### **Código**

```go
package main

import "fmt"

// Produto final (Product)
type Burger struct {
	Bread   string
	Meat    string
	Cheese  bool
	Lettuce bool
	Tomato  bool
}

// Builder interface
type BurgerBuilder interface {
	SetBread(bread string) BurgerBuilder
	SetMeat(meat string) BurgerBuilder
	AddCheese() BurgerBuilder
	AddLettuce() BurgerBuilder
	AddTomato() BurgerBuilder
	Build() Burger
}

// Implementação concreta do Builder
type CustomBurgerBuilder struct {
	burger Burger
}

func NewBurgerBuilder() *CustomBurgerBuilder {
	return &CustomBurgerBuilder{}
}

func (b *CustomBurgerBuilder) SetBread(bread string) BurgerBuilder {
	b.burger.Bread = bread
	return b
}

func (b *CustomBurgerBuilder) SetMeat(meat string) BurgerBuilder {
	b.burger.Meat = meat
	return b
}

func (b *CustomBurgerBuilder) AddCheese() BurgerBuilder {
	b.burger.Cheese = true
	return b
}

func (b *CustomBurgerBuilder) AddLettuce() BurgerBuilder {
	b.burger.Lettuce = true
	return b
}

func (b *CustomBurgerBuilder) AddTomato() BurgerBuilder {
	b.burger.Tomato = true
	return b
}

func (b *CustomBurgerBuilder) Build() Burger {
	return b.burger
}

// Director (opcional)
type BurgerDirector struct {
	builder BurgerBuilder
}

func NewBurgerDirector(builder BurgerBuilder) *BurgerDirector {
	return &BurgerDirector{builder: builder}
}

func (d *BurgerDirector) MakeClassicBurger() Burger {
	return d.builder.SetBread("White").SetMeat("Beef").AddCheese().AddLettuce().Build()
}

func (d *BurgerDirector) MakeVeganBurger() Burger {
	return d.builder.SetBread("Wholegrain").SetMeat("Plant-based").AddLettuce().AddTomato().Build()
}

// Função principal
func main() {
	// Forma simples, usando Builder sem Director
	builder := NewBurgerBuilder()
	customBurger := builder.SetBread("White").
		SetMeat("Chicken").
		AddCheese().
		AddLettuce().
		AddTomato().
		Build()

	fmt.Printf("Custom Burger: %+v\n", customBurger)

	// Usando o Director para criar hambúrgueres pré-definidos
	director := NewBurgerDirector(NewBurgerBuilder())
	classicBurger := director.MakeClassicBurger()
	veganBurger := director.MakeVeganBurger()

	fmt.Printf("Classic Burger: %+v\n", classicBurger)
	fmt.Printf("Vegan Burger: %+v\n", veganBurger)
}
```

---

## **Explicação detalhada do código**

1. **Produto (`Burger`)**: Representa o objeto final que desejamos construir.
2. **Builder Interface (`BurgerBuilder`)**: Define os métodos genéricos para configurar cada parte do objeto.
3. **ConcreteBuilder (`CustomBurgerBuilder`)**: Implementa a interface `BurgerBuilder` para efetivamente construir o
   hambúrguer.
4. **Director (`BurgerDirector`)**: Controla a sequência de ações do *builder* para criar variantes pré-definidas do
   hambúrguer (`MakeClassicBurger`, `MakeVeganBurger`).
5. **Uso prático**:
    - O padrão permite criar configurações customizadas de hambúrguer usando um *builder* direto.
    - Também permite criar configurações específicas e padronizadas usando o *director*.

---

## **Conclusão**

O padrão de projeto *Builder* é extremamente útil quando existe a necessidade de criar objetos complexos ou com várias
configurações. Ele divide a construção em etapas claras e facilmente reutilizáveis, promovendo um código mais legível e
organizado.

No entanto, como qualquer padrão, ele deve ser usado com cautela. Em sistemas simples, com objetos menos complexos, sua
aplicação pode adicionar uma complexidade desnecessária ao código.

---

## **Referências**

- [GoF Design Patterns](https://en.wikipedia.org/wiki/Builder_pattern)
- [Refactoring Guru - Builder Pattern](https://refactoring.guru/design-patterns/builder)
- [Effective Go](https://go.dev/doc/effective_go)
