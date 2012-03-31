#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0006(n): # n must be 100 in order to solve the problem
  """ What is the difference between the sum of the squares and the square of the sums? """
  t0,t1 = 0,0
  for x in range(1,n+1):
    t0,t1 = (t0+x),(t1+x**2)    
  t0 = t0**2
  return t0-t1

print euler0006(100)
