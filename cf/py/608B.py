""" 
01
00111
"""
def solve():
    s, t = input(), input()
    m, n = len(s), len(t)
    # 在 b 上滑窗
    # d = n - m + 1 窗口长度
    d = n - m + 1
    ans = 0
    one = t[0:d-1].count('1')
    for i, ch in enumerate(s):
        one += t[i + d - 1] == '1'
        ans += one if ch == '0' else (d - one)
        one -= t[i] == '1'
    return ans

if __name__ == '__main__':
    print(solve())
