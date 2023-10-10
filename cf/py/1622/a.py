def solve():
	t = int(input())
	for _ in range(t):		
		a = sorted(map(int, input().split()))
		if (a[-1] == a[0] + a[1]) or (a[0] % 2 == 0 and a[1] == a[-1]) or (a[0] == a[1] and a[-1] & 1 == 0):
			print("YES")
		else:
			print("NO")


if __name__ == "__main__":
	solve()