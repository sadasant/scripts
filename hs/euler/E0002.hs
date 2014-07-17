main = do
  print $ sum $ filter even $ takeWhile (< 4000000) fibs
  where
    fibs = 1 : 1 : zipWith (+) fibs (tail fibs)
