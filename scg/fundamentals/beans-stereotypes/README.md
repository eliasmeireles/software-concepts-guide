## Fundamentos de Programação: Stereotypes de Beans

Os *stereotypes* de beans são um conceito importante na programação orientada por objetos, especialmente em frameworks
que implementam injeção de dependência como o Spring Framework (usado em Java), ou ideias correlatas que podem ser
traduzidas ou adaptadas para outras linguagens como o Go. O objetivo principal de *stereotypes* é categorizar e rotular
classes para que o framework ou programa saiba como lidar com elas no contexto de gerenciamento de dependências.

### O que são Stereotypes de Beans?

No contexto de desenvolvimento de software, um *stereotype* de bean é um marcador atribuído a uma classe, indicando o
seu papel ou propósito dentro da aplicação. Por exemplo:

- Classes que representam componentes da lógica de negócios.
- Repositórios que interagem diretamente com bancos de dados.
- Controladores que lidam com a interação com cliente (API, interface gráfica, etc.).

Linguagens como Go não possuem *stereotypes* de forma nativa, mas é possível adotar práticas similares por meio de
convenções e uso de interfaces, separação de responsabilidades e uso de injeção de dependência manual ou com bibliotecas
específicas.

---

### Implementação em Go: Exemplo

Embora Go seja uma linguagem sem suporte nativo para frameworks pesados como o Spring, podemos implementar práticas
similares às de *stereotypes* por meio de organização do código. Aqui está um exemplo simples:

#### Exemplo em Go

Suponha que queremos implementar e categorizar diferentes funções em uma aplicação como `controllers`, `services` e
`repositories` seguindo princípios de *stereotypes*.

```go
package main

import (
	"fmt"
)

// Stereotype: Service
type GreetingService interface {
	GetGreeting(user string) string
}

type GreetingServiceImpl struct{}

func (g *GreetingServiceImpl) GetGreeting(user string) string {
	return fmt.Sprintf("Olá, %s! Bem-vindo ao sistema.", user)
}

// Stereotype: Controller
type UserController struct {
	service GreetingService
}

func NewUserController(service GreetingService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) HandleRequest(user string) {
	greeting := c.service.GetGreeting(user)
	fmt.Println(greeting)
}

// Stereotype: Repository (simulado para contexto mais amplo)
type UserRepository interface {
	SaveUser(user string) error
}

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) SaveUser(user string) error {
	fmt.Printf("Usuário '%s' salvo no banco de dados.\n", user)
	return nil
}

func main() {
	// Injeção manual de dependências (forma comum em Go)
	service := &GreetingServiceImpl{}
	controller := NewUserController(service)

	controller.HandleRequest("João")
}
```

#### Explicação do Código

1. `GreetingService` e sua implementação `GreetingServiceImpl` representam um *Service*.
2. `UserController` age como o *Controller*, interagindo com o cliente e usando o serviço para fornecer as informações
   necessárias.
3. Um exemplo de repostório (simulado) mostra como persistência de dados seria comunicada.
4. A injeção de dependência é feita manualmente ao criar as instâncias no `main()`.

---

### Prós e Contras de Usar Stereotypes em Go

#### Prós

1. **Organização**: O uso deste padrão facilita a estruturação do código, separando responsabilidades e tornando o
   código mais compreensível.
2. **Teste**: Com a separação clara de responsabilidades, fica mais fácil realizar testes unitários em partes isoladas
   do código.
3. **Legibilidade**: O uso de convenções como essas melhora a legibilidade, especialmente em equipes que desejam adotar
   boas práticas.

#### Contras

1. **Complexidade Extra**: O modelo baseado em *stereotypes* pode acrescentar complexidade ao código, o que pode não ser
   ideal para pequenas aplicações.
2. **Frameworks Pesados**: Em linguagens que possuem suporte nativo para isso (como Java), o uso extensivo pode criar
   dependência de frameworks.
3. **Injeção Manual**: Em Go, a necessidade de fazer injeção manual de dependências pode ser um esforço adicional se
   comparado a frameworks que fazem isso automaticamente.

---

### Conclusão

*Stereotypes* de beans são um conceito importante para organizar e estruturar aplicativos grandes. Embora Go não os
suporte diretamente como Java, práticas inspiradas neles podem ser empregadas para criar aplicativos escaláveis,
organizados e mais fáceis de manter. O exemplo acima mostra uma maneira simples de implementar algo similar em Go.

---

### Links de Referência

1. [Documentação sobre Dependency Injection em Go (Martin Fowler)](https://martinfowler.com/articles/injection.html)
2. [Go Design Patterns: Organizing Code by Responsibility](https://golang.org/doc/effective_go)
3. [Artigo: Stereotype Annotations in Spring](https://spring.io/guides)
