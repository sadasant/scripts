#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0016(n,e): # n must be 2 and e must be 1000 in order to solve the problem
  """ What is the sum of the digits of the number 2**1000? """
  N = n**e
  print N
  R = 0
  for x in str(N):
    R+=int(x)
  return R
  
print euler0016(2,1000)

