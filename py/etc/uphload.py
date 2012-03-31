# D:{SADASANT;}
# :;,._.,;:;,._.,;:;,._.,;:;,._.,;:;,._.,;:
#     uphload ~ project uploader.  v0.1.3
# ;:"'`'":;:"'`'":;:"'`'":;:"'`'":;:"'`'":;
#~ 

import os
import time
import string

# variables
# variables
# variables
u = "" # user
p = "" # port
h = "" # ip:/dir/
l = "" # local parent of the porject
o = "" # objective name, must be the same locally and remotely
a = "" # sudo, or empty
d = 3 # seconds to count before looping
# variables
# variables
# variables

print """
 :;,._.,;:;,._.,;:;,._.,;:;,._.,;:;,._.,;:
     uphload ~ project uploader.  v0.1.3
 ;:"'`'":;:"'`'":;:"'`'":;:"'`'":;:"'`'":;

 > this python will upload your projects using scp.
 > edit the variables inside this file to your own uphload.
 > with love, we remind you:  THERE ARE NO WARRANTIES
"""

l = "ls --time-style='+%d-%m-%Y %H:%M:%S' -lR "+l+o+"*" # usually you dont do this
slp = d               # i mean, really.
cnt = 0
d = 1
def changed(act,prv):
  who = direct = direct2 = ''
  for _act in act:
    __act = _act.split(' ')
    if (len(__act)==1)and(len(__act[0])>3 or len(__act[0])<1):
      direct = __act[0]
    for _prv in prv:
      __prv = _prv.split(' ')
      if (len(__prv)==1)and(len(__prv[0])>3 or len(__act[0])<1):
        direct2 = __prv[0]
      if (direct2 == direct):
        _direct = (((str)(direct)).split(':'))[0]+"/"
        if ( (__act[-1]==__prv[-1]) and (string.find(__act[-1],".")>=0) \
        and (string.find(who,__act[-1])<0) and (len(__act)>3) \
        and (((__act[-2])!=(__prv[-2])) or ((__act[-3])!=(__prv[-3]))) ):
          global cnt
          """printing"""
          print "\n upload n* "+((str)(cnt))+":"
          who+= """  file: """+_direct+(str)(__act[-1])+\
          "\n  previous change: "+(str)((__prv[:-1]).pop())+\
          " ~ actual change: "+(str)((__act[:-1]).pop())+\
          "\n  D: !!! DONE !!! ??? just check ;)"
          _direct_ = (str)((_direct.split(o))[1])
          scp  = a+" scp -P "+p+" "+_direct+(str)\
          (__act[-1])+" "+u+"@"+h+o+_direct_ 
          run  = os.popen(scp)
          cnt+=1
          break
  return who
while True:
  pop = os.popen(l)
  read = (pop.read()).split('\n')
  if(d==1): print ' loading ...\n ' #print read
  else:
    os.write(1,changed(read,d))
  d = read
  time.sleep(slp)
