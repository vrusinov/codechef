// https://www.codechef.com/problems/TAXSAVING

/* SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include <iostream>
using namespace std;

int tax_saving(int x, int y) {
    return x - y;
}

int main() {
	int t;
    cin >> t;
    for (int i = 0; i < t; i++) {
        int x, y;
        cin >> x;
        cin >> y;
        cout << tax_saving(x, y) << "\n";
    }
	return 0;
}
