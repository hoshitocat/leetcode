// https://leetcode.com/problems/excel-sheet-column-title/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    string convertToTitle(int columnNumber) {
      const int base = 26;
      vector<string> titles = {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"};

      string result = "";

      int remain = columnNumber % base;
      if (remain == 0) {
        result = titles[base-1];
      } else {
        result = titles[remain-1];
      }

      while (0 < columnNumber) {
        int quotient = columnNumber / base;
        if (quotient <= base) {
          result = titles[quotient-1] + result;
        } else {
          result = titles[base-1] + result;
        }

        columnNumber = quotient - base;
      }

      return result;
    }
};

int main() {
  Solution solution;

  vector<int> cases = {1, 28, 701, 1000};
  for (auto &num : cases) {
    cout << solution.convertToTitle(num) << endl;
  }
}

