# Diameter of binary tree

Tendo a seguinte árvore:

```bash
      1
     / \
    2   3
   / \
  4   5
```

O caminho mais longo seria:

```bash
4 -> 2 -> 1 -> 3
```

Logo, as **arestas** desse caminho são:

- 1: 4 -> 2
- 2: 2 -> 1
- 3: 1 -> 3

3 arestas nesse caminho mais longo.

**Importante:** O número de arestas é **SEMPRE** o número de nós - 1 de um caminho.

Se temos 4 nós no caminhos, a quantidade de arestas = 4 - 1 = 3

---

Existem casos, onde esse caminho não vai passar pela raíz da árvore:

```bash
        1
       / 
      2   
     / \
    4   5
       / \
      6   7
         / \
        8   9
```

Nesse caso, o caminho mais longo nem passa pela raíz, ficando:

```bash
9 -> 7 -> 5 -> 2 -> 4
```

Precisamos portanto considerar quando o diametro não passa pela raíz, ou seja, está apenas dentro de uma subárvore.

Para isso, precisamos checar **cada nó** da árvore como um possível "centro" do diâmetro, não só a raíz.

## Pontos importantes da implementação e recursão

### Porque 1 + max(l, r)?

É aqui que a soma acontece, caso contrário jamais ia somar nada.

Como é uma função recursiva, as chamadas vão acumulando / criando stack frames a cada recursão na stack do processo.

```Ao chegar no último nó (folha) a soma irá ficar 1 + max(l, r), que será 1 + max(0,0).```

Todos os outros stack frames da fila da stack, que serão somados subsequentemente, terão os valores que foram calculados nos filhos corretamente.

- Mas porque somar 1? 

Essa soma **representa a aresta entre o nó atual e o seu pai**. Cada nível de recursão vai adicionar uma aresta ao caminho.

Portanto, entendendo em uma árvore pequena:

```bash
    1
   / \
  2   3
 /
4
```

- Stack frames ficarão (ao chegar na folha)

Tendo a soma de arestas: ```1 + max(l, r)```

4 (folha esquerda) -> 1 + max(0, 0) = 1
3 (folha direita) -> 1 + max(0, 0) = 1
2 (pai do 4) -> 1 + max(1, 0) = 2

Ao retornar ao pai, ficará:

```bash
left = 2 (2 arestas)
right = 1 (1 aresta)
```

Portanto diametro = 3, que é o total de arestas.

**O Diametro é sempre left + right num nó especifico**. Para obtermos isso globalmente, atualizamos a cada recursão o valor do diametro (caso seja maior).

Outro ponto chave, é que começamos a contagem e empilhamento dos stack frames a partir dos nós filhos, não podemos contar a root, se não fica errado a soma (vai passar 1).

Portanto:

```go
func diameterOfBinaryTree(root *TreeNode) (diamater int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)

		// aqui atualizamos o diametro se o caminho passando por este nó for maior
		// -- somamos as alturas de ambos os lados...
		diamater = max(diamater, leftHeight+rightHeight)

		// retorna a altura entre esse nó e o pai dele
		// basicamente é a aresta. Precisamos do +1 se não nunca soma (folha = max(0,0))
		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)
	return
}
```