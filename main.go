package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli"
)

func listYamlFiles(rootpath string) ([]string, error) {
	var files []string
	var yamlRegex = regexp.MustCompile(`(?mi).*\.(yaml|yml)`)
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && yamlRegex.MatchString(info.Name()) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func listImageReferences(filepath string) ([]string, error) {
	data, err := ioutil.ReadFile(filepath)
	var imageRefs []string
	var imageRefRegex = regexp.MustCompile(`(?m)image:\s*(?P<image>[^{\s]+)\s+`)
	for _, match := range imageRefRegex.FindAllStringSubmatch(string(data), -1) {
		if len(match) > 1 {
			imageRefs = append(imageRefs, match[1])
		}
	}
	return imageRefs, err
}

func main() {
	var scanpath, registry string
	app := cli.NewApp()
	app.Name = "korp"
	app.Usage = "push images to a corporate registry based on Kubernetes yaml files"

	app.Commands = []cli.Command{
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "collects images referenced in yaml files and creates a kustomization file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "files, f",
					Value:       ".",
					Usage:       "path of yaml files to scan",
					Destination: &scanpath,
				},
			},
			Action: func(c *cli.Context) error {
				paths, _ := listYamlFiles(scanpath)
				for _, yamlPath := range paths {
					imageRefs, _ := listImageReferences(yamlPath)
					if len(imageRefs) > 0 {
						fmt.Println(listImageReferences(yamlPath))
					}
				}
				return nil
			},
		},
		{
			Name:    "pull",
			Aliases: []string{"p"},
			Usage:   "pulls referenced images to the local docker registry",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "path, p",
					Value:       ".",
					Usage:       "path of the kustomize file",
					Destination: &scanpath,
				},
			},
			Action: func(c *cli.Context) error {
				paths, _ := listYamlFiles(scanpath)
				for _, yamlPath := range paths {
					imageRefs, _ := listImageReferences(yamlPath)
					if len(imageRefs) > 0 {
						fmt.Println(listImageReferences(yamlPath))
					}
				}
				return nil
			},
		},
		{
			Name:    "push",
			Aliases: []string{"u"},
			Usage:   "tags images and pushes them to the corporate registry",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "path, p",
					Value:       ".",
					Usage:       "path of the kustomize file",
					Destination: &scanpath,
				},
				cli.StringFlag{
					Name:        "registry, r",
					Value:       "docker.io",
					Usage:       "name of the corporate registry to use",
					Destination: &registry,
				},
			},
			Action: func(c *cli.Context) error {
				paths, _ := listYamlFiles(scanpath)
				for _, yamlPath := range paths {
					imageRefs, _ := listImageReferences(yamlPath)
					if len(imageRefs) > 0 {
						fmt.Println(listImageReferences(yamlPath))
					}
				}
				return nil
			},
		},
		{
			Name:    "patch",
			Aliases: []string{"a"},
			Usage:   "patches the image references in all yaml files in the path",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "path, p",
					Value:       ".",
					Usage:       "path of the kustomize file and the yaml files",
					Destination: &scanpath,
				},
			},
			Action: func(c *cli.Context) error {
				paths, _ := listYamlFiles(scanpath)
				for _, yamlPath := range paths {
					imageRefs, _ := listImageReferences(yamlPath)
					if len(imageRefs) > 0 {
						fmt.Println(listImageReferences(yamlPath))
					}
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
