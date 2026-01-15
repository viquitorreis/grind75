# Maximum Depth of Binary Tree

## A Analogia do Prédio

Imagina que cada nó da árvore é um andar de um prédio e você quer descobrir quantos andares o prédio tem no total. Você está no térreo (a raiz) e precisa descobrir qual é o andar mais alto que você consegue alcançar.

Quando você está em um andar qualquer, você olha para cima e vê que existem duas escadas:
    - uma indo para a esquerda
    - outra para a direita

Cada escada leva para um conjunto de andares acima de você. Você faz duas perguntas:

1. Primeira pergunta: se eu subir pela escada da esquerda, quantos andares acima de mim eu consigo alcançar? Isso é o que ```dfs(root.Left)``` responde.

2. Segunda pergunta: se eu subir pela escada da direita, quantos andares acima de mim eu consigo alcançar? Isso é o que ```dfs(root.Right)``` responde.

Agora vem a parte crucial.

Quando você recebe as respostas, você adiciona mais um a cada uma delas. Por quê? Porque você precisa contar o andar onde você está atualmente.

Se a escada da esquerda te leva a dois andares acima de você, então do térreo até lá são três andares no total (o seu andar atual mais os dois de cima).
Por fim, você retorna o máximo entre os dois lados porque a profundidade do prédio é determinada pelo lado mais alto.