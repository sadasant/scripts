#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0022(n): # n must nothing in order to solve the problem
  """ What is the total of all the name scores in the file of first names? """
  LTRS = {'A':1,'B':2,'C':3,'D':4,'E':5,'F':6,'G':7,'H':8,'I':9,'J':10,'K':11,'L':12,'M':13,\
  'N':14,'O':15,'P':16,'Q':17,'R':18,'S':19,'T':20,'U':21,'V':22,'W':23,'X':24,'Y':25,'Z':26}
  FILE = open('names.txt')
  FILE = eval((FILE.readlines()).pop())
  FILE = sorted(FILE)
  SUM = 0
  for x in range(len(FILE)):
    value = 0
    print x,FILE[x]
    for xx in FILE[x]:
      value += LTRS[xx]
      print LTRS[xx],value
    value *= x+1
    print value,x+1
    SUM += value
#    if FILE[x] == "COLIN": return 'colin'
  return SUM
  
print euler0022(100)
