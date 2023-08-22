// https://leetcode.com/problems/excel-sheet-column-title/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    string convertToTitle(int columnNumber) {
      return "";
    }
};

int main() {
  Solution solution;

  vector<int> cases = {1, 28, 701};
  for (auto &num: cases) {
    cout << solution.convertToTitle(num) << endl;
  }
}

