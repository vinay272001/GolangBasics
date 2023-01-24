from typing import List, Tuple
def max_numbness(t: int, test_cases: List[Tuple[int, List[int]]]) -> List[int]:
    res = []
    for i in range(t):
        n = test_cases[i][0]
        a = test_cases[i][1]
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            dp[i] = max(dp[i-1], dp[i-1]^a[i-1])
        res.append(dp[n])
    return res

# test
test_cases = [(5, [1, 2, 3, 4, 5]), (3, [10, 47, 52])]
t = 2
print(max_numbness(t, test_cases))
