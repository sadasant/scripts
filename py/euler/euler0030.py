#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0030(n,v=0): # n must be 5 in order to solve the problem
  """ Find the sum of all the numbers that can be written as the sum of fifth powers of their digits. """
  R = 0
  for x in xrange(10,10**(n+1)):
    if v:print x,
    xtr = str(x)
    y = 0
    for xx in xtr:
      if v:print str(xx)+'^'+str(n),
      xxn = int(xx)**n
      if v:print str(xxn),
      y += xxn
    if v:print '=== '+str(y)
    if x == y:
      print x
      R += x
  return 'the answer is ',R

print euler0030(5,0)
