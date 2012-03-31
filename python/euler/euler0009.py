#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0009(n): # n must be 1000 in order to solve the problem
  """ Find the only Pythagorean triplet, {a, b, c}, for which a + b + c = 1000. """
  l = []
  for x in range(1,n/2):
    l.append(x)
  for xxx in l:
    for xx in l:
      for x in l:
        if x<xx<xxx and x+xx+xxx == n and x**2+xx**2 == xxx**2:
          return x,xx,xxx,x*xx*xxx
  return ''

print euler0009(1000)
