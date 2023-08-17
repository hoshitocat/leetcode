#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
      int m, n; m = obstacleGrid.size(); n = obstacleGrid[0].size();

      vector<vector<int>> dp;
      for (int i = 0; i < m; i++) {
        vector<int> val(n, 0);
        dp.push_back(val);
      }

      dp[0][0] = 1;
      for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
          if (obstacleGrid[i][j] == 1) {
            dp[i][j] = 0;
            continue;
          }

          if (i != 0) {
            dp[i][j] += dp[i-1][j];
          }

          if (j != 0) {
            dp[i][j] += dp[i][j-1];
          }
        }
      }

      return dp[m-1][n-1];
    }
};

int main() {
  vector<vector<int>> input = {
    {0, 0, 0},
    {0, 1, 0},
    {0, 0, 0},
  };

  Solution solution;
  solution.uniquePathsWithObstacles(input);
  return 0;
}
