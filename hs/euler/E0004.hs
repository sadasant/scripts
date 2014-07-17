main = do
  print $ maximum $ filter isPal [x * y | x <- set, y <- set]
  where
    set = [999,998..1]
    isPal x = do
      let s = show x
      s == reverse s
