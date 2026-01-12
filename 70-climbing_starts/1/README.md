# Climbing Stairs

Esse é um problema básico de programação dinâmica. Precisamos retornar todas as combinações possíveis de se chegar ao topo.
Podemos subir apenas um ou dois degraus.

## Abordagem Bottom Up

Nessa abordagem, vamos calcular as respostas para nossos base cases.

Quando n <= 3 e n > 0. Então n = n.

Porém, quando n = 4, muda, pois teremos 5 opções de chegar ao topo:

1 - 1,1,1,1
2 - 1,1,2
3 - 1,2,1
4 - 2,1,1
5 - 2,2

### Por quê Bottom Up?

Nesse caso, a partir de 3, estamos fazendo a **soma dos passos anteriores**:

- n = 1 -> 1
- n = 2 -> 2
- n = 3 -> 3 (steps = 2 + 1)
- n = 4 -> 5 (steps = 3 + 2)
- n = 5 -> 8 (steps = 4 + 3)

Basicamente, é uma sequencia fibonacci que começa no 1

### Como?

Por exemplo, n = 3, temos as seguintes combinações:

- 1,1,1
- 1,2
- 2,1

Nesse sentido, quando chegamos no degrau 3, podemos ter vindo de qual degrau?

```
Do 1 ou do 2
```

Ou seja:

- Todas as formas de chegar no degrau 2 - ultimo degrau anterior (+ 1)
- Todas as formas de chegar no degrau 1 - penultimo degrau anterior (+ 2)

Portanto,  precisamos guardar os previous (-1 e -2 de onde estamos).

Nosso código pode ficar assim:

```go
func climbStairs(n int) int {
	// base case
	if n <= 2 {
		return n
	}

	// dynamic programming
	res := []int{}
	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		// previous mudaram agora
		prev1 = prev2
		prev2 = current

		res = append(res, current)
	}

	return res[len(res)-1]
}
```

Mas não precisamos guardar toads as respostas, pois o prev2 será o que queremos, pois é o número mais atual sempre, ou seja, quantos passos para chegar no andar atual.

```go
func climbStairs(n int) int {
	// base case
	if n <= 2 {
		return n
	}

	// dynamic programming
	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		// previous mudaram agora
		prev1 = prev2
		prev2 = current
	}

	return prev2
}
```

