def solve():
  n = int(input())
  a = [list(map(int, input().split())) for _ in range(n)]
  a.sort()
  rmx0 = rmx1 = -1
  for L, R in a:
    if L > rmx0:
      rmx0 = max(rmx0, R)
    elif L > rmx1:
      rmx1 = max(rmx1, R)
    else:
      print("NO")
      return
  print("YES")
solve()
