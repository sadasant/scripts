#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0004(n): # n must be 1000 in order to solve the problem
  """ Find the largest palindrome made from the product of two 3-digit numbers. """
  return max([x*xx for x in reversed(xrange(100,n)) for xx in reversed(xrange(100,x)) \
  if str(x*xx)==str(x*xx)[::-1]])
  
print euler0004(1000)

