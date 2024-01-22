# This Code TLE!
# See go code: github.com/huxulm/apsr/cf/go/1748B/main.go
IS = lambda: input()
II = lambda: int(IS())
LI = lambda: map(int, IS().split())
def solve():
  t = II()
  for _ in range(t):
    ans = 0
    n, s = II(), IS()
    for i in range(n):
      cnt = [0] * 10
      k = mxC = 0
      for j in range(i, n):
        v = ord(s[j]) - ord('0')
        if cnt[v] == 10:
          break
        k += cnt[v] == 0
        cnt[v] += 1
        mxC = max(mxC, cnt[v])
        ans += mxC <= k
    print(ans)

if __name__ == '__main__':
  solve()



