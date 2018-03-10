# Terrapin

Terrapin is a command line tool that will install a specific version of [Terraform](https://www.terraform.io) to a desired directory.

Use terrapin in your continuous integration process to ensure the right terraform version is always used to apply your configuration. Leverage your continuous integration tool's caching functionality to ensure you only have to download the necessary terraform version once.

## Usage

```
$ terrapin
  -directory string
        Optional: destination directory (by default the current directory is used).
  -force
        Optional: force removal of terraform binary if found (default is false).
  -tf-version string
        terraform version to use.
```

## Examples

Download terraform 0.11.3 to the current directory:
```
$ terrapin -tf-version 0.11.3
```

Download terraform 0.11.3 to a different directory:
```
$ terrapin -tf-version 0.11.3 -directory ~/bin
```

Download terraform 0.11.3 to a different directory, overwriting any existing terraform binary in that directory:
```
$ terrapin -tf-version 0.11.3 -directory ~/bin -force
```