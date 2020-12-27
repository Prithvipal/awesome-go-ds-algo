# Efficient Iterative Solution

## Complexity

**Time Complexity:** Theta(log n) <br/>
**Auxiliary space Space** Theta(1)

## Facts
- Every number can be written as sum of powers of 2 (set bits in binary representation)
- We can traverse through all bits of a number (from LSB to MSB) in O(log n) times.

```
3^10 = 3^8 * 3^2          | 10: 1010 = take 1's position as power from left: 3^2 
                          | and 3^8
                          |
3^19 = 3^16 * 3^2 * 3^1   | 10011
```
