// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150

#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
  public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
      if (n == 0) return;

      for (int j = 0; j < n; j++) {
        nums1[m+j] = nums2[j];
      }

      sort(nums1.begin(), nums1.end());
    }
};

struct input {
  vector<int> nums1;
  int m;
  vector<int> nums2;
  int n;
};

void printVector(vector<int> arr) {
  int size = arr.size();
  if (size == 0) {
    cout << "[]";
    return;
  }

  cout << "[";
  for (int i = 0; i < size-1; i++) {
    cout << arr[i] << ", ";
  }
  cout << arr[size-1];
  cout << "]";
}

int main() {
  Solution solution;

  vector<input> inputs = {
    {
      {1, 2, 3, 0, 0, 0},
      3,
      {2, 5, 6},
      3,
    }
  };

  for (auto &i : inputs) {
    cout << "input: ";
    cout << "nums1 = ";
    printVector(i.nums1);
    cout << ", ";
    cout << "m = " << i.m;
    cout << ", ";
    cout << "nums2 = ";
    printVector(i.nums2);
    cout << ", ";
    cout << "n = " << i.n;
    cout << endl;

    solution.merge(i.nums1, i.m, i.nums2, i.n);

    cout << "out: ";
    printVector(i.nums1);
    cout << endl;
  }
}
