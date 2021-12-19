sieve [] = []            
sieve (x:xs) = x : sieve (filter (\a -> not $ a `mod` x == 0) xs)

n = 100
main = print $ sieve [2..n]
