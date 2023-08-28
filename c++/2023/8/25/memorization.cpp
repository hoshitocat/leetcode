// https://leetcode.com/problems/interleaving-string/
#include<iostream>
#include<vector>
#include<unordered_map>

using namespace std;

const int null = 2;

class Solution {
  string s1, s2, s3;
  vector<vector<int>> memo;

  public:
    bool isInterleave(string s1, string s2, string s3) {
      this->s1 = s1;
      this->s2 = s2;
      this->s3 = s3;

      int a = s1.size();
      int b = s2.size();
      int c = s3.size();

      if (a+b != c) return false;

      memo = vector<vector<int>>(a+1, vector<int>(b+1, null));

      return helper(0, 0, 0);
    }

  private:
    bool helper(int a, int b, int c) {
      if (c == s3.size()) return true;

      int m = memo[a][b];
      if (m != null) return m;

      bool ans = false;
      if (a < s1.size() && s1[a] == s3[c]) {
        ans = ans || helper(a+1, b, c+1);
      }

      if (b < s2.size() && s2[b] == s3[c]) {
        ans = ans || helper(a, b+1, c+1);
      }

      memo[a][b] = ans;
      return ans;
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
