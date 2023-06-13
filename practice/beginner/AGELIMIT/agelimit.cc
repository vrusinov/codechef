// https://www.codechef.com/problems/AGELIMIT

/* SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include <iostream>
using namespace std;

string age_limit(int x, int y, int a) {
    if ((a >= x) && (a < y)) {
        return "YES";
    }
    return "NO";
}

int main() {
	int t;
    cin >> t;
    for (int i = 0; i < t; i++) {
        int x, y, a;
        cin >> x;
        cin >> y;
        cin >> a;
        cout << age_limit(x, y, a) << "\n";
    }
	return 0;
}
