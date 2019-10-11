[![Build Status](https://cloud.drone.io/api/badges/jimsheldon/terrapin/status.svg)](https://cloud.drone.io/jimsheldon/terrapin)

# Terrapin

Terrapin is a command line tool that will install a specific version of [Terraform](https://www.terraform.io) to a desired directory.

Use terrapin in your continuous integration process to ensure the right terraform version is always used to apply your configuration. Leverage your continuous integration tool's caching functionality to ensure you only have to download the necessary terraform version once.

## Usage

```
$ terrapin -h
NAME:
   terrapin - install a specific version of Terraform to a desired directory

USAGE:
   terrapin [global options] command [command options] [arguments...]

VERSION:
   undefined

COMMANDS:
   install  install Terraform binary
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --print-version, -V  print only the version
```

## Examples

Download terraform 0.11.3 to the current directory:
```console
$ terrapin install --version 0.11.3
downloading file https://releases.hashicorp.com/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip
unzipped: /Users/janedoe/project/terraform
```

Download terraform 0.11.3 to a different directory:
```console
$ terrapin install --version 0.11.3 --directory ~/bin
downloading file https://releases.hashicorp.com/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip
unzipped: /Users/janedoe/bin/terraform
terraform 0.11.3 is available at /Users/janedoe/bin/terraform
```

Running the same command again, the specified terraform binary is found and no download is needed:
```console
$ terrapin install --version 0.11.3 --directory ~/bin
terraform 0.11.3 is available at /Users/janedoe/bin/terraform
```

By default, if terrapin finds a different version of terraform in the specified directory, it will fail:
```console
$ terrapin install --version 0.11.4 --directory ~/bin
wrong terraform version found at /Users/janedoe/bin/terraform, pass the --force flag to overwrite it
```

Pass the `--force` flag to overwrite the terraform binary:
```
$ terrapin install --version 0.11.4 --directory ~/bin --force
downloading file https://releases.hashicorp.com/terraform/0.11.4/terraform_0.11.4_darwin_amd64.zip
unzipped: /Users/janedoe/bin/terraform
terraform 0.11.4 is available at /Users/janedoe/bin/terraform
```