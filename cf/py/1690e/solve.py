""" 
输入
6
6 3
3 2 7 1 4 8
4 3
2 1 5 6
4 12
0 0 0 0
2 1
1 1
6 10
2 0 0 5 9 4
6 5
5 3 8 6 3 2
输出
8
4
0
2
1
5
"""

I = lambda : map(int, input().split())
LI = lambda : list(I())
def Print(arr):
    for v in arr: print(v)

def solve():
    t = LI()[0]
    res = []
    for _ in range(t):
        n, k = I()
        a = LI()
        ans = 0
        for i, x in enumerate(a):
            ans += x // k
            a[i] = x % k
        a.sort()
        i, j = 0, n - 1
        while i < j:
            if a[i] + a[j] < k:
                i += 1
            else:
                ans += 1
                i += 1
                j -= 1
        res.append(ans)
    Print(res)

solve()
