# work: Do the learning rules.
import cPickle as pickle
import time # i dont really need this.

class ballpy(object):
  """ ballpy """

  dYm = []

  def __init__(self):
    print 'welcome to ballpy'

  def pull(self):
    print " Wellcome to the pull function. \n"
    print " Now you must follow the steps to add stuff to a dictionary. \n"
    print " Ready? ... "
    time.sleep(1)  
    print " SET ... "
    time.sleep(1)  
    print " GO!!! \n"
    pull = False
    while pull is False:
      pull = raw_input("""1) Read actual values.    2) Add new value.
                          3) Remove existing value. 4) Quit\n    """)
      if pull == '1':
        pull1 = False 
        while pull1 is False:
          pull1 = raw_input('1) Print all. 2) Interact. 3) To go to main.')
          if pull1 == '1': 
            # print the main functions
            pull1 = False
          if pull1 == '2': 
            try:
              exec raw_input('Try to do something :3.\n')
            except:
              pass
            pull1 = False
          if pull1 == '3':
            print "Done."
      if pull == '2': 
        pull2 = False
        while pull2 is False: 
          pull2 = raw_input('1) Add a function to a place. 3) Quit\n')
          if pull2 == '1':
            pull2_2_1 = False
            pull2_2 = False
            dym = ''
            pull2_3 = raw_input('Please enter the name of the function:\n')
            dym += "  def " + pull2_3 + "(self" #self
            pull2_2_1 = False
            while pull2_2_1 is False or pull2_2_1 != '':
              pull2_2_1 = raw_input('Add a new parameter:\n')
              if pull2_2_1 != '':
                dym += "," + pull2_2_1
            dym += "):\n  "
            pull2_2_2 = raw_input('Please enter the description of the function:\n')
            dym += '  """' + pull2_2_2 + '"""'
            pull2_2_2 = False
            while pull2_2_2 is False or pull2_2_2 != '':
              print dym
              pull2_2_2 = raw_input("""Please enter a line of code
                                       or nothing to finish:\n""")
              dym += "\n    " + pull2_2_2
            self.dYm.append(dym)
            print dym
            d = {}
            exec dym.strip() in d
            setattr(self.__class__, pull2_3, d[pull2_3]) #asign to class
            print "Done =) check 1 to see the results"
          if pull2 == '3':
            pull2 = True
      if pull == '3': 
        unR = True
      if pull == '4':
        print "bye bye :3"
        pull = True
      else:
        pull = False

  def wake(self):
    self.dYm = self.dYm

  def doze(self):
    fail = open("ballpy.py", "r")
    flag = '  ### dYm ###\n'
    d = 0
    content = []
    line = fail.readline()
    while line:
      if line == flag:   
        content.append("\n")
        for d in range(len(self.dYm)):
          content.append(self.dYm[d])
        content.append("\n")
        r = False
      content.append(line)
      line = fail.readline()
    print content
    fail.close()
    fail = open("ballpy.py", "w")
    fail.write("".join(content))
    fail.close()

  def pyLove(self,love):
    """nut : python's love"""
    i = ''
    if love == 'iLoveU': i = 'I WUV U TOO!!! <3'
    return i

  def loling(self, lol):
    """loling"""
    print lol+"... yeah, loling hard."
    

  ### dYm ###


ball = ballpy()
ball.wake()
ball.pull()
ball.doze()
