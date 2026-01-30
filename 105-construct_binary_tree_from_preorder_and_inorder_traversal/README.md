Nesse algoritmo, temos uns pontos cruciais:

Inorder visita: esquerda -> root -> direita
Preorder visita: root -> esquerda -> direita

1. Podemos facilmente pegar a root a partir da preorder, é o primeiro elemento.
2. PRECISAMOS do índice da root na inorder, pois:
    - INORDER nos diz qual parte da árvore está à esquerda e qual está à direita da root.
3. Precisamos do **índice da root**, pois tudo antes dela na vai pertencer à subárvore esquerda. Tudo depois pertence à subárvore direita.

Exemplo: preorder=[3,9,20,15,7], inorder=[9,3,15,20,7]

- Root é 3 (preorder[0])
- Busca 3 no inorder: **índice 1**
- inorder[:1] = [9] -> subárvore esquerda
- inorder[2:] = [15,20,7] -> subárvore direita

Logo: tudo que vem antes do índice da root na inorder, vai ser os elementos até o índice da raíz na inorder. Então:

## Achando elementos da esquerda e direita na árvore**

1. Pega índice da raíz (preorder[0])
2. Elementos da esquerda: inorder[:idxRoot]
3. Elementos da direita: inorder[idxroot+1:]

## Como a recursão constrói a árvore

```go
root := &TreeNode {
    Val: 3,
    Left: buildTree(preorder[1:2], inorder[0:1]),    // constrói esquerda com [9], [9]
    Right: buildTree(preorder[2:], inorder[2:])      // constrói direita com [20,15,7], [15,20,7]
}
```

cada chamada recursiva, vai:

1. Pegar a fatia correta da preorder e inorder
2. Identificar a root atual desses elementos (preorder[0])
3. Divide de novo usando o inorder, e continua até o base case (len(preorder) == 0)


