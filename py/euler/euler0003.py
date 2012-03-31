#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0003(n): # n must be 600851475143 in order to solve the problem
  """ Find the largest prime factor of a composite number. """
  return ([x for x in euler.mult(n) if euler.prime(x)])
  
# print euler0003(13195)
print euler0003(600851475143)
