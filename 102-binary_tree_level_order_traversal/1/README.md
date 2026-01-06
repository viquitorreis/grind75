
Binary Tree é uma árvore não balanceada. Onde o elemento da esquerda pode não ser o menor elemento.

# <span style="color:rgb(0, 176, 240)">In Order traversal</span>

https://medium.com/geekculture/level-order-traversal-of-binary-trees-in-go-311b05a2abcf

É atravessar a árvore por ordem de nível (andar).

Na order traversal:

1. A root é visitada primeiro
2. O filho imediato
3. Os netos imediatos

etc...

Precisamos visitar todos os nodes em um level antes de visitar o próximo node.

## <span style="color:rgb(186, 82, 255)">Fila - FIFO</span>

Podemos implementar uma <span style="color:rgb(0, 255, 136)"><b>fila - FIFO</b>.</span> 

```
Numa fila, os elementos inseridos em uma ponta, saem na outra ponta.
```

Precisamos, no mínimo das seguintes operações:

- **insert:** inserir um novo elemento no fim da final
- **remove:** remover um elemento do início da fila
- **len:** total de item da fila

Podemos criar um tipo customizado para nossa fila em go, e acessar os métodos a partir dela, mas não precisa. Podemos usar slices simples:

```go
// Create a queue - a slice of integers here
queue := []int{}

// Insert an item at the end
// a new item into the slice at the end
queue = append(queue, 10)

// Remove an item from the queue
item := queue[0]
queue = queue[1:]

// Get the length of the queue
// len is a built-in function in Go
numberOfItems := len(queue)
```

## <span style="color:rgb(186, 82, 255)">Passos do Algoritmo</span>

### Passo 1:

Checar se a root é nula. Se for, evita todo o processo e retorna nada...

```go
if root == nil {
  // nada para fazer
  return [][]int{}
}
```

### Passo 2: Criar um slice de inteiros para o resultado final

```go
result := [][]int{}
```

### Passo 3: Cria a fila

Cria uma fila e insere a root na queue. Ou seja, começamos o traversal da root.

```go
queue := []*TreeNode{}
queue = append(queue, root)
```

### Passo 4: Processamento da Fila

Enquanto a fila não estiver vazia, processamos os elementos da fila.

<span style="color:rgb(255, 149, 0)"><b>Nota:</b></span> o tamanho total da fila indica a quantidade de elementos que existem naquele nível.

```go
// continua processando enquanto a fila não estiver vazia
for len(queue) > 0 {
  // pega o tamanho atual da fila -> quantidade de nós que devem estar no nível atual
  sz := len(queue)
  level := []int{}
  // processa "sz" número de elementos da fila, e enche os valores no slice "level"
}
```

### Passo 5: Processa cada nó pertencente ao mesmo nível

O processamento envolve:

1. Remover o node da fila
2. Inserir o valor num array temporário que armazena os nodes do nível atual
3. Inserir os filhos na filha para processamento posterior

```go
for i := 0; i < sz; i++ {
    // remove um node da fila (primeiro elemento)
    node := queue[0]
    queue = queue[1:]

    // ao visitar o node, coloca ele no array de output
    level = append(level, node.Val)
    
    // insere o filho do node na fila
    if node.Left != nil {
      queue = append(queue, node.Left)
    }
    if node.Right != nil {
      queue = append(queue, node.Right)
    }
}
```

### Passo 6:

Depois que todos os nodes que pertenciam ao mesmo nível forem processados, o output do nível atual é inserido no resultado final.

```go

```

### Resultado final

```go
func levelOrder(root *TreeNode) [][]int {
  if root == nil {
	// nada para fazer
    return [][]int{}
  }

  // cria a file e insere a raiz
  queue := []*TreeNode{}
  queue = append(queue, root)

  // slice dos resultados final
  result := [][]int{}

  // processa enquanto a fila não estiver vazia
  for len(queue) > 0 {
    // pega o tamanho atual da fila
    // indica o tamanho total de nodes que são parte do nível atual
    sz := len(queue)
    level := []int{}
    for i := 0; i < sz; i++ {
      // Remove o node da fila (primeiro elemento)
      node := queue[0]
      queue = queue[1:]

      // ao visitar o node, coloca ele no array de output
      level = append(level, node.Val)

      // insere os filhos do node na fila
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }
    
    // level é preenchido com o nível atual de nodes. depois disso insere no resultado final
    result = append(result, level)
  }
  // result is ready to be returned
  return result
}
```