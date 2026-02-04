# Word Search

Esse algoritmo é um pouco complicado num primeiro momento, mas resolvemos ele da seguinte forma:

- DFS
- Backtracking

Basicamente, percorremos cada elemento do grafo e para cada um, precisamos de DFS para percorrer a partir do vértice, todos os vizinhos dele. Temos algumas condições de parada:

**Parada falsa:**

- Out of bounds: se o valor da row ou coluna não existir no grafo, return false e vai para o pŕoximo elemento do grafo
- Se o par já foi visitado anteriormente (podemos fazer isso com um map, ou podemos mudar a letra atual no grafo para evitar revisitar)
- Se o byte não existe na string.

**Parada verdadeira:**

- Indice atual é igual ao índice da palavra -> significa que percorremos todos os elementos e nenhum retornou false.

**Condição de backtrack**

AtrFez DFS a partir de um vértice para todas directions possívels, depois disso re-colocamos o valor original do vértice (backtrack).