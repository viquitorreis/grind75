Esse problema, resolvemos com **DP e binary search**

## O problema da decisão

Imagina que estamos vendo os jobs **ordenados por endTime crescente**. E estamos no job `i` e precisamos decidir:

1. PEgo esse job? Ganho `profit[i]`, mas agora, só posso epgar jobs que começam depois de `endTime[i]`
2. **Pulo esse job?** Não ganho nada dele, mas mantenho todas as opções anteriores

É tipo o **knapsack**: cada item você decide pegar ou não.

## Parte principal - DP

Se já sabemos o **máximo de lucros possível considerando os jobs de 0 até i-1**, então para decidir sobre o job i, só precisamos:

- **Se pular**: o lucro é o mesmo que `dp[i - 1]` (o máximo até o job anterior)
- **Se pegar**: o lucro é `profit[i]` + o máximo lucro dos jobs que **terminaram antes** de `startTime[i]`

Isso significa que `dp[i]` depende de:

1. `dp[i - 1]` (opção de pular)
2. `profit[i] + dp[j]`, onde `j` é o último job que terminou antes de `startTime[i]` começar

---

## Onde usamos Binary Search?

Para achar esse `j` (ultimo job compatível), não queremos fazer um loop linear, que seria O(n²). Como os jobs estão ordenados por `endTime`, podemos usar **binary search** para achar o índice **j** do último job onde `endTime <= startTime[i]`

**De forma simples**: acha o ultimo job que não tem overlap de tempo até o i atual

**findLastCompatible**

No nosso loop de dp, quando estamos no job `i` e decidimos pegá-lo, precisamos achar o **índice do último job entre 0 e i-1 que terminou antes ou exatamente quando o job i começou**.

Ou seja: procurar o maior índice `j` tal que `jobs[j].endTime <= jobs[i].startTime`

Como os jobs estão ordenados por `endtime` crescente, teremos:

```
jobs[0].endTime <= jobs[1].endTime <= jobs[2].endTime ...
```

Isso significa pois estamos procurando o último que terminou antes de X, estamos **procurando uma fronteira no array ordenado**:

```
jobs:     [endTime=2] [endTime=4] [endTime=5] [endTime=7] [endTime=9]
               ✓           ✓           ✓           ✗           ✗
          (compatíveis com startTime=6)    (incompatíveis)
                                      ↑
                                  último compatível
```

```go
func findLastCompatible(jobs []job, i int) int {
	left, right := 0, i-1
	candidate := -1

	for left <= right {
		mid := (left + right) >> 1
		if jobs[mid].endTime <= jobs[i].startTime {
			// procuramos se existe um job mais recente
			candidate = mid
			left = mid + 1
		} else {
			// job termina tarde, vai para a direita
			right = mid - 1
		}
	}

	return candidate
}
```

A função findLastCompatible não compara profit nenhum. Ela apenas acha o último job compatível por tempo. Então como isso resolve o problema de maximizar lucro?

A mágica tá no **array de DP** e não no findLastCompatible. Quando a chamada ao findLastCompatible retorna, não vamos pegar o job de maior lucro dela, e sim de `dp[j]`, que é muito diferente.

**dp[j] = o máximo lucro possível considerando TODOS os jobs de 0 até j, escolhendo a melhor combinação deles.**

Então quando você faz take = jobs[i].profit + dp[j], você não está somando apenas o profit do job j. Você está somando o profit do job i com a melhor estratégia possível que terminou no job j ou antes.

```
O dp guarda o resultado da melhor escolha, não qual job foi escolhido.
```

.:. Então:

```
 o dp guarda o melhor profit até aquele indice, e quando pegamos o ultimo tempo compatível, o dp[j] quem pega o maior profit possível para aquele job (indice j)
```

---

## Estrutura da solução

1. Ordena jobs por `endTime` crescente (vai permitir o DP sequencial e o binary search)
2. Cria o array `dp` onde `dp[i]` = máximo lucro considerando jobs 0 até 1
3. Para cada job `i`:
    - Acha o último job compatível `j` (binary search em `endTime`)
    - `dp[i] = max(dp[i - 1], profit[i] + dp[j])` -> send j o último job compatível
4. Resposta é `dp[n - 1]`

### Por que ordenar por endTime

Pois precisamos saber qual é **a informação que precisamos acessar de forma eficiente**.

Ou seja: quando estamos no job `i` e decidimos pegá-lo, precisamos saber **qual o máximo de lucro dos jobs que terminarem antes de `startTime[i]`.

Se ordenamos por `endTime`:

- Todos os jobs antes de `i` no array já terminaram em tempo `<= endTime[i]`
- Fazemos o binary search para achar o último job onde `endTime <= startTime[i]`
- O `dp[j]` desse job já tem o máximo acumulado até ele


### Código

**take**: Quando estamos processando o job `i` temos duas opções:
- Pegar esse job (**take**): ganhamos `jobs[i].profit` + maximo lucro que conseguimos até o ultimo job compativel
- Pular esse job (**skip**): não ganha nada desse job, então o lucro continua sendo o mesmo que tinhamos até o job anterior (`dp[i-1]`)


```go
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := []job{}
	for i := range startTime {
		jobs = append(jobs, job{
			profit:    profit[i],
			startTime: startTime[i],
			endTime:   endTime[i],
		})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].endTime < jobs[j].endTime
	})

	dp := make([]int, len(profit))
	dp[0] = jobs[0].profit

	for i := 1; i < len(jobs); i++ {
		// achamos o indice do ultimo job compatível com o i
		j := findLastCompatible(jobs, i)

		take := jobs[i].profit
		if j != -1 {
			// pegamos o maximo de lucro que o ultimo job compatível teve, o array dp vai nos dizer isso
			take += dp[j]
		}

		// quando decidimos pular o job i, estamo pensando "não vou fazer esse job
		// vou manter exatamente o resultado que tinha antes de olhar para ele, que é o indice i-1 no dp
		skip := dp[i-1]

		// max entre:
		// 		- soma do profit do job atual + profit maximo do ultimo job compativel
		// 		- soma do profit no indice anterior do dp - sem contar esse job
		dp[i] = max(take, skip)
	}

	return dp[len(dp)-1]
}
type job struct {
	profit    int
	startTime int
	endTime   int
}
```