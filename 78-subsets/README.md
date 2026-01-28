# Subsets

Esse é um problema de **backtracking**, e para entender ele, precisamos entender como funciona a árvore de decisões.

## Backtracking

```go
dfs = func(i int) {
```

O i é o índice do elemento atual, ele representa o índice do array que estamos processando agora. Quando temos nums = [1,2,3], os índices são 0, 1 e 2.

Quando i chega em 3, significa que já passou por todos os elementos e tomou uma decisão para cada um deles (incluir ou não incluir). Nesse momento, chegamos em uma "folha" da árvore de decisões, e precisamos guardar o subset que construímos até ali.

Então, como no exemplo ficaria indice de 0 até 2:

1. Primeira decisão no índice 0
2. Segunda decisão no índice 1
3. Terceira decisão no índice 2

Após isso, ao chamar o backtracking novamente, o índice vai ser maior que o tamanho de nums, e vai entrar na condição de parada:

```go
dfs = func(i int) {
		if i >= len(nums) {
```

### Por que duas camadas recursivas?

Nso problemas de subsets, isso é essencial.

Para cada número no array, temos exatamente duas escolhas: ou incluimos ele no subset atual, ou não incluimos. Essas duas chamadas recursivas exploram essas duas possibilidades.

1. A primeira parte faz isso:

```go
// decisão para incluir o número em i
subset = append(subset, nums[i])
dfs(i + 1)
```

Basicamente: vamos incluir o número no índice i no subset, e depois chamamos a recursão para processar o próximo índice.

**Essa chamada recursiva vai descer por toda a árvore de decisões**

2. Desfazendo a decisão

```go
subset = subset[:len(subset)-1]   // DESFAZ: remove o número que acabou de adicionar
dfs(i + 1)                         // EXPLORA: vê o que acontece se não incluir
```

**=>** Nessa segunda condição é importante entender que, só irá ocorrer após a chamada recursiva retornar (stack frames forem executados), e isso só acontece depois da condição de parada / base case, ou seja, i >= len(nums).

- Por que desfazemos?

    É um princípio do backtracking. Quando termina e retorna de uma chamada recursiva, precisamos voltar ao que era antes daquela chamada começar. Todas as modificações que aconteceram dentro daquela árvore de chamadas já foram desfeitas pelas suas próprias operações de backtracking.

Como não chegamos em base case, vamos remover o número que acabamos de aduicionar (desfazer a decisão), e chamar o backtracking novamente.

Essa segunda chamada vai explorar todos os subsets que não contém o número em i.

A árvore de decisão ficará assim:

```
                        []  (i=0, início)
                       /  \
                      /    \
           (inclui 1)/      \(não inclui 1)
                    /        \
                  [1]         []  (i=1)
                 /  \        /  \
      (inclui 2)/    \      /    \(não inclui 2)
              /      \    /      \
           [1,2]    [1]  [2]     []  (i=2)
           /  \     / \  / \     / \
(inclui 3)/    \   /   \/   \   /   \(não inclui 3)
        /      \ /     /\    \ /     \
    [1,2,3] [1,2] [1,3] [1] [2,3] [2] [3] []  (i=3, adiciona ao resultado)
```

**Exemplo com as iterações:**

```go

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		fmt.Println("subset", subset)
		fmt.Println("res", res)
		if i >= len(nums) {
			combo := make([]int, len(subset))
			copy(combo, subset)
			res = append(res, combo)
			return
		}

		// decisão para incluir o número em i
		subset = append(subset, nums[i])
		dfs(i + 1)

		fmt.Println("unstacking")

		// decisão para não incluir o numero na i
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)

	return res
}
```

**Logs:**

```
subset []
res []
subset [1]
res []
subset [1 2]
res []
subset [1 2 3]
res []
unstacking
subset [1 2]
res [[1 2 3]]
unstacking
subset [1]
res [[1 2 3] [1 2]]
subset [1 3]
res [[1 2 3] [1 2]]
unstacking
subset [1]
res [[1 2 3] [1 2] [1 3]]
unstacking
subset []
res [[1 2 3] [1 2] [1 3] [1]]
subset [2]
res [[1 2 3] [1 2] [1 3] [1]]
subset [2 3]
res [[1 2 3] [1 2] [1 3] [1]]
unstacking
subset [2]
res [[1 2 3] [1 2] [1 3] [1] [2 3]]
unstacking
subset []
res [[1 2 3] [1 2] [1 3] [1] [2 3] [2]]
subset [3]
res [[1 2 3] [1 2] [1 3] [1] [2 3] [2]]
unstacking
subset []
res [[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3]]
[[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3] []]
```