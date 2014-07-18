main = do
  print $ foldr1 (lcm) [1..20]
