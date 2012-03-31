#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0026(n): # n must be 1000 in order to solve the problem
  """ Find the value of d < 1000 for which 1/d contains the longest recurring cycle. """
  IDC = ''
  N = 0
  for x in range(1,n):
    if euler.prime(x):
      idc = (((euler.iDivision(x)).partition('('))[-1].partition(')'))[0]
      if len(idc) > len(IDC): IDC,N = idc,x
  return IDC,'r',N
  
print euler0026(1000)
