#include<iostream>
#include<vector>

using namespace std;

class Solution {
  public:
    vector<vector<int>> substringXorQueries(string s, vector<vector<int>>& queries) {
      vector<vector<int>> result;
      for (int i = 0; i < queries.size(); i++) {
        int first = queries[i][0];
        int second = queries[i][1];
        int xored = first ^ second;
        string b = toBinary(xored);
        size_t index = s.find(b);
        if (index == string::npos) {
          result.push_back({-1, -1});
          continue;
        }

        int left = static_cast<int>(index);
        int right = static_cast<int>(b.size() - 1);

        result.push_back({left, left+right});
      }

      return result;
    }

  private:
    string toBinary(int n) {
      string r;
      while (n != 0) {
        r += (n % 2 == 0 ? "0" : "1");
        n /= 2;
      }

      return r;
    }
};

int main() {
  Solution solution;
  vector<vector<int>> queries = {
    {0, 5},
    {1, 2},
  };

  vector<vector<int>> result = solution.substringXorQueries("101101", queries);
  for (int i = 0; i < result.size(); i++) {
    printf("{ %d, %d }\n", result[0], result[1]);
  }
}
