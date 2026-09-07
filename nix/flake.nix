{
  description = "CLI tool for managing Antybrowser profiles, proxies, and automations";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        antybrowser-cli = pkgs.callPackage ./default.nix { };
      in
      {
        packages.default = antybrowser-cli;
        packages.antybrowser-cli = antybrowser-cli;

        devShells.default = pkgs.mkShell {
          buildInputs = [ antybrowser-cli ];
        };
      }
    );
}
