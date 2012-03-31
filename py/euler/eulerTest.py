#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0024(n): # n must be 123456789 in order to solve the problem
  """ What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9? """
  n = list(n)
  c = 0
  flag1,flag2 = 1,1
  print n,c+1
  while c < 20:
    c += 1
    if not c%2:
      if not c%3:
        chng1 = n[-1]
        n[-1] = n[-3]
        n[-3] = chng1
        if flag2>0:
          chng1 = n[-3]
          n[-3] = n[-4]
          n[-4] = chng1
        else:
          chng1 = n[-2]
          n[-2] = n[-4]
          n[-4] = chng1
        flag2 *= -1
      else:
        chng1 = n[-1]
        n[-1] = n[-2]
        n[-2] = chng1
        if flag1>0:
          chng1 = n[-2]
          n[-2] = n[-3]
          n[-3] = chng1
        else:
          chng1 = n[-1]
          n[-1] = n[-3]
          n[-3] = chng1
        flag1 *= -1
    else:
      chng1 = n[-1]
      n[-1] = n[-2]
      n[-2] = chng1
    print n,c+1
  return ''
  
print euler0024('0123456789')
