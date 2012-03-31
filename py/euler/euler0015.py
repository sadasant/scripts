#! /usr/bin/python
# D:{SADASANT;}
import euler
from math import factorial
def euler0015(n): # n must be 20 in order to solve the problem
  """ Starting in the top left corner in a 20 by 20 grid, how many routes are there to the bottom right corner? """
  print factorial(2*n)/(factorial(n)*factorial(n))
  return ''
  
print euler0015(20)
#print euler0015(4)
