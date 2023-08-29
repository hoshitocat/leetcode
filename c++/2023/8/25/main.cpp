// https://leetcode.com/problems/interleaving-string/
#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    bool isInterleave(string s1, string s2, string s3) {
      int m = s1.size();
      int n = s2.size();
      int l = s3.size();
      if (m+n != l) return false;

      if (m < n) isInterleave(s2, s1, s3);

      vector<bool> dp(n+1, false);
      dp[0] = true;

      for (int j = 1; j < n+1; j++) {
        dp[j] = dp[j-1] && s2[j-1] == s3[j-1];
      }

      for (int i = 1; i < m+1; i++) {
        dp[0] = dp[0] && s1[i-1] == s3[i-1];
        for (int j = 1; j < n+1; j++) {
          dp[j] = (dp[j] && s1[j-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1]);
        }
      }

      return dp[n];
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
