#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0010(n): # n must be 2000000 in order to solve the problem
  """ Calculate the sum of all the primes below two million. """
  R = 0
  for i in range(2,n):
    mult = euler.mult(i)
    if len(mult) < 3:
      R+=i
      print i,mult,R
  return R

print euler0010(2000000)
