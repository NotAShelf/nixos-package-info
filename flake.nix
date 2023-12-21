{
  description = "Description for the project";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixos-search.url = "github:nixos/nixos-search";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [];
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      perSystem = {
        inputs',
        self',
        pkgs,
        ...
      }: {
        packages = {
          default = self'.packages.nixos-package-info;
          update-sources = pkgs.callPackage ./packages/update-sources {};
          nixos-package-info = pkgs.callPackage ./packages/nixos-package-info {};
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            golint
            jq
          ];
        };
      };
    };
}
