// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestExample(t *testing.T) {
	if luckyGame(3, 2, []int{1, 2, 3, 4, 5}) != "ALICE" {
		t.Errorf("Failed example 1")
	}
	if luckyGame(2, 4, []int{1, 2, 3, 4, 5}) != "BOB" {
		t.Errorf("Failed example 2")
	}
}

func TestEmpty(t *testing.T) {
	if luckyGame(2, 4, []int{}) != "ALICE" {
		t.Errorf("Failed on empty slice")
	}
}

func TestSame(t *testing.T) {
	if luckyGame(2, 2, []int{1, 2, 3, 4, 5}) != "BOB" {
		t.Errorf("Failed on same numbers")
	}
}