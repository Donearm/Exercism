#!/usr/bin/env python2
# -*- coding: utf-8 -*-
import datetime

def add_gigasecond(date):
    return date + datetime.timedelta(seconds=10**9)
