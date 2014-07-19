main = do
  print $ sum $ takeWhile (<2000000) primesPT1
  where
    primesPT1       = 2 : sieve primesPT1 [3..]
    sieve (p:ps) xs = let (h,t) = span (< p*p) xs
                      in h ++ sieve ps [x | x<-t, rem x p /= 0]
