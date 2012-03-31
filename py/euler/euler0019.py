#! /usr/bin/python
# D:{SADASANT;}
import euler
def euler0019(d,m,y): # d must be 31, m must be 12, y must be 2000 in order to solve the problem
  """ How many Sundays fell on the first of the month during the twentieth century? """
  D,M,Y = 1,1,1901
  m30d = [4,6,9,11]
  m31d = [1,3,5,7,8,10,12]
  week = {1:'monday',2:'tuesday',3:'wednesday',4:'thursday',5:'friday',6:'saturday',7:'sunday'}
  wekc = 2
  if m > 12: return 'month out of range'
  if m == 2 and d > 28: return 'day out of range'
  elif m in m30d and d > 30: return 'day out of range'
  elif m in m31d and d > 31: return 'day out of range'
  print D,M,Y,week[wekc]
  answercounter = 0
  while (D,M,Y) != (d,m,y):
    if D is 1 and wekc is 7: answercounter += 1 
    if M is 2 and D is 28 and Y%4 and (not Y%100 or Y%400): M,D = M+1,0
    elif M is 2 and D is 29: M,D = M+1,0
    elif M in m30d and D is 30: M,D = M+1,0
    elif M in m31d and D is 31: M,D = M+1,0
    D,wekc = D+1,wekc+1
    if wekc > 7: wekc = 1
    if M is 13: Y,M,D = Y+1,1,1
    print D,M,Y,week[wekc]
  return answercounter

print euler0019(31,12,2000)

