#!/usr/bin/env python
# -*- coding: utf-8 -*-

def distance(seq1, seq2):
    result = 0
    for i, c in enumerate(seq1):
        if seq1[i] is not seq2[i]:
            result += 1

    return result

