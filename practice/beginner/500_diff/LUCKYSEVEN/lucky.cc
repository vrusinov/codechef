/* SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Lucky Seven
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/LUCKYSEVEN

#include <bits/stdc++.h>
#include <iostream>
using namespace std;

char lucky_seven(string s) {
    return s[6];
}

int main() {
    string s;
    cin >> s;
	cout << lucky_seven(s);

}
