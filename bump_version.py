#!/usr/bin/env python

import argparse
import re


def _get_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('filename')
    parser.add_argument('bump_type')

    return parser.parse_args()


def _get_current_version(filename):
    with open(filename) as f:
        contents = ''.join(f.readlines())
        match_results = re.search('(?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)', contents).groupdict()
        
        current_major = int(match_results['major'])
        current_minor = int(match_results['minor'])
        current_patch = int(match_results['patch'])
        
        return current_major, current_minor, current_patch


def main():
    args = _get_args()
    
    current_major, current_minor, current_patch = _get_current_version(args.filename)

    bump_type = args.bump_type.lower()

    next_major = current_major + 1 if bump_type == 'major' else current_major
    next_minor = 0 if bump_type == 'major' else current_minor + 1 if bump_type == 'minor' else current_minor
    next_patch = 0 if bump_type in ['major', 'minor'] else current_patch + 1

    next_version = '%s.%s.%s' % (next_major, next_minor, next_patch)

    with open(args.filename, 'w') as f:
        f.write('%s\n' % next_version)

    print next_version


if __name__ == '__main__':
    main()
