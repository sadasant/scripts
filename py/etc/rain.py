#!/usr/bin/env python
# By: D:{SADASANT;}
# Special thanks to Ray Myers (matrix.py)
# License: Whatever. Public domain sounds good. ;)

#  RrRRRRr     A    IiiIiiI nNn   Nn
#  rRr  RRr   aAa     iIi   nNNn  Nn
#  RrR rRR*  aA Aa    iIi   nN Nn Nn
#  rRr'*Rr  aAaAaAa   iIi   nN  N Nn
#  RrR  RRr aA   Aa   iIi   nN  nNNn
#  rRr  RRr aA   Aa IiiIiiI nN   nNn

import curses, os, random, string, sys, thread, traceback

################################################################################

if len(sys.argv)<3: 
  sys.exit("""
    By: D:{SADASANT;}
    Special thanks to Ray Myers (matrix.py)
    License: Whatever. Public domain sounds good. ;)

    RrRRRRr     A    IiiIiiI nNn   Nn
    rRr  RRr   aAa     iIi   nNNn  Nn
    RrR rRR*  aA Aa    iIi   nN Nn Nn
    rRr'*Rr  aAaAaAa   iIi   nN  N Nn
    RrR  RRr aA   Aa   iIi   nN  nNNn
    rRr  RRr aA   Aa IiiIiiI nN   nNn
  
    Hello :]
    
    try:
      python rain.py 50 50 15
  
    - The first value is the number of drops that will appear at the top every turn.
    - The second value is the number of ms between each step.
    - The third value is how long this program will wait for fadin a character.
    
    Enjoy!
  """)
  
rate  = int(str(sys.argv[1]))    # The number of drops that appear at the top every turn
delay = int(str(sys.argv[2]))   # number of ms in between each step
fade  = int(str(sys.argv[3]))   # how long until a character fades back to black
if len(sys.argv)>4: calm = int(str(sys.argv[4]))   # how long until a character fades back to black
else: calm = 0
water = 'cloud'
heat = "ps ax"
lapse = 0
if calm:
  heat = "ls" # "ps rax" # "ls -Rat -1"
  lapse = delay

################################################################################

# A graphics abstraction:
class X:
  setup = False
  def __init__(self):
    global green, white
    # Init curses itself:
    self.w = curses.initscr() ; X.setup = True
    # Set all of our modes:
    curses.noecho() ; curses.cbreak()
    curses.start_color() ; curses.curs_set(0)
    # Set up the color pairs:
    curses.init_pair(1, curses.COLOR_GREEN, curses.COLOR_BLACK)
    curses.init_pair(2, curses.COLOR_WHITE, curses.COLOR_BLACK)
    self.green = curses.color_pair(1)
    self.white = curses.color_pair(2)
    # Ok, now for a grid:
    self.data = []
    for i in xrange(curses.LINES):
      self.data.append( [ (0, ord(' ')) ] * curses.COLS )

  def set(self, (x,y), char):
    if x < 0 or x >= curses.COLS or y < 0 or y >= curses.LINES: return
    self.data[y][x] = (fade, char)

  def step(self):
    for i in xrange(len(self.data)):
      for j in xrange(len(self.data[i])):
        newcount, char = self.data[i][j]
        if newcount > 0:
          newcount -= 1
          self.data[i][j] = (newcount, char)

  def draw(self):
    for i in xrange(len(self.data)):
      for j in xrange(len(self.data[i])):
        count, char = self.data[i][j]
        char = ord(' ') if count<=0 else char
        color = self.green if count<fade else self.white
        try: self.w.addch(i, j, char, color)
        except: pass

# Each letter on the screen:
class drop:
  def __init__(self, x,lst):
    self.x, self.y = x, -1
    self.alive = True
    self.char = ''
    self.lst = lst
    self.newchar()
  
  def newchar(self):
    if self.lst == '': return
    self.char = ord(self.lst[:1])
    self.lst = self.lst[1:]
  
  def move(self):
    # Are we dead?
    if not self.alive:
      return
    # Did we die?
    if len(self.lst)==0:
      self.alive = False
      return
    # Else, move:
    self.y += 1
  
  def step(self, g):
    if not self.alive:
      return
    self.move()
    self.newchar()
    g.set((self.x,self.y), self.char)
  def dead(self):
    fade = len(self.lst)

def thunder():
  global rain
  global calm
  global lapse
  if not len(rain):
    if lapse > 0:
      lapse-=1
      return ''
    else:
      lapse = delay*calm
      thread.start_new_thread(storm,tuple())
      return ''
  return rain.pop(0)

def storm():
  global condense
  global lapse
  global calm
  global rain
  global rate
  pop = os.popen(heat)
  read = (pop.read()).split('\n')
  read.pop(0)
  writes = []
  try: #fix this asshole
    for I in range(len(read)):
      do = read[I].split(' ')
      write = '' ;
      for i in range(0,len(do)):
        if do[i] != '' and ':' not in do[i] and ( i<3 or i>11 ):
          write += '+'+do[i]
      writes.append(write)
    for I in range(len(rain)):
      if rain[I] and rain[I] == writes[I]:
        writes[I] = ''
  except: ''
  rain = writes
  return

def main():
  global rain
  rain = []
  g = X()
  thunder()
  tears = [] ; starts = []
  for a in range(0,rate):
    tears.append\
    (drop(random.randint(0, curses.COLS),''));
    starts.append(random.randint(5,fade))

  while True:
    g.step()
    for I in xrange(len(tears)):
      if starts[I] > 0:
        lastA, lastS = tears[I], starts[I]
        tears.pop(I)
        starts.pop(I)
        tears = [lastA]+tears
        starts = [lastS-1]+starts
      else:
        tears[I].step(g)
        if not tears[I].alive:
          tears[I].dead()
          tears[I]=drop(random.randint(0, curses.COLS),thunder())
          starts[I] = random.randint(5,fade)
    g.draw()
    g.w.refresh()
    curses.napms(delay)

if __name__ == "__main__":
  try:
    main()
  except:
    if X.setup:
      curses.nocbreak()
      curses.echo()
      curses.endwin()
    traceback.print_exc()
