#!/usr/bin/env python3


import csv
import io


# Show how exceptions propagate in Python at the call to next(it),
# but in go you need to carry them through explicitly in Go.


def load_csv_data(text):
    result = []
    for row in csv.reader(io.StringIO(text)):
        if len(row) != 2:
            raise ValueError('Rows must have two entries')
        point = float(row[0]), float(row[1])
        result.append(point)
    return result


def load_csv_data_generator(text):
    for row in csv.reader(io.StringIO(text)):
        if len(row) != 2:
            raise ValueError('Rows must have two entries')
        yield float(row[0]), float(row[1])


def load_csv_data_streaming(stream):
    for row in csv.reader(stream):
        if len(row) != 2:
            raise ValueError('Rows must have two entries')
        yield float(row[0]), float(row[1])
