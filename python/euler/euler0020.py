#! /usr/bin/python
# D:{SADASANT;}
import euler
from math import factorial
def euler0020(n): # n must 100 in order to solve the problem
  """ Find the sum of digits in 100! """
  fomg = factorial(n)
  fomgIsNowStr = str(fomg)
  theSumOfTheSillyNumbers = 0
  for x in fomgIsNowStr: theSumOfTheSillyNumbers+=int(x)
  return theSumOfTheSillyNumbers
  
print euler0020(100)
