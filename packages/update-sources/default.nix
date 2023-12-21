{
  stdenv,
  makeWrapper,
  lib,
  jq,
}: let
  inherit (lib.fileset) unions toSource;

  pname = "update-sources";
  version = "0.1.0";
in
  stdenv.mkDerivation {
    inherit pname version;

    src = toSource {
      root = ./.;
      fileset = unions [
        ./update.sh
      ];
    };

    buildInputs = [makeWrapper];
    phases = ["installPhase" "postInstall"];

    installPhase = ''
      runHook preInstall
      mkdir -p $out/bin
      cp $src/update.sh $out/bin/update-sources
      chmod +x $out/bin/update-sources
      runHook postInstall
    '';

    postInstall = ''
      wrapProgram $out/bin/update-sources \
        --prefix PATH : ${lib.makeBinPath [jq]}
    '';

    meta = {
      license = lib.licenses.mit;
      mainProgram = pname;
      maintainers = with lib.maintainers; [NotAShelf];
    };
  }
