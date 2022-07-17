#!/usr/bin/python3

COPY='c'
PUSH_BACK='p'

def copypush(line, prev_op=''):
    """Copy and Push Back.

    Idea: start from given line and try to apply Anon's actions in reverse.
    Try splitting or cutting characters from line to see if we can end up
    in empty string. Keep track of previous actions.
    """
    #print(f'{line}, op={prev_op}')
    l = len(line)
    if l == 0:
        return 'yes'
    # First, try to see if we can split the string into two equals.
    if l % 2 == 0:
        if line[:l//2] == line[l//2:]:
            if copypush(line[l//2:], prev_op=COPY) == 'yes':
                return 'yes'
    # We tried splitting and it didn't work. Let's see if we can append.
    if prev_op == PUSH_BACK:
        return 'no'
    return copypush(line[:-1], PUSH_BACK)



if __name__ == '__main__':
    t = input()
    for _ in range(int(t)):
        unused_n = input()
        line = input()
        print(copypush(line))
