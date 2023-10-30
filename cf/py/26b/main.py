def solve():
    s = input()
    ans = 0
    st = 0
    for c in s:
        if c == '(':
            st += 1
        else:
            if st:
                st -= 1
                ans += 1
    print(ans * 2)
if __name__ == "__main__":
    solve()   