Problema: Para cada posição i, multiplique tudo exceto nums[i].

Exemplo: nums = [2, 3, 4, 5]
Resultado esperado:

res[0] = 3×4×5 = 60 (tudo exceto 2)
res[1] = 2×4×5 = 40 (tudo exceto 3)
res[2] = 2×3×5 = 30 (tudo exceto 4)
res[3] = 2×3×4 = 24 (tudo exceto 5)

- Loop 1: Constrói tudo a esquerda

```go
left = 1  // começa com 1 (neutro da multiplicação)
res = [0, 0, 0, 0]  // array vazio

// i=0 (nums[0]=2)
res[0] = left  // res[0] = 1 (não tem nada à esquerda de 2)
left = left * nums[0]  // left = 1 * 2 = 2
// Estado: res=[1,0,0,0], left=2

// i=1 (nums[1]=3)
res[1] = left  // res[1] = 2 (à esquerda de 3 tem: 2)
left = left * nums[1]  // left = 2 * 3 = 6
// Estado: res=[1,2,0,0], left=6

// i=2 (nums[2]=4)
res[2] = left  // res[2] = 6 (à esquerda de 4 tem: 2×3=6)
left = left * nums[2]  // left = 6 * 4 = 24
// Estado: res=[1,2,6,0], left=24

// i=3 (nums[3]=5)
res[3] = left  // res[3] = 24 (à esquerda de 5 tem: 2×3×4=24)
left = left * nums[3]  // left = 24 * 5 = 120
// Estado: res=[1,2,6,24], left=120
```

- Loop 2: Multiplica pelos produtos à DIREITA:

```go
right = 1  // começa com 1
// res já tem: [1, 2, 6, 24]

// i=3 (nums[3]=5)
res[3] = res[3] * right  // res[3] = 24 * 1 = 24 (não tem nada à direita)
right = right * nums[3]  // right = 1 * 5 = 5
// Estado: res=[1,2,6,24], right=5

// i=2 (nums[2]=4)
res[2] = res[2] * right  // res[2] = 6 * 5 = 30 (direita tem: 5)
right = right * nums[2]  // right = 5 * 4 = 20
// Estado: res=[1,2,30,24], right=20

// i=1 (nums[1]=3)
res[1] = res[1] * right  // res[1] = 2 * 20 = 40 (direita tem: 4×5=20)
right = right * nums[1]  // right = 20 * 3 = 60
// Estado: res=[1,40,30,24], right=60

// i=0 (nums[0]=2)
res[0] = res[0] * right  // res[0] = 1 * 60 = 60 (direita tem: 3×4×5=60)
right = right * nums[0]  // right = 60 * 2 = 120
// Estado: res=[60,40,30,24], right=120
```
