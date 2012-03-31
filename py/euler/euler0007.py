#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0007(n): # n must be 10001 in order to solve the problem
  """ Find the 10001st prime. """
  P,c = 2,0
  while True:
    if euler.prime(P):
      c+=1
      print P,c
      if c == n: return P,c
    P+=1

print euler0007(10001)
