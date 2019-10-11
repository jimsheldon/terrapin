package install

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli"
)

// Command exports the install command
var Command = cli.Command{
	Name:      "install",
	Usage:     "install Terraform binary",
	ArgsUsage: "<source>",
	Action:    install,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrite terraform binary if found",
		},
		cli.StringFlag{
			Name:  "directory, d",
			Usage: "destination directory",
			Value: ".",
		},
		cli.StringFlag{
			Name:     "version, v",
			Usage:    "terraform version to install",
			Required: true,
		},
	},
}

func downloadTerraform(version string, directory string) error {
	filePath := directory + "/terraform.zip"
	url := "https://releases.hashicorp.com/terraform/" + version + "/terraform_" + version + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".zip"

	// create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// get the data
	log.Println("downloading file", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		os.Remove(filePath)
		return fmt.Errorf("failed to download %s, status code received: %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	filenames, err := unzip(out.Name(), directory)
	if err != nil {
		return err
	}

	log.Println("unzipped: " + strings.Join(filenames, ", "))

	err = os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}

	return nil
}

// unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {
			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()
			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}

func install(c *cli.Context) error {
	log.SetFlags(0)

	// set directory to current directory if directory flag was not set
	directory := c.String("directory")
	force := c.Bool("force")
	version := c.String("version")

	if directory == "." {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		directory = dir
	}

	// verify the directory exists
	_, err := os.Stat(directory)
	if err != nil {
		log.Fatalln(err)
	}
	tfPath := directory + "/terraform"

	// if the file does not exist, or if the -force flag was set, download terraform
	_, err = os.Stat(tfPath)
	if err != nil || force {
		if os.IsNotExist(err) || force {
			err := downloadTerraform(version, directory)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	// ensure 'terraform version' output includes "Terraform vX.Y.Z"
	out, err := exec.Command(tfPath, "version").Output()
	if err != nil {
		log.Fatalln(err)
	}
	versionString := "Terraform v" + version
	if !strings.Contains(string(out), versionString) {
		log.Fatalf("wrong terraform version found at %s, pass the --force flag to overwrite it", tfPath)
	}

	log.Println("terraform", version, "is available at", tfPath)
	return nil
}
