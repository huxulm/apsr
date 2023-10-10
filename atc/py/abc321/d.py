# https://atcoder.jp/contests/abc321/tasks/abc321_d
from bisect import *
from itertools import *
I = lambda : map(int, input().split())
LI = lambda : list(I())
def solve():
  n, m, p = I()
  a = LI()
  b = LI()
  b.sort()
  preSum = list(accumulate(b))
  ans = 0
  for x in a:
    i = bisect_right(b, p - x) # [0, i) 取x + y, [i, n) 取 p
    ans += x * i + (preSum[i-1] if i else 0) + (m - i) * p
  print(ans)

solve()

"""
输入
2 2 7
3 5
6 1
输出
24
"""
