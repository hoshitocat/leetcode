// https://leetcode.com/problems/reorganize-string/
#include<iostream>
#include<vector>
#include<unordered_map>
#include<queue>

using namespace std;

class Solution {
  public:
    string reorganizeString(string s) {
      unordered_map<char, int> mp;
      for (auto c : s) {
        mp[c]++;
      }

      priority_queue<pair<int, char>> q;
      for (auto &[c, freq] : mp) {
        q.push({freq, c});
      }

      string res = "";
      while (q.size() >= 2) {
        auto [freq1, char1] = q.top(); q.pop();
        auto [freq2, char2] = q.top(); q.pop();

        res += char1;
        res += char2;

        if (--freq1 > 0) q.push({freq1, char1});
        if (--freq2 > 0) q.push({freq2, char2});
      }

      if (!q.empty()) {
        auto [freq, ch] = q.top();
        if (freq > 1) return "";
        res += ch;
      }

      return res;
    }
};

int main() {
  Solution solution;
  vector<string> cases = {"aab", "aaab"};

  for (string &s : cases) {
    cout << solution.reorganizeString(s) << endl;
  }
}
