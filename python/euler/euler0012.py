#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0012(n): # n must be 500 in order to solve the problem
  """ What is the value of the first triangle number to have over five hundred divisors? """
  a,c,M = 0,0,0
  while True:
    a += 1
    c += a
    f = len(euler.mult(c))
    if f > M:
      M = f
      print a,c,M
    if M > n:
      return a,c,M
      
      

print euler0012(500)
76576500

