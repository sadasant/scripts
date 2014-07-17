main = do
  print $ div n (last m)
  where
    m = takeWhile (<n) $ scanl1 (*) [x | x <- [3,5..], mod n x == 0]
    n = 600851475143

