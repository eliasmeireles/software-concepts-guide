# Fluent Interface

O padrão **Fluent Interface** é um estilo de design de API que proporciona uma experiência mais legível e expressiva ao
encadear chamadas de métodos. Através deste padrão, métodos retornam o próprio objeto, permitindo que várias operações
sejam chamadas em sequência.

## Como funciona?

O Fluent Interface utiliza o conceito de encadeamento de métodos (method chaining), permitindo que o código seja escrito
de forma mais linear e fluida. Para isso, os métodos da classe retornam a própria instância (`this` em outras linguagens
ou `*struct` no caso do Go) para permitir chamadas contínuas.

---

## Prós

- **Legibilidade**: Torna o código mais legível e expressivo, simulando quase uma descrição natural do que o código faz.
- **Redução da verbosidade**: Menos chamadas repetidas e desnecessárias ao mesmo objeto.
- **Facilidade na configuração**: Ideal para configurações complexas de objetos ou inicializações.

---

## Contras

- **Depuração complicada**: Erros em chamadas encadeadas podem ser difíceis de rastrear, especialmente com linhas
  longas.
- **Não intuitivo para iniciantes**: O conceito de retornar a própria instância em métodos pode gerar confusão.
- **Encadeamento rígido**: Pode dificultar a extensão e a reutilização do código, já que o fluxo de chamadas é linear.

---

## Exemplo em Go

Aqui está um exemplo prático de um **Fluent Interface** em Go para uma configuração de uma conexão com banco de dados:

```go
package main

import "fmt"

// Config é a estrutura que implementa o padrão Fluent Interface
type Config struct {
	host     string
	port     int
	username string
	password string
}

// SetHost define o host da conexão e retorna a própria instância
func (c *Config) SetHost(host string) *Config {
	c.host = host
	return c
}

// SetPort define o port da conexão e retorna a própria instância
func (c *Config) SetPort(port int) *Config {
	c.port = port
	return c
}

// SetUsername define o nome de usuário e retorna a própria instância
func (c *Config) SetUsername(username string) *Config {
	c.username = username
	return c
}

// SetPassword define a senha e retorna a própria instância
func (c *Config) SetPassword(password string) *Config {
	c.password = password
	return c
}

// Connect simula a conexão usando a configuração definida
func (c *Config) Connect() {
	fmt.Printf("Conectando ao host %s:%d como %s... Conexão estabelecida!\n", c.host, c.port, c.username)
}

func main() {
	// Utilizando o padrão Fluent Interface para configurar e conectar
	config := &Config{}
	config.
		SetHost("localhost").
		SetPort(5432).
		SetUsername("admin").
		SetPassword("senha123").
		Connect()
}
```

---

## Conclusão

O Fluent Interface é um padrão útil para situações em que configuração e inicialização fluida são necessárias, como na
criação de objetos complexos ou na construção de DSLs internas. Apesar de suas limitações em rastreamento de erros e
flexibilidade, pode ser implementado de forma eficiente em Go quando usado com moderação e planejamento.

---

## Links de referência

- [Martin Fowler - Fluent Interface](https://martinfowler.com/bliki/FluentInterface.html)
- [Go Programming Language Documentation](https://go.dev/doc/)
