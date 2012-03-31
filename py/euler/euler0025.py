#! /usr/bin/python
# D:{SADASANT;}
#import euler
def euler0025(n): # n must be 1000 in order to solve the problem
  """ What is the first term in the Fibonacci sequence to contain 1000 digits? """
  x,y = 1,0
  c = 0
  while len(str(x))<n or len(str(y))<n:
    c += 1
    y,x=x,(x+y)
    print len(str(x))
  print 'feel the touch of the tao... '
  return x,'\n',len(str(x)),'\n',y,'\n',len(str(y)),'','','',c
  
print euler0025(1000)
