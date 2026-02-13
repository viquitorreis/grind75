## modelo mental para esse problema

A mediana tem uma propriedade importante: ela separa os dados em dois grupos iguais:

- Tudo abaixo da mediana e tudo acima

Não precisamos que cada grupo esteja totalmente ordenado internamente, só precisamos saber:

- O **maior elemento do grupo inferior** 
- O **menor elemento do grupo superior**

Um exemplo visual, pensa que temos dois grupos, separados por uma linha do meio:

```
grupo esquerdo: [1, 2, 3]  |  grupo direito: [4, 5, 6]
                        mediana
```

Para calcular a mediana, só precisamos **olhar para as bordas que estão viradas para o centro** que no caso são `3` e `4`. Não importa a ordem interna de cada grupo.

Agora precisamos pensar na estrutura de dados para isso, qual estrutura nos dá acesso ao maior elemento de um grupo de forma instantanea? Uma **max heap**. E acesso ao menor elemento de forma instantanea? Uma **min heap**.

---

**Entendendo o balanceamento.**

Um trabalho importante é garantir *duas invariantes a cada inserção.*.

- Primeira: os grupos precisam estar ordenados entre si. Todo elemento da esquerda deve ser menor ou igual a todo elemento da direita.
- Segunda: os tamanhos devem ser iguais ou ser diferentes em no máximo 1, para que a mediana seja o topo de um ou a média dos dois topos.

---

### Lógica de inserção

Quando um elemento novo chega, precisamos **classificar** ele:

- Ele pertente à metade esquerda ou a direita?

O que vai dividir os dois lados é o **topo do MaxHeap**, que é o maior elemento da metade esquerda.

- Então faz sentido colocar o número lá primeiro e deixar o heap se reorganizar.
- Depois disso verificamos se ele realmente pertence à esquerda ou se ele é grande demais e precisa ir para a direita. (`maxHeap.top > minHeap.top`)

**Passo 1: Corrigir a ordem entre os heaps**

Depois de inserir no maxHeap, verificamos `maxHeap.top > minHeap.top`.
    - Sim: o maior elemento da esquerda é maior que o menor da direita, cruzamos a fronteira e a invariante de ordem foi quebrada. Movemos o topo do maxHeap para o minHeap

```
Antes:  maxH=[7,3,1]  minH=[5,6]    <- 7 > 5, inválido
Depois: maxH=[3,1]    minH=[5,6,7]  <- topo esquerdo 3 <= topo direito 5
```

**Passo 2: Corrigir o balanço de tamanhos**

Depois de fazer isso, os heaps estão ordenados corretamente, mas podem estar desbalanceados. Se o maxHeap tem 2 ou mais elementos a mais que o minHeap, a mediana está errada, não está mais nos topos -> vai estar dentro do heap maior.
    Então movemos o excesso, o topo do maxHeap vai para o minHeap.

A convenção é deixar o maxHeap ter no máximo 1 elemento a mais.

Então quando o total for ímpar, a mediana está no topo do maxHeap, e quando for par, a mediana é a média dos dois topos.

Exemplo:

```
addNum(1):
  maxH=[1]        minH=[]
  ordem ok, balanço ok → median = 1.0

addNum(7):
  insere no maxH -> maxH=[7,1]
  topo maxH(7) > topo minH(∞)? não (minH vazio)
  balanço: maxH.len(2) > minH.len(0)+1 -> move 7 para minH
  maxH=[1]        minH=[7]    -> median = (1+7)/2 = 4.0

addNum(3):
  insere no maxH -> maxH=[3,1]
  topo maxH(3) > topo minH(7)? não
  balanço ok (2 vs 1, diferença de 1 é permitida)
  maxH=[3,1]      minH=[7]    -> median = 3.0

addNum(5):
  insere no maxH -> maxH=[5,3,1]
  topo maxH(5) > topo minH(7)? não
  balanço: maxH.len(3) > minH.len(1)+1 -> move 5 para minH
  maxH=[3,1]      minH=[5,7]  -> median = (3+5)/2 = 4.0
```