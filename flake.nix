# SPDX-FileCopyrightText: 2025 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0
{
  description = "Codechef recipes";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        fhs = pkgs.buildFHSUserEnv {
          name = "fhs-shell";
          targetPkgs = pkgs: [
            pkgs.git
            pkgs.go_1_23
            pkgs.nodejs_24  # For pre-commit hooks
            pkgs.ruby_3_4  # For pre-commit hooks
          ];
        };
      in
      {
        #devShell = pkgs.mkShell { buildInputs = devDeps; };
        devShells.default = fhs.env;
      });
}
