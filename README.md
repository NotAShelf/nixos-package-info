# 🔍 nixos-package-info

> Tools and collections around parsing Nixpkgs package data

## 🎛️ Usage

**nixos-package-info** generates a JSON database for the latest nixpkgs commit every hour. The database contains **package_name**, **package_version** and
**package_description** by default. The `--full` flag can be used to display **package_longDescription** and **package_homepage** in addition.

The database is intended to be parsed by third party tools, possibly to replace the undesirable interface of `nix search`. No assumptions are made about
the formatting of output files, and parsing is left entirely to the user's discretion.

### Example Data

This is an example package field, generated without the `--full` flag and formatted with `jq` to be human readable.

```json
{
  "package_description": "A simple terminal UI for both docker and docker-compose",
  "package_name": "lazydocker",
  "package_version": "0.23.1"
}
```

With the `nixos-package-info` tool we extract only the relevant sections. Below is an example with the `--full` flag for extended
information.

```json
{
  "package_description": "Integration Services for running NixOS under HyperV",
  "package_homepage": [
    "https://kernel.org"
  ],
  "package_longDescription": "<rendered-html><p>This packages contains the daemons that are used by the Hyper-V\nhypervisor on the host.</p>\n<p>Microsoft calls their guest ag
ents “Integration Services” which is\nwhy we use that name here.</p>\n</rendered-html>",
  "package_name": "linuxKernel.packages.linux_zen.hyperv-daemons",
  "package_version": "6.6.7"
},
```

Once again, this is a formatted example (aside from the html tags in longDescription, those are not our doing) with the help of `jq`.

## 📦 Packages

The repository contains several packages, each with their own purpose. Below is an explanation of each package.

### update-sources

Updates the Nixpkgs revision in `data/target.json`. The `target.json` file is used by the [flake-info](https://github.com/NixOS/nixos-search/tree/main/flake-info)
tool to generate raw Nixpkgs data.

#### Usage

```console
nix run .#update-sources data/targets.json
```

### nixos-package-info

The main tool of this flake, it uses data generated by `flake-info` to generate a much smaller database with relevant package info.
By default, it returns package name, version and description. Using the `--full` flag will also return long description and the package
homepage.

#### Usage

Generate lightweight database

```console
nix run .#nixos-package-info -- --input data/nixpkgs.json > database.json
```

Generate full database

```console
nix run .#nixos-package-info -- --full --input data/nixpkgs.json > database-full.json
```

## 🛠️ Hacking

You can use `direnv` or `nix develop` to enter a shell with required packages. All "source code" is available
under `packages/`

## ✨ Contributing

Contributions are always welcome. If you would like to make a change, but cannot decide how to go about it
please open an issue and we will discuss it.

## 📜 License

`nixos-package-info` is licensed under EUPL-1.2 license. See [LICENSE](LICENSE) for more details.
