#include<iostream>
#include<vector>
#include<unordered_map>

using namespace std;

class Solution {
  public:
    vector<vector<int>> substringXorQueries(string s, vector<vector<int>>& queries) {
      unordered_map<int, vector<int>> mp;
      int len = s.length();
      for (int i = 0; i < len; i++) {
        int x = 0;
        for (int j = 0; i+j < len && j < 32; j++) {
          x = (x << 1) + (s[j+i] == '1');
          if (mp.count(x) == 0) {
            mp[x].push_back(i);
            mp[x].push_back(i+j);
          }

          if (s[i] == '0') break;
        }
      }

      vector<vector<int>> ans;
      for (auto &i : queries) {
        int xored = i[0] ^ i[1];
        if (mp.count(xored))
          ans.push_back(mp[xored]);
        else
          ans.push_back({-1, -1});
      }
      return ans;
    }
};

int main() {
  Solution solution;

  string s = "101101";
  vector<vector<int>> queries = {
    {0, 5},
    {1, 2},
  };


  vector<vector<int>> result;
  result = solution.substringXorQueries(s, queries);
  for (auto &r : result) {
    printf("{ %d, %d }\n", r[0], r[1]);
  }
}
