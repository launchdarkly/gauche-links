package extension

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
)

const TemplateSuffix = ".tmpl"

type Config struct {
	DevHost  string
	Host     string
	Icon     string
	Name     string
	Prefix   string
	Version  string
	Platform string
}

var extensionBox packr.Box

func init() {
	extensionBox = packr.NewBox("./source")
}

func Build(config Config, targetDirPath string) error {
	fileName := fmt.Sprintf("%s_gauche_links_extension_%s_%s.zip", config.Prefix, config.Platform, config.Version)
	path := filepath.Join(targetDirPath, fileName)
	archive, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(`unable to create archive at "%s": %s`, path, err)
	}

	zipFile := zip.NewWriter(archive)
	if config.Icon != "" {
		sourceIconFile, iconErr := os.Open(config.Icon)
		if iconErr != nil {
			return fmt.Errorf(`unable to read icon at "%s": %s`, config.Icon, iconErr)
		}
		defer func() { _ = sourceIconFile.Close() }()
		iconBaseName := filepath.Base(config.Icon)
		if err := archiveFile(zipFile, iconBaseName, sourceIconFile); err != nil {
			return err
		}
		config.Icon = iconBaseName
	}

	err = extensionBox.Walk(func(name string, file packr.File) (err error) {
		destName := name
		var r io.Reader = file
		if strings.HasSuffix(name, TemplateSuffix) {
			destName = destName[:len(destName)-len(TemplateSuffix)]
			inputData, err := ioutil.ReadAll(file)
			if err != nil {
				return fmt.Errorf(`unable to read template for "%s": %s`, name, err)
			}
			tmpl, err := template.New(name).Parse(string(inputData))
			if err != nil {
				return fmt.Errorf(`unable to parse template for "%s": %s`, name, err)
			}
			b := bytes.NewBuffer(nil)
			if tmplErr := tmpl.Execute(b, config); tmplErr != nil {
				return fmt.Errorf(`unable to execute template for "%s": %s`, name, tmplErr)
			}
			r = bytes.NewReader(b.Bytes())
		}
		return archiveFile(zipFile, destName, r)
	})
	if err != nil {
		return err
	}
	return zipFile.Close()
}

func archiveFile(writer *zip.Writer, name string, src io.Reader) error {
	w, err := writer.Create(name)
	if err != nil {
		return fmt.Errorf(`unable to create file named "%s": %s`, name, err)
	}
	_, err = io.Copy(w, src)
	if err != nil {
		return fmt.Errorf(`unable to write file named "%s": %s`, name, err)
	}
	return nil
}
