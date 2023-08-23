// https://leetcode.com/problems/excel-sheet-column-title/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    string convertToTitle(int columnNumber) {
      string result = "";
      while (columnNumber > 0) {
        columnNumber--;
        char c = 'A' + columnNumber % 26;
        result = c + result;
        columnNumber /= 26;
      }

      return result;
    }
};

int main() {
  Solution solution;

  vector<int> cases = {1, 28, 701, 702, 1000};
  for (auto &num : cases) {
    cout << solution.convertToTitle(num) << endl;
  }
}

