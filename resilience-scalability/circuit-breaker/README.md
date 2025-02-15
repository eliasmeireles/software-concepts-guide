# Circuit Breaker

O conceito de **Circuit Breaker** (ou disjuntor em português) é amplamente utilizado em desenvolvimento de sistemas para
melhorar a resiliência e a estabilidade de aplicações distribuídas. Ele serve como um padrão de design para isolar e
manipular falhas em partes de sistemas, especialmente em cenários envolvendo chamadas para serviços externos ou
operações que podem falhar.

## Para que serve o Circuit Breaker?

1. **Prevenir falhas cascatas:**  
   Quando uma parte de um sistema fica sobrecarregada ou começa a falhar, o Circuit Breaker impede que outras partes do
   sistema sejam afetadas. Por exemplo, ao evitar chamadas contínuas para um serviço externo que já está indisponível, o
   sistema pode manter um estado consistente.

2. **Manter a estabilidade do sistema:**  
   Ao detectar falhas recorrentes em uma operação, o Circuit Breaker abre o "circuito" e impede chamadas adicionais para
   o serviço problemático, dando tempo para que ele se recupere sem sobrecarregá-lo.

3. **Melhorar a experiência do usuário:**  
   Em vez de forçar o usuário a esperar um longo tempo causado por timeouts frequentes, o Circuit Breaker pode retornar
   uma resposta alternativa ou um erro amigável mais rapidamente.

4. **Registrar métricas e alertar:**  
   O Circuit Breaker coleta métricas que podem ser usadas para monitoramento e alertas sobre a saúde dos serviços de
   terceiros ou de componentes internos.

---

## Como funciona?

O Circuit Breaker funciona em três estados principais:

1. **Closed (Fechado):**  
   O fluxo de chamadas é permitido normalmente. O Circuit Breaker monitora a taxa de erros para determinar se deve "
   abrir" o circuito.

2. **Open (Aberto):**  
   No estado aberto, o Circuit Breaker interrompe todas as chamadas para o serviço problemático. Isso é necessário para
   evitar sobrecarregar o sistema ou deixar o serviço externo ainda mais instável.

3. **Half-Open (Meio Aberto):**  
   Depois de um tempo determinado, o Circuit Breaker entra no estado "meio aberto", permitindo algumas chamadas de teste
   para verificar se o serviço problemático se recuperou. Se as chamadas teste forem bem-sucedidas, o Circuit Breaker
   retorna ao estado fechado. Caso contrário, ele volta ao estado aberto.

---

## Exemplo em Go

Aqui está um exemplo prático de implementação de um Circuit Breaker usando a
biblioteca [Go-resilience](https://github.com/eapache/go-resiliency):

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

func main() {
	// Cria o Circuit Breaker com as configurações:
	// - Máximo de 3 falhas permitidas
	// - Janela de tempo para rastrear falhas
	// - Tempo de espera antes de testar novamente (meio aberto)
	cb := breaker.New(3, 1*time.Second, 5*time.Second)

	for i := 0; i < 10; i++ {
		err := cb.Run(func() error {
			// Simula uma chamada ao serviço externo
			return simulatedRequest(i)
		})

		if err != nil {
			if errors.Is(err, breaker.ErrBreakerOpen) {
				fmt.Println("Circuito aberto, ignorando chamada")
			} else {
				fmt.Println("Erro:", err)
			}
		} else {
			fmt.Println("Chamada bem-sucedida")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func simulatedRequest(attempt int) error {
	// Simula uma falha nas 5 primeiras tentativas
	if attempt < 5 {
		fmt.Println("Simulando falha...")
		return errors.New("falha simulada")
	}
	fmt.Println("Serviço respondeu normalmente")
	return nil
}
```

---

## Vantagens

- Aumenta a resiliência de um sistema distribuído.
- Evita sobrecarregar serviços já falhando.
- Mantém o controle de erros em ambientes críticos.

## Desvantagens

- Introduz complexidade no gerenciamento do fluxo do sistema.
- Requer ajustes finos na configuração (por exemplo, número de erros permitidos, tempos de espera, etc.).

O Circuit Breaker é uma ferramenta essencial em sistemas altamente disponíveis e que dependem de vários serviços ou APIs
externas.