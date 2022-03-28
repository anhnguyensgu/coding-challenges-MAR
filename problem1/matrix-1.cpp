#include<bits/stdc++.h>

using namespace std;

int main() {
    int n = 4;
    int matrix[n][n] = {
        {0, 1, 2, 3}, 
        {3,1,2,4},
        {1,0,2,3},
        {5,9,2,5}
    };

    int checkRow = 0;
    int checkCol = 0;

    for(int i = 0; i < n; i++) {
        for(int j = 0; j < n; j++) {
            if (matrix[i][j] == 1) {
                checkRow |= 1 << i;
                checkCol |= 1 << j;
            }
        }
    }

    for(int i = 0; i < n; i++) {
        int isOne = checkRow & (1 << i);
        if (isOne != 0) {
            for(int j = i; j < n; j++) {
                matrix[i][j] = 1;
            }
        }

        isOne = checkCol & (1 << i);    
        if (isOne != 0) {
            for(int j = i; j < n; j++) {
                matrix[j][i] = 1;
            }
        }
    }

    for(int i = 0; i < n; i++) {
        for(int j = 0; j < n; j++) {
            cout << matrix[i][j] << " ";
        }
        cout << endl;
    }
}
//time complexity: O(n^2)
//space complexity: O(1)