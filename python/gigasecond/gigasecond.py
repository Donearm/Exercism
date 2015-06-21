#!/usr/bin/env python2
# -*- coding: utf-8 -*-
import datetime

def add_gigasecond(date):
    a = date + datetime.timedelta(seconds=10**9)
    return a
