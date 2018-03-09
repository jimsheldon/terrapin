# Terrapin

Terrapin is a command line tool to install a specific version of [Terraform](https://www.terraform.io) to a desired directory.

## Usage

```
$ terrapin
  -directory string
        Optional: destination directory (if unspecified current directory is used).
  -force
        Optional: force removal of terraform binary if found (default is false).
  -tf-version string
        terraform version to use.
```

# Examples

Download `terraform` 0.11.3 to the current directory
```
$ terrapin -tf-version 0.11.3
```
