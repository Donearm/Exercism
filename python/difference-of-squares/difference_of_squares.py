#!/usr/bin/env python
# -*- coding: utf-8 -*-

def square_of_sum(num):
    tot = 0
    for n in range(1, num + 1):
        tot += n
    return tot*tot

def sum_of_squares(num):
    tot = 0
    for n in range(1, num + 1):
        tot += n*n
    return tot

def difference(num):
    return square_of_sum(num) - sum_of_squares(num)
