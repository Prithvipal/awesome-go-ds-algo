# Efficient Solution

**Time Complexity:** log n <br/>
**Auxilary Space** log n

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