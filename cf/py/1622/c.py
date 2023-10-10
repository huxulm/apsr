def solve():
	t = int(input())
	for _ in range(t):
		n, k = map(int, input().split())
		a = sorted(map(int, input().split()))
		for i in range(1, n):
			a[i] += a[i - 1]

		r = float('inf')
		for y in range(n):
			dr = a[0] * y + a[n - y - 1] - k # 减小后的总和 - k = (需要减少的总数)
			x = 0 if dr <= 0 else (dr + y) // (y + 1)
			if x + y > r: # 减小一不会导致总次数减少了
				break
			r = x + y
		print(r)

if __name__ == "__main__":
	solve()





		




