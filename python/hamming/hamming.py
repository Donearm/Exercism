#!/usr/bin/env python
# -*- coding: utf-8 -*-

def distance(seq1, seq2):
    difference = 0
    for a, b in zip(seq1, seq2):
            if a != b:
                difference += 1
    return difference
