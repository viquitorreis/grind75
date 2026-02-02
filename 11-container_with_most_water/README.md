para resolvermos esse problema, precisamos entender que:

- A altura do menor elemento é a fonte de verdade.

Portanto, sempre que formos calcular, precisamos considerar a altura do elemento menor para performar o cálculo.

Esse problema resolvemos com **two pointers**. Onde:

- max = área do container máximo
- start = altura de onde iniciou o container (ponteiro)
- end = altura de onde terminou o container (ponteiro)

Sempre que acharmos um start < end, vamos atualizar o ponteiro start para frente.
Sempre que acharmos end < start, vamos atualizar o ponteiro end para trás.

Dentro do loop:

- Achamos width = end - start
- Height = menor elemento (nums em start ou em end)
- Área = width * height -> atualiza max se for maior.