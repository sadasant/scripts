#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0027(n1,n2): # n1 and n2 must be 1000 in order to solve the problem
  """ Find a quadratic formula that produces the maximum number of primes for consecutive values of n. """
  MAX,R = 0,0
  for a in range(-n1,n1):
    for b in range(-n2,n2):
      n,prime = 0,True
      while prime: prime,n = euler.prime(((n*n)+(a*n)+(b))),n+1
      n-=1
      if n > MAX: MAX,R = n,(a,b,a*b)
#      print MAX,n,a,b
  return 'the answer is ',R
  
print euler0027(1000,1000)
