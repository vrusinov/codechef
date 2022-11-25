<!--
SPDX-FileCopyrightText: 2022 Google Inc
SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>

SPDX-License-Identifier: Apache-2.0
-->

# [Lucky Number Game](https://www.codechef.com/submit-v2/HP18)

Solution in go.

Bob and Alice are playing a game with the following rules:

Initially, they have an integer sequence A_1, A_2, \ldots, A_NA
1
​
 ,A
2
​
 ,…,A
N
​
 ; in addition, Bob has a lucky number aa and Alice has a lucky number bb.
The players alternate turns. In each turn, the current player must remove a non-zero number of elements from the sequence; each removed element should be a multiple of the lucky number of this player.
If it is impossible to remove any elements, the current player loses the game.
It is clear that one player wins the game after a finite number of turns. Find the winner of the game if Bob plays first and both Bob and Alice play optimally.