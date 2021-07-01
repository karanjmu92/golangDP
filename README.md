# Dynamic Programming

***

Dynamic Programming is mainly an optimization over plain recursion. Wherever we see a recursive solution that has repeated calls for same inputs, we can optimize it using Dynamic Programming. The idea is to simply store the results of subproblems, so that we do not have to re-compute them when needed later. This simple optimization reduces time complexities from exponential to polynomial

## Dynamic Programming Recipe

**notice any overlapping subproblems**

**decide what is the trivially smallest input**

**think recursively to use memoization**

**think iteratively to use tabulation**

**draw a strategy first**



### Dynamic Prgramming Problems using Go language.

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

 6. **CanConstruct**
    
  Explanation : Write a function `canConstruct(target, wordBank)` that accepts a target string and an array of strings.
                The function should return a boolean indicating whether or not the `target` can be constructed by concatenating elements of the wordBank array.
  
                You may reuse elements of wordBank array as many times as needed.

 7. **CountConstruct**
  
  Explanation : Write a function `countConstruct(target, wordBank)` that accepts a target string and an array of strings.
                The function should return the number of ways the `target` can be constructed by concatenating elements of the wordBank array.

                You may reuse elements of wordBank array as many times as needed.

 8. **AllConstruct**

  Explanation : Write a function `countConstruct(target, wordBank)` that accepts a target string and an array of strings.
                The function should return a 2D array containing all of the ways that the `target` can be constructed by concatenating elements of the wordBank array.
                Each element of the 2D array should represent one combinatin that constructs the `target`.

                You may reuse elements of wordBank array as many times as needed.    
  
    
