#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0005(n): # n must be 20 in order to solve the problem
  """ What is the smallest number divisible by each of the numbers 1 to 20? """
  R = 1
  for x in euler.powD(range(1,n),1,0,1,1): R *= x
  return R
  

print euler0005(20)
