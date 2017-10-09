#!/usr/bin/env python

from random import randint
import sys

# settings
levels = 12
maxchildren = 6
maxnum = 100000

try:
    index = sys.argv[1]
except (IndexError):
    index = ""

f = open("generated_test" + index + ".txt", "w")

def rand_nums():
    global maxchildren
    global maxnum

    list = []
    for i in range(0, randint(1, maxchildren)):
        list.append(randint(1, maxnum))
    return list

def gen(lvl):
    global levels
    global f

    if levels != lvl:
        s = "-" * lvl
        nums = rand_nums()
        i = 0
        for n in nums:
            n+=1
            f.write(s + " " + str(n) + "\n")
            gen(lvl + 1)

f.write(" 0\n")

gen(1)

f.close()
