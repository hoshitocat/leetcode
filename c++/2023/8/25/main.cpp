// https://leetcode.com/problems/interleaving-string/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    bool isInterleave(string s1, string s2, string s3) {
      return true;
    }
};

int main() {
  vector<vector<string>> cases = {
    vector<string>{"aabcc", "dbbca", "aadbbcbcac"},
    vector<string>{"aabcc", "dbbca", "aadbbbaccc"},
    vector<string>{"", "", ""},
  };

  Solution solution;
  for (auto &c : cases) {
    bool result = solution.isInterleave(c[0], c[1], c[2]);
    cout << result << endl;
  }

  return 0;
}
