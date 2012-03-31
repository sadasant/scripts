#! /usr/bin/python
##### E U L E R #####
import math

""" variables """
# here be variables
below19 = {0:'zero',1:'one',2:'two',3:'three',4:'four',5:'five',6:'six',7:'seven',8:'eight',9:'nine',10:'ten',11:'eleven',12:'twelve',\
13:'thirteen',14:'fourteen',15:'fifteen',16:'sixteen',17:'seventeen',18:'eighteen',19:'nineteen'}
over19 = {2:'twenty',3:'thirty',4:'forty',5:'fifty',6:'sixty',7:'seventy',8:'eighty',9:'ninety'}
more = {1:'hundred',2:'thousand',4:'million'}

""" functions """
# here be functions
def prime(n): # n:number
  """ is prime? """
  if n > 1 and len(mult(n)) < 3: return True
#  if n > 1 and not ([x for x in xrange(2,int(n/3.0)+1) if n%x==0]): return True
  else: return False

def iDivision(n):
#  print n # hell
  numerator, denominator = 1, n
  fraction = []
  remainders = []
  while True:
    numerator *= 10
    r = numerator % denominator
    q = (numerator - r)/denominator
    if(r == 0):
      fraction.append(str(q))
      break
    if r in remainders:
      for x in range(len(remainders)): 
        if remainders[x] == r and int(fraction[x]) == q:
          fraction.insert(x,"(")
          fraction.append(")")
          return "0."+"".join(fraction)
    remainders.append(r)
    fraction.append(str(q))
    numerator = r
  return "0."+"".join(fraction)

def n2str(n,nr=1,v=0): # n:number. nr: number of recurtions. v:erbosity
  """ pass number, return string """
  global nine,teen,decn,more
  if v: print n,nr
  txt,ttxt = '',''
  if not n: return 'zero'
  nstr = str(n)
  z = len(nstr)
  if z > 3:
    txt += n2str(int(nstr[:-3]),nr+1)+more[nr+1]
    nstr = nstr[-4:]
  else:
    if z > 2:
      ttxt += below19[int(nstr[0])]+more[1]
      if nstr[-2]+nstr[-1] != '00': ttxt += 'and'
      nstr = nstr[1:]
    if int(nstr) < 20:
      if int(nstr): ttxt += below19[int(nstr)]
    else:
      tttxt = ''
      if int(nstr[1]): tttxt += below19[int(nstr[1])]
      ttxt += over19[int(nstr[0])]+tttxt
  return txt+ttxt

def LEXPERM(n,v=0): # n: string of number. v:erbosity
  """ Lexicographic Permutation """
  lexperm = []
  if len(n)<2: return ''
  if len(n)>2:
    n = list(n)
    for x in range(len(n)):
      tmpn = n[:]
      tmpn.pop(x)
      for xx in LEXPERM(tmpn):
        xxx = [n[x]]
        xxx.extend(xx)
        if v: print xxx
        lexperm.append(xxx)
  else:
    lexperm = n,n[::-1]
  return lexperm

def mult(n,p=0,v=0): # n:number. p:prime. v:erbosity.
  """ find the multiples or prime factorize """
  if not p:
    factors = []
    for x in xrange(1, int(math.sqrt(n)+1)):
      if n%x == 0:
        factors.append(x)
        factors.append(n/x)
    return sorted(list(set(factors)))
  o,d,c = n,[1,],2
  while n>1:
    if v: print n,c
    if not n%c:
      d.append(c)
      n = n/c
    else: c+=1
  if v: print n
  return sorted(d)
  
def powD(nl,G=0,L=0,cm=0,ncm=0,v=0):
# nl: number list. G:great, L:low. cm:commons. ncm:noncommons. v:erbosity 
  """ find the lowest|greatest power divisors """
  l = []
  for x in range(len(nl)):
    if v>1: print nl[x]
    fx = mult(nl[x],1)
    c = 1
    l.append([[],[]])
    for xx in range(1,len(fx)):
      if v>1: print fx[xx],fx[xx+1:]
      if fx[xx] in fx[xx+1:]: c += 1
      else:
        l[x][0].append(fx[xx])
        l[x][1].append(c)
        c = 1
  if v: print l
  if cm:
    cm = []
    for x in range(len(l)):
      for xx in range(len(l[x][0])):
        if len([yy for y in range(len(l)) \
        for yy in range(len(l[y][0])) \
        if y != x and l[y][0][yy] == l[x][0][xx]]) == len(l)-1:
          if ([l[x][0][xx],l[x][1][xx]]) not in cm:
            cm.append([l[x][0][xx],l[x][1][xx]])
    if v: print cm
    if not ncm: l = cm
  if ncm:
    ncm = []
    for x in range(len(l)):
      for xx in range(len(l[x][0])):
        if len([yy for y in range(len(l)) \
        for yy in range(len(l[y][0])) \
        if y != x and l[y][0][yy] == l[x][0][xx]]) != len(l)-1:
          if ([l[x][0][xx],l[x][1][xx]]) not in ncm:
            ncm.append([l[x][0][xx],l[x][1][xx]])
    if v: print ncm
    if not cm: l = ncm
  else:
    tl = []
    for i in l:
      for x in range(len(i[0])):
        if [i[0][x],i[1][x]] not in tl: tl.append([i[0][x],i[1][x]]) 
    l = tl
  if cm and ncm:
    l = cm
    l.extend(ncm)
  if G or L:
    if v: print l
    for x in range(len(l)):
      for xx in range(len(l)):
        if l[x][0] == l[xx][0]:
          if G and l[x][1] > l[xx][1]:
            l[xx][0],l[xx][1] = l[x][0],l[x][1]
            l[x][0],l[x][1] = 0,0
          elif L and l[x][1] < l[xx][1]:
            l[xx][0],l[xx][1] = l[x][0],l[x][1]
            l[x][0],l[x][1] = 0,0
    c,R = [],[]
    for x in l:
      if x[0] not in c:
        c.append(x[0])
        R.append(x[0]**x[1])
  else: return l
  return sorted(R)
  
