# Spiral Matrix

Esse algoritmos tem um ponto, num primeiro momento ele é meio complexo, mas precisamos entender o seguinte:

Precisamos de 4 ponteiros:

1. Esquerda
2. Direita
3. Cima
4. Baixo

Percorre a matriz em espiral de fora pra dentro. Usa 4 bounds que vão fechando a cada volta.

Configuração inicial:

- top = 0 (primeira linha)
- bottom = len(matrix)-1 (última linha)
- left = 0 (primeira coluna)
- right = len(matrix[0])-1 (última coluna)

## Passos para cada iteração

1. Direita: Percorre linha top de left até right, depois top++
2. Baixo: Percorre coluna right de top até bottom, depois right--
3. Esquerda: Percorre linha bottom de right até left (só se top <= bottom), depois bottom--
4. Cima: Percorre coluna left de bottom até top (só se left <= right), depois left++

Por que os checks nos passos 3 e 4?

    Evita processar linha/coluna duas vezes quando matriz tem apenas uma linha ou coluna no centro.

**Condição de parada:** ```top > bottom``` ou ```left > right```

    Porque quando os bounds se cruzam, já processou todos os elementos.