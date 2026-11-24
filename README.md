# Matrix Chain Multiplication

> Striver SDE Sheet 176 · Day 25 DP — interval DP: solve every chain segment shortest-first, choosing where the final multiplication splits it. Go, stdlib only, tested.

## Problem

Matrix multiplication is associative but the parenthesization changes the work enormously: multiplying a 10×20, 20×30, 30×40 chain costs 18,000 scalar multiplications one way and 24,000 the other. Find the minimum for a whole chain (matrix i is dims[i]×dims[i+1]).

## Approach

The archetype of **interval DP**: the state is a segment of the chain, and the decision is where the *outermost* multiplication splits it — `dp[i][j] = min over k of dp[i][k] + dp[k+1][j] + dims[i]·dims[k+1]·dims[j+1]`, the last term being the final multiply of the two resulting blocks. Filling by increasing segment length guarantees both halves are ready when needed. The same solve-segments-shortest-first shape powers palindrome partitioning II, burst balloons, and polygon triangulation — MCM is the one to learn it on.

**Complexity:** O(n³) time, O(n²) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
