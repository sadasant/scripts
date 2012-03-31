#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0028(n): # n must be 1001 in order to solve the problem
  """ What is the sum of both diagonals in a 1001 by 1001 spiral? """
  n *= n
  R,c,l,r = 1,1,4,2
  while c < n:
    c += r
    R += c
    l -= 1
    if l == 0: l,r = 4,r+2
      
  return 'the answer is ',R
  
print euler0028(1001)
