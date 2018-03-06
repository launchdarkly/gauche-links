package extension

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/mcuadros/go-crxmake"
)

const TemplateSuffix = ".tmpl"

type ExtensionConfig struct {
	DevHost string
	Host    string
	Icon    string
	Name    string
	Prefix  string
	Version string
}

var extensionBox packr.Box

func init() {
	extensionBox = packr.NewBox("./source")
}

func Build(config ExtensionConfig, targetDirPath string) error {
	b := crxmake.NewBuilder()

	tempDir, err := ioutil.TempDir("", "gauche-links-build-extension")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tempDir) }()
	if config.Icon != "" {
		sourceIconFile, iconErr := os.Open(config.Icon)
		if iconErr != nil {
			return fmt.Errorf(`unable to read icon at "%s": %s`, config.Icon, iconErr)
		}
		defer func() { _ = sourceIconFile.Close() }()
		iconBaseName := filepath.Base(config.Icon)
		iconFilePath := filepath.Join(tempDir, iconBaseName)
		iconErr = copyFile(iconFilePath, config.Icon)
		if iconErr != nil {
			return fmt.Errorf(`unable to copy icon file "%s" to "%s": %s`, config.Icon, iconFilePath, iconErr)
		}
		config.Icon = iconBaseName
	}

	err = extensionBox.Walk(func(name string, file packr.File) (err error) {
		targetPath := filepath.Join(tempDir, name)
		if strings.HasSuffix(name, TemplateSuffix) {
			targetPath = targetPath[:len(targetPath)-len(TemplateSuffix)]
		}
		outputFile, openErr := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, 0600)
		if openErr != nil {
			return openErr
		}
		defer func() {
			closeErr := outputFile.Close()
			if err != nil {
				err = closeErr
			}
		}()
		err = func() error {
			w := bufio.NewWriter(outputFile)
			if strings.HasSuffix(name, TemplateSuffix) {
				inputData, readErr := ioutil.ReadAll(file)
				if readErr != nil {
					return fmt.Errorf(`unable to read template for "%s": %s`, name, readErr)
				}
				tmpl, readErr := template.New(name).Parse(string(inputData))
				if readErr != nil {
					return fmt.Errorf(`unable to parse template for "%s": %s`, name, readErr)
				}
				if tmplErr := tmpl.Execute(w, config); tmplErr != nil {
					return fmt.Errorf(`unable to execute template for "%s": %s`, name, tmplErr)
				}
			} else if _, copyErr := io.Copy(w, file); copyErr != nil {
				return fmt.Errorf(`unable to copy file "%s": %s`, name, copyErr)
			}
			if flushErr := w.Flush(); flushErr != nil {
				return flushErr
			}
			return nil
		}()
		return
	})
	if err != nil {
		return fmt.Errorf(`unable to generate templates for chrome extension: %s`, err)
	}

	if zipErr := b.BuildZip(tempDir); zipErr != nil {
		return zipErr
	}

	chromeExtensionName := fmt.Sprintf("%s_gauche_links_extension_%s.crx", config.Prefix, config.Version)
	chromeExtension, err := os.Create(filepath.Join(targetDirPath, chromeExtensionName))
	if err != nil {
		return fmt.Errorf("unable to create chrome extension: %s", err)
	}

	defer func() { _ = chromeExtension.Close() }()

	err = b.WriteToFile(chromeExtension)
	if err != nil {
		return fmt.Errorf("unable to write chrome extension: %s", err)
	}

	firefoxExtensionName := fmt.Sprintf("%s_gauche_links_extension_%s.xpi", config.Prefix, config.Version)
	firefoxExtension, err := os.Create(filepath.Join(targetDirPath, firefoxExtensionName))
	if err != nil {
		return fmt.Errorf("unable to create firefox extension: %s", err)
	}
	archive := zip.NewWriter(bufio.NewWriter(firefoxExtension))
	defer func() { _ = archive.Close() }()

	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		writer, err := archive.Create(info.Name())
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() { _ = file.Close() }()

		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		return fmt.Errorf("unable to create firefox extension: %s", err)
	}
	return nil
}

func copyFile(dst string, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf(`unable to read file at "%s": %s`, src, err)
	}
	defer func() { _ = srcFile.Close() }()

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf(`unable to write file "%s": %s`, dst, err)
	}
	defer func() { _ = dstFile.Close() }()
	_, err = io.Copy(dstFile, srcFile)
	return err
}
