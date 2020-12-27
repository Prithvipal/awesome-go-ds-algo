# Efficient Recursive Solution

**Time Complexity:** Theta(log n) <br/>
**Auxiliary space Space** Theta(log n)

```
                |-- if n % 2 == 0
                |      power(x, n/2) * power(x * n/2)
                |
Power(x,n) -----|
                |
                |
                |-- Else
                        power(x,n-1) * x
```