#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0024(n): # n must be 123456789 in order to solve the problem
  """ What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9? """
  lexperm = euler.LEXPERM(n,1)
  print '   '
  print ' looping long stocking '
  print '   '
  c = 0
  answer = []
  for x in lexperm:
    c += 1
    if c == 1000000: answer = x
    print x
  print 'Feel the waves of the tao.'
  return answer
  
print euler0024('0123456789')
