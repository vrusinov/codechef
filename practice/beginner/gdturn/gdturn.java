/*
 * SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vladimir.rusinov@gmail.com>
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Append for OR
 * https://www.codechef.com/problems/APPENDOR
 */

/* package codechef; // don't place package name! */


import java.util.*;
import java.io.*;

/* Name of the class has to be "Main" only if the class is public. */
class Codechef
{
    private static int[] StrToArray(String s) {
        String[] stringArray = s.split(" ");
        int[] result = new int[stringArray.length];
        for (int i = 0; i < stringArray.length; i++) {
            result[i] = Integer.parseInt(stringArray[i]);
        }
        return result;
    }

	public static void main (String[] args) throws java.lang.Exception
	{
        Scanner scanner = new Scanner(System.in);
        String line = scanner.nextLine();
        int t = Integer.parseInt(line);
        while (t-- > 0) {
            line = scanner.nextLine();
            int[] a = StrToArray(line);
            int n = a[0];
            int y = a[1];
            line = scanner.nextLine();
            a = StrToArray(line);

            System.out.println(solve(n, y, a));
        }
	}

    private static int solve(int n, int y, int[] a) {
        int bitWiseOr = 0;
        for(int element : a) {
            bitWiseOr |= element;
        }
        int maybeResult =  bitWiseOr ^ y;
        if ((maybeResult | bitWiseOr) == y) {
            return maybeResult;
        }
        return -1;
    }
}
