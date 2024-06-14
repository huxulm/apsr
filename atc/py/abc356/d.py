from functools import cache
I = lambda : map(int, input().split())
LI = lambda : list(I())

def solve():
    n, m = I()
    s = bin(n)[2:]
    MOD = 998244353
    @cache
    def f(i, cnt, isLimit):
        if i == len(s):
            return cnt
        up = int(s[i]) if isLimit else 1
        res = 0
        for d in range(up + 1):
            res = (res + f(i + 1, cnt + (d & (m >> (len(s) - i - 1))), isLimit and d == up)) % MOD
        return res
    ans = f(0, 0, True) % MOD
    print(ans)

if __name__ == '__main__':
    solve()