// https://leetcode.com/problems/spiral-matrix-ii/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    vector<vector<int>> generateMatrix(int n) {
      if (n == 1) return {{1}};

      vector<vector<int>> matrix(n, vector<int>(n, 0));

      pair<int, int> current = {0, 0};
      vector<int> directions = {0, 1, 0, -1, 0};
      int right = 0;
      int down = 1;
      int left = 2;
      int up = 3;
      int dir = right;
      for (int i = 1; i <= n*n; i++) {
        matrix[current.first][current.second] = i;

        int next_x = current.first + directions[dir];
        int next_y = current.second + directions[dir+1];
        if (next_x < 0 ||
            next_y < 0 ||
            n <= next_x ||
            n <= next_y ||
            matrix[next_x][next_y] != 0) {
          switch(dir) {
            case 0:
              dir = down;
              break;
            case 1:
              dir = left;
              break;
            case 2:
              dir = up;
              break;
            case 3:
              dir = right;
              break;
          }
        }

        current.first += directions[dir];
        current.second += directions[dir+1];
      }

      return matrix;
    }
};

int main() {
  Solution solution;

  vector<int> cases = {3, 1, 4};
  for (auto c : cases) {
    vector<vector<int>> result = solution.generateMatrix(c);
    for (auto &r : result) {
      cout << "{";
      for (auto n : r) {
        cout << n << ",";
      }
      cout << "}," << endl;
    }
  }
}
