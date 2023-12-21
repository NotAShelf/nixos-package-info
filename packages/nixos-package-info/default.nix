{
  buildGoModule,
  lib,
  ...
}: let
  inherit (lib.fileset) unions toSource fileFilter difference;

  filterNixFiles = fileFilter (file: lib.hasSuffix ".nix" file.name) ./.;
  baseSrc = unions [
    ./internal
    ./main.go
    ./go.mod
  ];

  filter = difference baseSrc filterNixFiles;

  pname = "nixos-package-info";
  version = "0.1.0";
in
  buildGoModule {
    inherit pname version;

    src = toSource {
      root = ./.;
      fileset = filter;
    };

    vendorHash = null;
    doCheck = false;
    ldflags = ["-s" "-w"];

    meta = {
      description = "A simple parser to get packages and their names from flake-info tool output";
      license = lib.licenses.gpl3Only;
      maintainers = with lib.maintainers; [NotAShelf];
    };
  }
