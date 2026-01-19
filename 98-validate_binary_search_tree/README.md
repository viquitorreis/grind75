# Validando Árvore binária

NÃO PODE OLHAR UM LADO OU OUTRO, tipo assim:

```go
if root.Left != nil {
    return validateRec(root.Left, min, root.Val)
} else {
    return validateRec(root.Right, root.Val, max)
}
```

**TEM QUE OLHAR DOS DOIS LADOS SIMULTANEAMENTE PARA PODER CHECAR TODOS OS NÓS FILHOS A PARTIR DO NÓ ATUAL!!!**.

Fica assim:

```go
return validateRec(root.Left, min, root.Val) && validateRec(root.Right, root.Val, max)
```