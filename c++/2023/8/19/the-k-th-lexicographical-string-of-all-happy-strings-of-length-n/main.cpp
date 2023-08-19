#include<iostream>
#include<vector>
#include<queue>

using namespace std;

class Solution {
  public:
    string getHappyString(int n, int k) {
      queue<string> q;

      q.push("a");
      q.push("b");
      q.push("c");
      while (q.front().size() < n) {
        string tmp = q.front();
        q.pop();

        if (tmp[tmp.size() - 1] == 'a') {
          q.push(tmp + "b");
          q.push(tmp + "c");
        }
        if (tmp[tmp.size() - 1] == 'b') {
          q.push(tmp + "a");
          q.push(tmp + "c");
        }
        if (tmp[tmp.size() - 1] == 'c') {
          q.push(tmp + "a");
          q.push(tmp + "b");
        }
      }

      string result = "";
      for (int i = 0; i < k; i++) {
        if (q.empty())
          return "";

        result = q.front();
        q.pop();
      }

      return result;
    }
};

int main() {
  Solution solution;
  cout << solution.getHappyString(1, 3) << endl;
  cout << solution.getHappyString(1, 4) << endl;
  cout << solution.getHappyString(3, 9) << endl;
}
