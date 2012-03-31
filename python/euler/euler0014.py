#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0014(n): # n must be 1000000 in order to solve the problem
  """ Find the longest sequence using a starting number under one million. """
  c,N,ln = 0,0,[]
  while c < n:
    ch = []
    c += 1
    N = c
    print N
    while N > 1:
      ch.append(N)
      if not N%2: N /= 2
      else: N = (3*N)+1
      if N == 1: break
    if len(ch) > len(ln): ln = ch
  return ln
      
print euler0014(1000000)
