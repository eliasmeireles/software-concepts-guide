# Design Pattern: Adapter (Adaptador)

O Padrão de Projeto Adapter (ou Adaptador) é utilizado para "adaptar" a interface de uma classe ou objeto existente,
tornando-a compatível com outra interface que um cliente espera. Este é um padrão estrutural que atua como um
intermediário entre duas interfaces incompatíveis.

---

## Motivação

Imagine que você está desenvolvendo um sistema que precisa utilizar uma biblioteca externa para realizar determinadas
operações. No entanto, a interface fornecida por essa biblioteca é diferente da interface esperada pelo seu sistema. Em
vez de modificar o código da biblioteca (o que pode ser inviável), você pode criar um Adaptador que "traduza" as
chamadas do seu código para a interface da biblioteca.

---

## Estrutura

O padrão Adapter é composto por:

1. **Classe-alvo (Target):** Define a interface esperada pelo cliente.
2. **Classe existencial (Adaptee):** Classe com a interface incompatível que queremos adaptar.
3. **Adaptador (Adapter):** Classe que implementa a interface esperada pela Classe-alvo e "traduza" as chamadas para a
   classe Adaptee.
4. **Cliente:** A parte do sistema que utiliza a interface do Adaptador para interagir com o Adaptee.

---

## Exemplo em Go

Imagine um sistema que usa componentes de áudio, mas quer integrar uma biblioteca que opera com um sistema de som
diferente.

### Código:

```go
package main

import "fmt"

// Target: Interface que o cliente espera utilizar.
type AudioPlayer interface {
	PlaySound()
}

// Adaptee: Classe com a interface existente e incompatível.
type AdvancedAudioSystem struct{}

func (a *AdvancedAudioSystem) PlayAudio() {
	fmt.Println("Tocando som através do sistema de áudio avançado!")
}

// Adapter: Adapta AdvancedAudioSystem para a interface AudioPlayer.
type AudioAdapter struct {
	Adaptee *AdvancedAudioSystem
}

func (adapter *AudioAdapter) PlaySound() {
	// Tradução da chamada para a interface existente.
	adapter.Adaptee.PlayAudio()
}

// Cliente que usa a interface do Target (AudioPlayer).
func main() {
	// Instanciando o Adaptee.
	advancedAudio := &AdvancedAudioSystem{}

	// Criando o Adapter.
	audioAdapter := &AudioAdapter{Adaptee: advancedAudio}

	// Cliente interage com o Adapter como se fosse um AudioPlayer.
	var player AudioPlayer = audioAdapter
	player.PlaySound()
}
```

### Explicação:

1. **`AudioPlayer` (Target):** Define a interface padrão.
2. **`AdvancedAudioSystem` (Adaptee):** Classe com uma interface incompatível.
3. **`AudioAdapter` (Adapter):** Adapta as chamadas de `AudioPlayer` para a API definida pelo `AdvancedAudioSystem`.
4. **Cliente (`main`):** Usa a interface `AudioPlayer` sem precisar modificar o código do sistema avançado.

---

## Prós e Contras

### Prós:

1. **Reutilização de código:** Permite integrar classes ou bibliotecas existentes sem a necessidade de modificá-las.
2. **Flexibilidade:** Fornece uma solução clara para problemas de incompatibilidade entre interfaces.
3. **Código desacoplado:** O cliente não conhece os detalhes da classe adaptada.

### Contras:

1. **Complexidade adicional:** Pode adicionar camadas extras de lógica, tornando o código mais difícil de entender.
2. **Overhead:** Se mal implementado, pode trazer problemas de performance devido às chamadas indiretas.

---

## Conclusão

O padrão Adapter é uma ferramenta valiosa para lidar com incompatibilidades de interfaces em sistemas. Ele oferece uma
solução elegante para integrar códigos de maneira eficiente e sem modificações no código original. Contudo, é importante
utilizá-lo moderadamente, já que sua aplicação sem necessidade pode aumentar a complexidade do sistema.

---

## Links de Referência

- [Wikipedia - Adapter Pattern](https://en.wikipedia.org/wiki/Adapter_pattern)
- [Refactoring Guru - Adapter Design Pattern](https://refactoring.guru/design-patterns/adapter)
- [GoF Design Patterns - Adapter](https://www.gofpatterns.com/adapter-pattern)
