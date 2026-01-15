# Balanced Binary Tree

Vamos entender, porque esse código não funciona para calcular a altura:

```go
dfs = func(root *TreeNode) (res int) {
    if root == nil {
        return 0
    }
    if root.Left != nil {
        return dfs(root.Left) + 1   // ← Só explora a esquerda!
    } else {
        return dfs(root.Right) + 1  // ← Ou só explora a direita!
    }
}
```

**O bug:** Está explorando APENAS UM lado da árvore. Se existe filho esquerdo, vai só pela esquerda e ignora completamente a direita. Se não existe filho esquerdo, aí vai pela direita.

Mas a **altura** de um nó deveria ser baseada no **máximo entre as duas subárvores**, não em apenas uma delas.

---

## Vamos Traçar esse Test Case Que Falhou

```
       2
      / \
     1   4
        / \
       3   5
            \
             6
```

Quando chamamos dfs(4) (o nó com valor 4):

- root.Left != nil? Sim (existe o nó 3)
- Então: return dfs(root.Left) + 1
- Esse código explora apenas o nó 3, calcula altura 1, retorna 2
- Ele NUNCA olhou para o lado direito (nó 5 → 6)

A altura real da subárvore com raiz em 4 deveria ser 3 (caminho 4 → 5 → 6), mas esse código retorna 2 porque só olhou para o nó 3.

## DFS Correto

```go
func getHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(getHeight(root.Left), getHeight(root.Right)) + 1
}
```

Aqui **sempre** explora ambos os lados e retorna o máximo. Não tem if/else escolhendo um ou outro, precisamos dos dois!

---

## Por Que Precisa do `&& isBalanced(root.Left) && isBalanced(root.Right)`

Olha essa árvore:

```
      1
     / \
    2   2
   /     \
  3       3
 /         \
4           4
```

Na raiz (1):

- Altura esquerda: 3
- Altura direita: 3
- Diferença: 0 ✓ (parece balanceado na raiz!)

Mas olha o nó esquerdo (2):

- Altura esquerda: 2 (2 → 3 → 4)
- Altura direita: 0
- Diferença: 2 ❌ (NÃO balanceado!)

E o nó direito (2) tem o mesmo problema!
Então mesmo que a raiz pareça balanceada (ambos os lados têm altura 3), as subárvores individuais não estão balanceadas. Por isso você precisa verificar recursivamente todos os nós, não apenas a raiz.