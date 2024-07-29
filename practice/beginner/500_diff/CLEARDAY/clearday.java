/*
 * SPDX-FileCopyrightText: 2024 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Clear Day
 * https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/CLEARDAY
 */

import java.util.*;

class Codechef {
  public static void main(String[] args) throws java.lang.Exception {
    Scanner scanner = new Scanner(System.in);
    int x = scanner.nextInt();
    int y = scanner.nextInt();

    int result = 7 - x - y;
    System.out.println(result);

    scanner.close();
  }
}
