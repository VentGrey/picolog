#!/usr/bin/python3

from typing import Match
import matplotlib.pyplot as plt
import re

def count_log_levels(log_file_path: str) -> None:
    log_level_count: dict[str, int] = {'INFO': 0, 'DEBUG': 0, 'WARNING': 0, 'ERROR': 0, 'OK': 0}
    with open(log_file_path, 'r') as file:
        for line in file:
            match: Match[str] | None = re.search(r'\[(\w+)\]', line)
            if match:
                log_level = match.group(1)
                if log_level in log_level_count:
                    log_level_count[log_level] += 1
    plt.bar(log_level_count.keys(), log_level_count.values(), color=['blue', 'green', 'yellow', 'red', 'purple'])
    plt.xlabel('Log Level')
    plt.ylabel('Occurrences')
    plt.title('Count of Each Log Level')
    plt.show()

if __name__ == '__main__':
    log_file_path: str = 'fake_logs.txt'
    count_log_levels(log_file_path)
