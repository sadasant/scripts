#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0021(n): # n must 10000 in order to solve the problem
  """ Evaluate the sum of all amicable pairs under 10000. """
  if n < 220: return 'the smallest number is over 220. So please... '
  Mtmp = []
  pmtM = []
  for x in range(220,n):
#    print (euler.mult(x))
    Mtmp.append([x,sum((euler.mult(x))[:-1])])
    pmtM.append(Mtmp[-1][::-1])
  R = []
  for x in Mtmp:
    if x!=x[::-1] and x in pmtM and x[::-1] not in R: R.append(x)
  print Mtmp
  print 'and'
  print pmtM
  return R,sum([xx for x in R for xx in x])
  
print euler0021(10000)
#print euler0015(4)
