Nesse problema precisamos usar BFS, pegar os elementos em cada nível.

Para cada nível, vamos pegar o elemento que está mais a direita, ou seja, o **último elemento de cada nível da árvore**.

Então é um BFS, e a cada nível que estivermos processando, vamos chegar se o índice processado está do tamanho do nível atual, se sim fazemos o push do elemento atual na resposta.

