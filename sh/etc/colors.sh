part_1="signup"
part_2="    Creates a new account at "
part_3="sadasant.com"

# \e doesn't work on OSX
normal="\x1b[0m"

for i in `seq 0 7`; do
  color="\x1b[3"$i"m"
  echo -e "$color$part_1$normal$part_2$color$part_3$normal."
  color="\x1b[3$i;1m"
  echo -e "$color$part_1$normal$part_2$color$part_3$normal."
done
