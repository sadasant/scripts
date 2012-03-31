#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0001(n): # n must be 1000 in order to solve the problem
  """ Add all the natural numbers below one thousand that are multiples of 3 or 5. """
  global m
  def setMult(n): global m ; m = euler.mult(n) ; return 1
  return sum([x for x in range(n) if setMult(x) and (3 in m or 5 in m)])

# print euler0001(10)
print euler0001(1000) # 233168
