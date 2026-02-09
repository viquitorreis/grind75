Nesse algoritmo, precisamos armazenar a frequência dos elementos.

- Heap
- Frequency

Vai ser uma **max heap**. Pois precisamos armazenar os elementos mais frequentes antes.

Toda vez que processarmos um elementos, precisamos esperar o cooldown / slots até processar esse elementos novamente.

Exemplo: ```tasks = ['A','A','A','B','B','B'], n = 2```

Nesse exemplo, toda vez que executar A por exemplo, precisamos esperar dois slots antes de executar A de novo.

Para resolver isso com uma max heap fazemos:

1. Criamos uma max heap de Tasks. Cada task precisa ter sua frequencia atrelada.
2. Processamos a heap em um JANELA DE COOLDOWN (i <= n). Ao processar a heap (pop), decrementamos a frequency do elemento.
    a. Se frequency ainda > 0, voltamos para o heap novamnete para processar em outra camada.
    b. Se < 0, não.
3. Ao contar os ciclos mínimos