#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0023(n): # n must be nothing in order to solve the problem
  """ Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers. """
  abundants = []
  sumOfAbundants = []
  answer = 0
  for x in range(1,28126):
    mult = (euler.mult(x))[:-1]
    sumM = sum(mult)
    state = ''
    if   sumM >  x: 
      state = 'abundant'
      abundants.append(x)
    elif sumM == x: state = 'perfect'
    elif sumM <  x: state = 'deficient'
    print x,mult,'<>',sumM,state
  for x in abundants:
    print x
    for xx in abundants:
      sumOfAbundants.append(x+xx)
  abundants = ''
  for x in range(1,28126):
    print x
    if x not in sumOfAbundants:
      answer += x
  return answer
  
print euler0023(100)
