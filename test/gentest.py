#!/usr/bin/env python3

from random import randint
import sys

levels = 10
try:
    index = sys.argv[1]
except (IndexError):
    index = ""

f = open("generated_test" + index + ".txt", "w")

def rand_nums():
    list = []
    for i in range(0, randint(1, 3)):
        list.append(randint(1, 100000))
    return list

def gen(lvl):
    global levels
    global f

    if levels != lvl:
        s = "-" * lvl
        nums = rand_nums()
        for n in nums:
            f.write(s + " " + str(n) + "\n")
            gen(lvl + 1)


if __name__ == "__main__":
    f.write(" 0\n")
    gen(1)


f.close()
