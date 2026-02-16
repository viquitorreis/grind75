# Word letter

## Modelo mental

**O que procuramos**

Temos uma palavra inicial e final, e queremos encontrar o **caminho mais curto** entre elas onde cada passo troca exatamente uma letra.

"Mais curto" + "cada passo válido", isso tá gritando BFS. Pois BFS garante que você encontra o destino pelo menor número de passos possível.

O grafo existe, mas não precisa construir explicitamente antes de começar. Cada palavra é um nó, e dois nós têm uma aresta entre eles vão se diferir em exatamente uma letra.

Mas como encontramos os vizinho de uma palavra de forma eficiente?

**O problema de encontrar vizinhos**

Algumas soluções possíveis:

1. Podemos comparar cada palavra com todas as outras para ver se diferem em 1 letra. Isso funciona, mas é O(n²) para construir o grafo. Para wordLists grandes é muito lento.

2. A segunda solução: ao invés de comparar palavras diretamente, usamos **padrões com wildcard**.
    Para a palavra `"hot"`, geramos 3 padrões: `"*ot"`, `"h*t"`, `"ho*"`.
    Qualquer palavra que tem esses padrões vai diferir de hot em apenas uma letra.
    Então construímos, um map de padrão -> lista de palavras que encaixam nele.
    Para encontrar os vizinhos de uma palavra, vamos gerar seus padrões e consultamos o map - O(L) onde L é o tamanho da palavra, não O(n).

---

Trace com `hit -> cog`

Usando `wordList = ["hot","dot","dog","lot","log","cog"]`:

1. Começamos com "hit" na fila. Os padrões de "hit" são `"*it"`, `"h*t`, `"hi*`. Consultando o map `"h*t"` encontra `"hot"`. Então `"hot"` é vizinho de `"hit"`, então adicionamos à fila com distância 2.

2. Agora proecssamos "`hot`". Os padrões "`hot`" são "`*ot`", "`h*t`", "`ho*`". O padrão "`*ot`" encontra "`dot`" e "`lot`". Adiciona esses com distância 3.

Continua assim até chegar em "`cog`", O caminho vai ficar: `hit -> hot -> dot -> dog -> cog` ou `hit -> hot -> lot -> log -> cog`

---

**Sobre o `visited`**

Sem controlar os visitados, processaríamos "`hot`" muitas vezes, depois de vir de "`hit`". O BFS ia explodir em loop.