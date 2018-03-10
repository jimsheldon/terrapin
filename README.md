# Terrapin

Terrapin is a command line tool that will install a specific version of [Terraform](https://www.terraform.io) to a desired directory.

Use terrapin in your continuous integration process to ensure the right terraform version is always used to apply your configuration. Leverage your continuous integration tool's caching functionality to ensure you only have to download the necessary terraform version once.

## Usage

```
$ terrapin
  -directory string
        Optional: destination directory (by default the current directory is used).
  -force
        Optional: overwrite terraform binary if found (default is false).
  -tf-version string
        terraform version to use.
```

## Examples

Download terraform 0.11.3 to the current directory:
```
$ terrapin -tf-version 0.11.3
terrapin: configuration:
terrapin:     -directory     = /Users/jsheldon/go/src/github.com/jimsheldon/terrapin
terrapin:     -force         = false
terrapin:     -tf-version    = 0.11.3
terrapin: downloading file https://releases.hashicorp.com/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip
terrapin: unzipped: /Users/jsheldon/go/src/github.com/jimsheldon/terrapin/terraform
terrapin: terraform 0.11.3 is available at /Users/jsheldon/go/src/github.com/jimsheldon/terrapin/terraform
```

Download terraform 0.11.3 to a different directory:
```
$ terrapin -tf-version 0.11.3 -directory ~/bin
terrapin: configuration:
terrapin:     -directory     = /Users/jsheldon/bin
terrapin:     -force         = false
terrapin:     -tf-version    = 0.11.3
terrapin: downloading file https://releases.hashicorp.com/terraform/0.11.3/terraform_0.11.3_darwin_amd64.zip
terrapin: unzipped: /Users/jsheldon/bin/terraform
terrapin: terraform 0.11.3 is available at /Users/jsheldon/bin/terraform
```

Running the same command again, the specified terraform binary is found and no download is needed:
```
$ terrapin -tf-version 0.11.3 -directory ~/bin
terrapin: configuration:
terrapin:     -directory     = /Users/jsheldon/bin
terrapin:     -force         = false
terrapin:     -tf-version    = 0.11.3
terrapin: terraform 0.11.3 is available at /Users/jsheldon/bin/terraform
```

By default, if terrapin finds a different version of terraform in the specified directory, it will fail:
```
$ terrapin -tf-version 0.11.2 -directory ~/bin
terrapin: configuration:
terrapin:     -directory     = /Users/jsheldon/bin
terrapin:     -force         = false
terrapin:     -tf-version    = 0.11.2
terrapin: wrong terraform version found at /Users/jsheldon/bin/terraform, pass the -force flag to overwrite it
```

Pass the `-force` flag to overwrite the terraform binary:
```
$ terrapin -tf-version 0.11.2 -directory ~/bin -force
terrapin: configuration:
terrapin:     -directory     = /Users/jsheldon/bin
terrapin:     -force         = true
terrapin:     -tf-version    = 0.11.2
terrapin: downloading file https://releases.hashicorp.com/terraform/0.11.2/terraform_0.11.2_darwin_amd64.zip
terrapin: unzipped: /Users/jsheldon/bin/terraform
terrapin: terraform 0.11.2 is available at /Users/jsheldon/bin/terraform
```
