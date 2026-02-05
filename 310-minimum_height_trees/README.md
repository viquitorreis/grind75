Certo, para resolvermos esse problema, primeiro precisamos entender o que é ele:

-> Precisamos retornar os elementos centrais, que formam a menor árvore a partir do grafo não direcionado.

Os nós do centro, vão formar arvores mais baixas, pois estão mais pertos dos outros. Os nós folhas, estão mais longe, por isso formam árvores maiores.

## Intuição:

1. Criar o grafo não direcionado. Armazenando quantos vizinhos ele tem, e quais são.
2. Armazenar o degree (grau) de um vértice, e também as folhas.
3. Processar e ir removendo as folhas (diminuir a quantidade de seus vizinhos), até sobrar no mínimo 2. Atualizar a quantidade de folhas a cada iteração. No fim retorna as folhas restantes, que vão ser os elementos centrais.

**Por quê 1 ou 2 no centro?**

- **Grafo Assímetrico**

Nesse caso, tem um nó que é claramente o ponto do centro

```
      0
      |
      1
     /|\
    2 3 4
```

Ao removermos em camadas:

- Remove folhas [2,3,4]     -> sobra [0, 1]
- Remove folhas [0]         -> sobra [1]

- **Grafo simétrico**

Quando o grafo é simétrico, dois nós são centro igualmente.

```0---1---2---3```

Ao removermos em camadas:

- Remove folhas [0,3]     -> sobra [1,2]
- Remaining = 2           -> para


Um outro exemplo seria:

```
    0
   / \
  1   2
   \ /
    3
```

## 1. Criar o grafo não direcionado. Armazenando quantos vizinhos ele tem, e quais são.

```go
// graph -> connections
	graph := make(map[int][]int)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		// grafo bidirecionado (não direcionado)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
```

## 2. Armazenar o degree (grau) de um vértice, e também as folhas.

Os nós com grau 1 nos indicam as folhas, que são as que vamos iniciar o processamento de diminuir a quantidade de vizinhos (subtrair o degree / grau)

```go
// precisamos encontrar os nós folhas com grau 1
	leaves := []int{}
	degree := make([]int, n)
	for node := 0; node < n; node++ {
		degree[node] = len(graph[node])
		if degree[node] == 1 {
			leaves = append(leaves, node)
		}
	}
```

## 3. Processar e ir removendo as folhas (diminuir a quantidade de seus vizinhos), até sobrar no mínimo 2.

Aqui processamos todas as folhas:

- Remaining: vai monitorar quantos nós ainda faltam ser processados no grafo. Processamos até ser menor ou igual a 2 (elementos centrais)
- Para cada loop, diminuimos o remaining pela quantidade de folhas
- Vamos iterar as folhas, para cada vizinho da folha, vamos diminuir o grau. Se chegar a 1, aumentamos a quantidade de novas folhas.
- Ao final do loop externo, atualizamos as novas folhas para iteração seguinte.

```go
// precisamos remover camada por camada até sobrar no maximo 2 (que podem ser centro)
	remaining := n // quantos nós ainda estão no grafo
	for remaining > 2 {
		numLeaves := len(leaves)
		remaining -= numLeaves

		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--

				// se vizinho vira folha adiciona como leave atual
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		// atualiza leaves para proxima iteração
		leaves = newLeaves
	}

	// vai sobrar 1 ou 2
	return leaves
```