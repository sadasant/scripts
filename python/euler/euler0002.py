#! /usr/bin/python
# D:{SADASANT;}
#import euler
def euler0002(n): # n must be 4000000 in order to solve the problem
  """ Sum all the even-valued terms in the Fibonacci sequence which do not exceed four million. """
  x,y,R = 1,0,0
  while R<n:R,y,x=((x%2==0)and[R+x]or[R])[0],x,(x+y)
  return R
  
print euler0002(4000000)
