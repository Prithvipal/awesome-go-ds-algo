# Efficient Solution

`
                |-- if n % 2 == 0
                |      power(x, n/2) * power(x * n/2)
                |
Power(x,n) -----|
                |
                |
                |-- Else
                        power(x,n-1) * x
`