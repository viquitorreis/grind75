# 01 Matrix

Chave: Inverte o problema

Ao invés de começar de cada célula com valor um e procurar o zero mais próximo, podemos **inverter o problema**.

-> Começar de TODOS os zeros AO MESMO TEMPO e for expandindo para fora. A primeira vez que alcançarmos cada célula vai ser pela distância mais curta.

```
Início (todos os zeros na queue):
[[0,0,0],
 [0,1,0],
 [1,1,1]]

Distâncias iniciais:
[[0,0,0],
 [0,?,0],
 [?,?,?]]
```

O nome desse patern é **Multi-Source BFS**.

**Importante**: Ao fazer o BFS, adicionamos apenas os **quatro vizinhos diretos** de cada célula (cima, baixo, esquerda, direita). Primeiro todas as células na distância 1, depois distância 2, etc...

## Passos

1. Iniciar matriz resultado com todos valores inicialmente infinito ou muito grande, com excessão do 0, que deve ser 0.
2. Precisamos das direções possíveis para cada vizinho (cima,baixo,esquerda, direita). Enquanto a fila não estiver vazia, tiramos o primeiro elemento. Para cada vizinho ainda não visitado, calculamos a distancia dele = distancia do elemento atual + um. Adiciona esse vizinho na fila e marca como visitando trocando a distancia calculada na matriz resultado.

## Visualização

Matriz inicial:

```
[[0, 1, 1],
 [1, 1, 1],
 [1, 1, 0]]
```

1. Após o passo 1:

```
res = [[0, ∞, ∞],
       [∞, ∞, ∞],
       [∞, ∞, 0]]

q = [{x:0, y:0}, {x:2, y:2}]  // os dois zeros que temos
```

2. Aqui começa o loop BFS. Processamos o primeiro elemento da fila, que por acaso é o {x:0, y:0}

A distância desse nó é res[0][0] = 0

Aqui, vamos olhar os 4 vizinhos dele. Vamos olhar apenas o da direita para simplificar:

As coordenadas do vizinho da direita são:

```
newX = curr.x + 0 = 0 + 0 = 0
newY = curr.y + 1 = 0 + 1 = 1
```

Então newX = 0, newY = 1 é a célula a direita do zero (que tem valor um na matriz original).

A distância desse nó até o mais próximo é a distância dele mesmo + 1, pois o vizinho NÃO É ZERO.

Então.

Ele mesmo: ```res[0][0] = 0```
Soma do vizinho (nova coordenada): nó atual + 1

Dessa forma, sempre será incrementado o vizinho uma vez que percorrermos todos os zeros, pois a cada zero processado, ele vai somar um vizinho infinito positivo (logo, que nao foi visitado), + 1.

Assim garantimos o número mínimo, pois a cada nó do grafo, somamos o vizinho apenas uma vez sempre.

```
res[newX][newY] = res[curr.x][curr.y] + 1 // distancia do vizinho
```


## Implementação

É BFS normal:

1. Começa com zeros na queue (distância 0)
2. Processa zeros → descobre vizinhos com distância 1 → adiciona eles na queue
3. Processa células distância 1 → descobre vizinhos com distância 2 → adiciona eles na queue
4. Continua até processar tudo

1. Iniciando a matriz resultado

```go
func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	q := []Coordinate{}

	for i := range len(mat) {
		res[i] = make([]int, len(mat[0]))

		for j := range len(mat[0]) {
			res[i][j] = math.MaxInt64
			if mat[i][j] == 0 {
				q = append(q, Coordinate{
					x: i, // posicação da linha
					y: j, // posição da coluna
				})

				res[i][j] = 0
			}
		}
	
```

2. Precisamos das direções possíveis para cada vizinho (cima,baixo,esquerda, direita). Enquanto a fila não estiver vazia, tiramos o primeiro elemento. Para cada vizinho ainda não visitado, calculamos a distancia dele = distancia do elemento atual + um. Adiciona esse vizinho na fila e marca como visitando trocando a distancia calculada na matriz resultado.

Olha esse grid:

```
0  1  1
```

Se você NÃO adicionar o 1 da posição [0,1] na queue:

Queue: [{0,0}]
Processa {0,0}:

Vizinho direita {0,1} tem distância 1
res[0][1] = 1 ✓
NÃO adiciona na queue

Queue: [] (vazia)
PROBLEMA: {0,2} nunca é processado! Fica MaxInt pra sempre! ❌

Se você ADICIONAR na queue:
Queue: [{0,0}]
Processa {0,0}:

Vizinho {0,1} tem distância 1
res[0][1] = 1
Adiciona {0,1} na queue

Queue: [{0,1}]
Processa {0,1}:

Vizinho {0,2} tem distância 2
res[0][2] = 2 ✓

Queue: []
FUNCIONOU! Todas as células calculadas! ✓

Resumo:
Você adiciona na queue para que essa célula processe os vizinhos DELA depois.
Sem isso, você só processa vizinhos diretos dos zeros. Células mais longe ficam sem calcular.

```go
for len(q) > 0 {
    curr := q[0]
    q = q[1:]

    // precisamos percorrer o elemento atual em todas as direcoes possiveis
    // ou seja, para cada um dos vizinhos
    for _, dir := range directions {
        newX := curr.x + dir[0]
        newY := curr.y + dir[1]

        // checa se o vizinho está dentro dos limites
        if newX >= 0 && newX < len(mat) && newY >= 0 && newY < len(mat[0]) {
            // se nao foi visitada ainda, calculamos a distancia
            if res[newX][newY] == math.MaxInt64 {
                res[newX][newY] = res[curr.x][curr.y] + 1 // distancia do vizinho
                q = append(q, Coordinate{
                    x: newX,
                    y: newY,
                }) // adicionamos o vizinho na queue
            }
        }
    }
}
```