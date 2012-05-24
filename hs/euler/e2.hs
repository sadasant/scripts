-- Project Euler, Problem 2
-- by Daniel Rodríguez <http://sadasant.com/>
-- 23/05/2012

-- We need this to getArgs
import System.Environment

-- Part 1: Fibonacci

-- We start defining the Canonical zipWith Implementation
-- of the Fibonacci secuence, like the ones we can find
-- commonly on the internet xD but lets explain it a bit.

-- zip is a function, but as haskell behaves with
-- list comprehension, zip literaly *iterates* over
-- two lists, *while* mixing them into touples.

-- > zip [1,2] [1,2]
-- > zip [(1,1),(2,2)]

-- Using zipWith, once that both elements are zipped,
-- they're passed as arguments to the given action

-- > zipWith (+) [1,2] [1,2]
-- > [(1+1),(2,2)]
-- > [2, 4]

-- As lists comprehensions iterates both lists,
-- in the case of the fib function, they iterate
-- progressively through the same construction of
-- fibonacci elements.

-- I decided to start with 1 : 2 because I didn't
-- care about 0 and 1.
-- The function start having only [1,2]
-- zipWith (+) [1,2] (tail [1,2]), gives us [(1+2)]
-- appended to the original [1,2], it's [1,2,3],
-- now the *list comprehension* continues,
-- zipWith (+) [2,3] (tail [2,3]), gives us [(2+3)]
-- now we have [1,2,3,5]
-- and so on...

fib = 1 : 2 : zipWith (+) fib (tail fib)

-- Now, we can takeWhile that infinite fib sequence.
-- takeWhile stops the excecution of fib if the
-- condition is not satisfied.

-- takeWhile (<=n) fib will give us all the fib numbers
-- chopping it only with whose values do not exceed n
-- Some lines below you'll see how I'm using takeWhile...

-- Also, after chopping the list, we must loop through it
-- there are many ways, one way is to filter those that
-- aren't even, by dropWhile or filter, or others,
-- and then, sum the whole list up, which would require
-- looping again over the list.

-- On the other way, with fold we can effectively
-- pick those who are pair and sum them.

-- Part 2: Foldl (aka filter, for JavaScripters)

-- So we create a function that we will use inside
-- foldl, this function will check if the new value b
-- is even, and if it is, it will sum the old value a
-- with the new value b; else, it will return the old
-- value a.

sumIfEven a b = if even b then a + b else a


-- Using it inside foldl, we pass a first parammeter 0
-- and as second parammeter the chopped fibonacci,
-- foldl will iterate the secuence, adding just
-- the even valued elements.

e2 n = foldl (sumIfEven) 0 (takeWhile (<=n) fib)

-- and that's it \m/

-- Part 3: Let's use it.

-- First, we get the args (just the first)
-- then, we change it to Int with (read a + 0)
-- and we run e2 with the given value.

main = do
  (a:_) <- getArgs
  print (e2 (read a + 0))

-- Test it out!
-- $ runhaskell e2.hs 4000000
