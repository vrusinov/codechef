// https://www.codechef.com/problems/SQUATS

/* SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include <iostream>
using namespace std;

int main() {
	int t;
    cin >> t;
    for (int i = 0; i < t; i++) {
        int x;
        cin >> x;
        cout << x*15 << "\n";
    }
	return 0;
}
