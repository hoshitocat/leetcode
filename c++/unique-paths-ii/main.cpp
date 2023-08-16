#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
      vector<vector<int>> dp;
      for (int i = 0; i < obstacleGrid.size(); i++) {
        vector<int> val(obstacleGrid[i].size(), 0);
        dp.push_back(val);
      }



      for (int i = 0; i < dp.size(); i++) {
        for (int j = 0; j < dp[i].size(); j++) {
          pair<int, int> left = {i, j - 1};
          pair<int, int> up = {i - 1, j};
          // WIP
        }
      }

      // debug
      for (int i = 0; i < obstacleGrid.size(); i++) {
        for (int j = 0; j < obstacleGrid[i].size(); j++) {
          cout << dp[i][j];
        }
        cout << endl;
      }

      return 0;
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
