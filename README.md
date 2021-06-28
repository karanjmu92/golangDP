# golangDP
Dynamic Prgramming Problems using Go language.

# Problems

1. **Fibonaci Series**

Explanation : Write a function `fib(n)` that takes in a number as an argument.
              The function should return  the n-th number of the Fibonacci sequence.
   
              The 1st and 2nd number of the sequence is 1.
              To generate the next number of the sequence, we sum the previous two.
   
               n:      1, 2, 3, 4, 5, 6, 7,  8,  9, ....
               fib(n): 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
   
   
2. **gridTraveller**

Explanation : Say that you are a traveller on a 2D grid. You begin in the top-left corner. You may only move down or right.
              In how many ways can you travel to the goal on a grid with dimension m * n?
              Write a function `gridTraveller(m, n)` that calculate this.
              
        
3. **canSum**

Explanation : Write a function `canSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.
              The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers from the array.
              
              You may use an element of the array as many times as needed.
              
              You may assume that all input numbers are nonnegative.
              
4. **howSum**

Explanation : Write a function `howSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.
              The function should return an array containing any combination of elements that add up to exactly the targetSum.
              If there is no combination that adds up to the targetSum, then return nil.
              
              If there are multiple combinations possible, you may return any single one.
              
 5. **BestSum**

  Explanation : Write a function `BestSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.
                The function should return an array containing the shortest combination of elements that add up to exactly the targetSum.
                
                If there is a tie for the shortest combination, you may return any one of the shortest.
