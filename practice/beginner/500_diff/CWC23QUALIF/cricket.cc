/* SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

/* https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/CWC23QUALIF
 *
 * Cricket World Cup Qualifier
 *
 * The cricket World Cup has started in Chefland. There are many teams
 * participating in the group stage matches. Any team that scores 12 or more
 * points in the group stage matches qualifies for the next stage.
 *
 * You know the score that a particular team has scored in the group stage
 * matches. Determine if the team has qualified for the next stage or not.
 */

#include <bits/stdc++.h>
#include <iostream>
using namespace std;

int main() {
  // your code goes here
  int x;
  cin >> x;

  if (x >= 12) {
    cout << "Yes";
  } else {
    cout << "No";
  }
}
