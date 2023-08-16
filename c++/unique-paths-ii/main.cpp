#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
      cout << "hoge" << endl;
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
