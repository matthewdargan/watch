{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    parts.url = "github:hercules-ci/flake-parts";
  };
  outputs = inputs:
    inputs.parts.lib.mkFlake {inherit inputs;} {
      perSystem = {pkgs, ...}: {
        packages = {
          watch = pkgs.buildGoModule {
            pname = "watch";
            src = ./.;
            vendorHash = null;
            version = "0.1.3";
          };
        };
      };
      systems = ["aarch64-darwin" "aarch64-linux" "x86_64-darwin" "x86_64-linux"];
    };
}
