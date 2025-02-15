# Padrão de Projeto: Composite

## Visão Geral

O **Composite** é um padrão de projeto estrutural que permite que você trate de maneira uniforme objetos individuais e
composições de objetos. Ele organiza os objetos em uma estrutura hierárquica de árvore, onde tanto os objetos simples
quanto os compostos podem ser tratados como se fossem objetos individuais.

Este padrão ajuda na criação de estruturas que representam árvores, como sistemas de arquivos, representações gráficas
ou qualquer outra estrutura hierárquica.

---

## Intenção

> "Compor objetos em estruturas de árvore para representar hierarquias parte-todo. O padrão **Composite** permite que os
clientes tratem objetos individuais e composições de objetos de maneira uniforme."  
> — *Design Patterns: Elements of Reusable Object-Oriented Software (Gang of Four)*

---

## Aplicabilidade

Use o padrão **Composite** quando:

1. Você tiver que implementar uma estrutura hierárquica, como árvores.
2. Você quiser que o cliente interaja com os elementos simples e compostos de forma uniforme, sem depender de suas
   diferenças.
3. Você precisa de operações recursivas sobre a estrutura de dados.

---

## Estrutura

A estrutura do Composite é geralmente definida com os seguintes participantes:

1. **Componente (Component)**  
   Define a interface comum para objetos individuais e compostos. Declara métodos que devem ser implementados tanto
   pelos objetos simples quanto pelos compostos.

2. **Folha (Leaf)**  
   Representa os objetos individuais na composição que não contêm nenhum filho. Implementa os métodos definidos pela
   interface do componente.

3. **Composição (Composite)**  
   Representa objetos que podem ter filhos (outras folhas ou compostos). Implementa métodos para gerenciar os objetos
   filhos e delegar operações para eles.

---

## Exemplo com Código

Aqui está um exemplo de implementação simples do padrão **Composite** em Go para representar uma estrutura de hierarquia
de arquivos:

```go
package main

import "fmt"

// Component: Define a interface comum
type FileSystemComponent interface {
	Display(indentation string)
}

// Leaf: Representa arquivos individuais
type File struct {
	Name string
}

func (f *File) Display(indentation string) {
	fmt.Println(indentation + f.Name)
}

// Composite: Representa diretórios (que podem conter arquivos ou outros diretórios)
type Directory struct {
	Name     string
	Children []FileSystemComponent
}

func (d *Directory) Add(component FileSystemComponent) {
	d.Children = append(d.Children, component)
}

func (d *Directory) Display(indentation string) {
	fmt.Println(indentation + d.Name + "/")
	for _, child := range d.Children {
		child.Display(indentation + "  ")
	}
}

func main() {
	root := &Directory{Name: "root"}
	home := &Directory{Name: "home"}
	user := &Directory{Name: "user"}
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}
	pic := &File{Name: "photo.jpg"}

	// Construindo a hierarquia
	user.Add(file1)
	user.Add(file2)
	home.Add(user)
	home.Add(pic)
	root.Add(home)

	// Exibe a hierarquia
	root.Display("")
}

```

### Saída do Código

```plaintext
root/
  home/
    user/
      file1.txt
      file2.txt
    photo.jpg
```

---

## Benefícios

- **Simplicidade de design**: Permite tratar folhas e composições de forma uniforme.
- **Facilidade de extensão**: Adicionar novos tipos de componentes (folhas ou composições) é relativamente simples.
- **Flexibilidade**: Funciona bem com estruturas hierárquicas dinâmicas.

---

## Desvantagens

- Pode exigir muito cuidado ao gerenciar referências recursivas em estruturas grandes.
- Tornar-se complexo se houver muitas operações variantes entre folhas e composições.

---

## Comparação com Outras Estruturas

Embora o **Composite** seja frequentemente usado para representar hierarquias, ele é diferente de outros padrões como:

- **Decorator** – Dá funcionalidade adicional sem alterar a estrutura hierárquica.
- **Flyweight** – Evita a duplicação de objetos na memória.

---

## Quando Não Usar

Evitar o **Composite** caso:

1. Não haja uma hierarquia clara nos objetos.
2. A performance seja crítica, já que operações recursivas podem ser custosas.

---

## Resumo do Composite

O padrão **Composite** é uma solução elegante para problemas de hierarquia parte-todo, permitindo que você trate folhas
e composições de maneira similar. Ideal para representar sistemas hierárquicos como diretórios, gráficos de elementos ou
documentos com partes aninhadas.
