#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0017(n): # n must be 1000 in order to solve the problem
  """ How many letters would be needed to write all the numbers in words from 1 to 1000? """
  txt = ''
  for x in range(1,n+1):
    txt += euler.n2str(x)
  print txt
  return len(txt)
  
print euler0017(1000)

