
lines = [line.rstrip('\n') for line in open("switches.txt")]
switches = [False] * int(lines[0])
for line in lines[1:len(lines)]:
    values = line.split()

    low, hi = int(values[0]), int(values[1])
    if low > hi:
        low, hi = hi, low

    switches[low:hi+1] = [not sw for sw in switches[low:hi+1]]

print(switches.count(True))
