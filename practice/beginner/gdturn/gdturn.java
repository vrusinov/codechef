/*
 * SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 */

/* Good Turn
 * https://www.codechef.com/problems/GDTURN
 */

/* package codechef; // don't place package name! */

import java.util.*;
import java.lang.*;
import java.io.*;

class Codechef
{
	public static void main (String[] args) throws java.lang.Exception
	{
        Scanner scanner = new Scanner(System.in);
        String line = scanner.nextLine();
        int t = Integer.parseInt(line);
        while (t-- > 0) {
            line = scanner.nextLine();
            String[] a = line.split(" ");
            int x = Integer.parseInt(a[0]);
            int y = Integer.parseInt(a[1]);

            System.out.println(solve(x, y));
        }
	}

    public static String solve(int x, int y) {
        if ((x + y) > 6) {
            return "YES";
        } else {
            return "NO";
        }

    }
}
