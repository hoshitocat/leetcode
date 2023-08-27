// https://leetcode.com/problems/interleaving-string/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    bool isInterleave(string s1, string s2, string s3) {
      int n = s1.size();
      int m = s2.size();

      if ((n+m) != s3.size()) return false;

      n++;
      m++;

      vector<vector<bool>> dp(n, vector<bool>(m, false));
      dp[0][0] = true;

      for (int i = 1; i < n; i++) {
        dp[i][0] = dp[i-1][0] && (s1[i-1] == s3[i-1]);
      }

      for (int j = 1; j < m; j++) {
        dp[0][j] = dp[0][j-1] && (s2[j-1] == s3[j-1]);
      }

      for (int i = 1; i < n; i++) {
        for (int j = 1; j < m; j++) {
          bool isMatchS1 = dp[i-1][j] && (s1[i-1] == s3[i+j - 1]);
          bool isMatchS2 = dp[i][j-1] && (s2[j-1] == s3[i+j - 1]);

          dp[i][j] = isMatchS1 || isMatchS2;
        }
      }

      return dp[n-1][m-1];
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
