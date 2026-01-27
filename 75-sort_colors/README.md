# Sort Colors

É quicksort. Divide And Conquer:

1. Se min < max: chama quicksort recursivo:
    a. pivo = partition
    b. recursivo (min, piv-1)
    c. recursivo (piv+1, max)

2. Partition (nums, min, max)
    a. piv = nums[min]
    b. left, right := min+1, max (min já é usado no pivo...)
    c. Loop enquanto left <= right
        - loop: left procura elementos maiores que pivot
        - loop: right procura elementos menores ou iguais que pivot
        - swap:
            left com right, pois left vai ser maior que o pivo, mas está na parte esquerda
            right é menor que o pivot, mas está na parte direita

