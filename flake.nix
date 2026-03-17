{
  description = "Generated Steampipe plugin for Akeyless";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        packages.default = pkgs.runCommand "steampipe-akeyless-gen" {
          src = self;
        } ''
          mkdir -p $out/share/steampipe
          cp -r $src/*.go $out/share/steampipe/ 2>/dev/null || true
          cp -r $src/tables $out/share/steampipe/ 2>/dev/null || true
          touch $out/share/steampipe/.generated
        '';

        # When go.mod exists and vendorHash is known, replace with:
        # packages.default = pkgs.buildGoModule {
        #   pname = "steampipe-plugin-akeyless";
        #   version = "0.1.0";
        #   src = self;
        #   vendorHash = null; # update after first build
        # };
      }
    );
}
