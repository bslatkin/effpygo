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


def main():
    data = "1.0,2.5\n3.5,4.1\n"

    # Single load
    try:
        rows = load_csv_data(data)
    except (ValueError, IOError):
        print('Broke reading file')
        raise

    for i, row in enumerate(rows):
        print('Row %d is %r' % (i, row))

    # Streaming
    stream = io.StringIO(data)
    it = load_csv_data_streaming(stream)
    i = 0
    while True:
        try:
            row = next(it)
        except StopIteration:
            break
        except (ValueError, IOError):
            print('Broke on row %d' % i)
            raise

        print('Row %d is %r' % (i, row))
        i += 1


if __name__ == '__main__':
    main()
