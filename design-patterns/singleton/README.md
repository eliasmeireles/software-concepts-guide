# Design Pattern Singleton (Padrão de Projeto Singleton)

## Descrição

O Singleton é um padrão de projeto criacional que tem como objetivo garantir que uma classe tenha apenas uma única
instância durante o ciclo de vida da aplicação e fornecer um único ponto de acesso a ela. Esse padrão é amplamente usado
em situações onde é necessário controlar um acesso centralizado a recursos compartilhados, como configurações de
aplicativos, registros de log ou conexões a bancos de dados.

## Estrutura

A estrutura básica do Singleton é composta por:

1. **Construtor privado**: Impede que a classe seja instanciada externamente.
2. **Uma variável estática da própria classe**: Responsável por armazenar a única instância da classe.
3. **Método público estático**: Garante a criação ou recuperação da única instância.

---

## Implementação em Go (Golang)

Veja abaixo um exemplo funcional de como implementar o padrão Singleton em Go:

```go
package main

import (
	"fmt"
	"sync"
)

// Singleton define a estrutura da classe Singleton
type Singleton struct {
	message string
}

var (
	instance *Singleton
	once     sync.Once // Garante que a inicialização ocorra somente uma vez
)

// GetInstance cria ou retorna a única instância existente
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{message: "Instância única criada"}
	})
	return instance
}

func main() {
	// Recuperando a instância do Singleton
	s1 := GetInstance()
	fmt.Println(s1.message)

	// Modificando a mensagem da instância
	s1.message = "Instância modificada"

	// Recuperando novamente a instância para confirmar que é a mesma
	s2 := GetInstance()
	fmt.Println(s2.message)

	// Comparando as instâncias
	if s1 == s2 {
		fmt.Println("s1 e s2 são a mesma instância")
	} else {
		fmt.Println("s1 e s2 são instâncias diferentes")
	}
}
```

### **Como funciona?**

1. A função `GetInstance` utiliza o pacote `sync` e o método `Once.Do` para executar apenas uma vez o bloco que cria a
   instância (`instance`).
2. Apenas uma instância da estrutura `Singleton` será criada, independentemente de quantas vezes o método `GetInstance`
   for chamado.
3. `sync.Once` é altamente seguro para uso em contextos concorrentes, garantindo a thread safety do Singleton.

---

## Prós e Contras do Singleton

### **Prós**

- **Controle centralizado**: Garante que apenas uma instância seja criada, facilitando o controle de recursos
  compartilhados.
- **Economia de recursos**: Evita o desperdício de memória com a criação de múltiplas instâncias desnecessárias.
- **Thread-Safe (em Go)**: Com o uso de `sync.Once`, a implementação é segura em ambientes concorrentes.

### **Contras**

- **Acoplamento Global**: Facilita o acoplamento excessivo em diferentes partes do código, dificultando a manutenção e o
  teste.
- **Dificuldade de teste**: Trabalhar com a única instância de uma classe pode complicar a realização de testes
  unitários, especialmente ao simular estados diferentes da instância.
- **Falta de flexibilidade**: Como a classe é responsável por gerenciar sua própria instância, fica mais difícil
  substituir sua implementação ou estendê-la.

---

## Conclusão

O Singleton é uma ferramenta poderosa para gerenciar recursos globais compartilhados. Ele deve ser usado com moderação,
pois pode introduzir problemas de acoplamento, manutenção e teste. No geral, é uma ótima escolha para cenários como
logs, caches e gerenciadores de configuração, mas deve ser evitado sempre que existirem alternativas mais ajustadas ao
problema.

---

## Links de Referência

- [Documentação oficial de Go](https://go.dev/doc/)
- [Design Patterns em Go no refactoring.guru](https://refactoring.guru/design-patterns/singleton/go)
- [Postagem sobre Singleton no blog do Golang Cafe](https://golang.cafe/blog/golang-singleton-pattern.html)
