main = do
  print $ a - b
  where
    set = [1..100]
    a = sum set ^ 2
    b = sum $ map (^2) set
