# Optimized 3Sum Solution

We will first create a 2D slice to hold our results. After that, we will sort the input array. The entries in the nums slice will be easier to manage with multiple pointers if they are in ascending order. This will make the checks for duplicates and pairs simpler logically.

After that, we begin to loop through our input slice. We loop until i == size-2 so we don’t get an out-of-bounds error. The index i will be our left-most pointer.

We can handle some edge cases in our first if statement. There, we will check if i is equal to 0 or if i is greater than 0 and also if the current nums element does not equal the previous nums element. We check if i is equal to 0 because we must also check for matching values when i is greater than 0. When i is 0, there is no previous element to check.

We then get the low, high, and sum. The low will be the second leftmost pointer, located at the index immediately after i. The high variable will be the far right pointer, located at the last entry in the slice. We will then subtract the element at nums[i] from 0 to get the sum. We will need the other 2 pointers, low and high, to add up to this sum.

Next, we will begin a while loop where we look for combinations of elements that add up to the desired sum. We will set the condition for the loop to run while low is less than high. If the values at nums[low] and nums[high] add up to the desired sum, we add them and nums[i] to our result slice for we have found one of the solutions.

We then enter 2 more loops. These loops are used to bypass duplicates. The first one continues while low < high and also while the element at the current pointer position equals the element at the next pointer position. We move the left pointer forwards until a unique element is found.

We do the same logic to check for duplicates at the rightmost pointer, high. Except instead of moving the pointer forwards, we move it backwards. In other words, we move both pointers inwards. After those loops have ended and we have skipped all duplicates, we will move both pointers inwards.

If the sums at low and high don’t add up to the desired sum, we check if the result is greater than the desired sum. If it is, we move high inwards. The reason we do this is because since we sorted the values at the beginning, moving inward means moving to a smaller value. Otherwise, we assume the calculated sum is less than the desired sum, and we move the low pointer forwards.