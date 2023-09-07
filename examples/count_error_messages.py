#!/usr/bin/python3

from typing import Match
import matplotlib.pyplot as plt
import re
from collections import Counter

def count_error_messages(log_file_path: str) -> None:
    error_messages = []
    with open(log_file_path, 'r') as file:
        for line in file:
            match: Match[str] | None = re.search(r'\[(\w+)\] - .* : (.+?) -', line)
            if match:
                log_level, message = match.groups()
                if log_level == 'ERROR':
                    error_messages.append(message)
    counter = Counter(error_messages)
    plt.bar(counter.keys(), counter.values(), color='red')
    plt.xlabel('Error Message')
    plt.ylabel('Occurrences')
    plt.title('Most Frequent Error Messages')
    plt.xticks(rotation=30, ha='right')
    plt.show()

if __name__ == '__main__':
    log_file_path: str = 'fake_logs.txt'
    count_error_messages(log_file_path)
