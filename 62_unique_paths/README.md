# Unique Paths

Podemos resolver esse problema de diversas formas, uma delas é:

1. Preencher a primeira linha e primeira coluna com 1
2. Depois cada célula é a **esquerda + cima**

```
Início (m=3, n=3):
[1] [1] [1]
[1] [ ] [ ]
[1] [ ] [ ]
```

Depois disso, preenchemos o resto.

- Para cada célula, de quantas formas posso chegar aqui?

R .:. **somando os caminhos de cima com os caminhos da esquerda**

```
Preenchendo [1][1]:
[1] [1] [1]
[1] [2] [ ]  <- veio de cima (1) + veio da esquerda (1) = 2
[1] [ ] [ ]

Preenchendo [1][2]:
[1] [1] [1]
[1] [2] [3]  <- veio de cima (1) + veio da esquerda (2) = 3
[1] [ ] [ ]

Preenchendo [2][1]:
[1] [1] [1]
[1] [2] [3]
[1] [3] [ ]  <- veio de cima (2) + veio da esquerda (1) = 3

Preenchendo [2][2]:
[1] [1] [1]
[1] [2] [3]
[1] [3] [6]  <- veio de cima (3) + veio da esquerda (3) = 6
```

O **canto inferior direito vai ter a resposta**.

## Passos

1. Preenche a primeira coluna com 1
2. Preenche a primeira linha com 1
3. Preenche o resto -> cima + esquerda

```go
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for y := range matrix {
		matrix[y] = make([]int, n)
	}

	for row := range m {
		matrix[row][0] = 1
	}

	for col := range n {
		matrix[0][col] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			matrix[row][col] = matrix[row-1][col] + matrix[row][col-1]
		}
	}

	return matrix[m-1][n-1]
}
```