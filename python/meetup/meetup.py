#!/usr/bin/env python
# -*- coding: utf-8 -*-

from datetime import date
import calendar

def meetup_day(year, month, day_of_week, which):
    if which == 'last':
        return date(year, month, calendar.monthrange(year, month)[1])
    elif which == '1st':
        return find_date(year, month, 1, day_of_week)
    elif which == '2nd':
        return find_date(year, month, 8, day_of_week)
    elif which == 'teenth':
        return find_date(year, month, 13, day_of_week)
    elif which == '3rd':
        return find_date(year, month, 15, day_of_week)
    elif which == '4th':
        return find_date(year, month, 22, day_of_week)
    elif which == '5th':
        return find_date(year, month, 28, day_of_week)


def find_date(year, month, day, day_of_week):
    weekdays = {
            "Monday": 0,
            "Tuesday": 1,
            "Wednesday": 2,
            "Thursday": 3,
            "Friday": 4,
            "Saturday": 5,
            "Sunday": 6
            }
    while weekdays[day_of_week] != calendar.weekday(year, month, day):
        day = day + 1

    return date(year, month, day)
