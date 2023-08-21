#include<iostream>
#include<vector>
#include<stack>

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
      if (n == 0) return "0";

      stack<string> s;
      while (n != 0) {
        s.push(n % 2 == 0 ? "0" : "1");
        n /= 2;
      }

      string r;
      while (!s.empty()) {
        r += s.top();
        s.pop();
      }

      return r;
    }
};

int main() {
  Solution solution;
  vector<vector<int>> queries = {
    {4, 2},
    {3, 3},
  };

  vector<vector<int>> result;
  result = solution.substringXorQueries("111010110", queries);
  for (auto &r : result) {
    printf("{ %d, %d }\n", r[0], r[1]);
  }
}
