package manual_synunit

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ExeHdl struct {
	SourceDir  string
	AchieveDir string
}

func NewExeHdl(SourceDir, AchieveDir string) *ExeHdl {
	return &ExeHdl{SourceDir: SourceDir, AchieveDir: AchieveDir}
}

func (exeHdl *ExeHdl) Execute(handle func(val string)) error {
	// https://hakk.dev/docs/golang-sort-file-modified-time/

	// By default order by file name

	files, err := ioutil.ReadDir(exeHdl.SourceDir)
	if err != nil {
		return err
	}

	stgErrors := []string{}
	if len(files) == 0 {
		log.Printf("info: Don't have file for process \n")
		return nil
	}

	log.Printf("info: Begin proces file \n")
	for _, file := range files {
		sourceFiel := fmt.Sprintf("%v/%v", exeHdl.SourceDir, file.Name())
		err := exeHdl.readfile(sourceFiel, handle)
		if err != nil {
			stgErrors = append(stgErrors, err.Error())
		} else {
			err = exeHdl.moveToAchieve(true, file.Name())
			if err != nil {
				stgErrors = append(stgErrors, err.Error())
			}
		}
		// fmt.Println(file.Name(), file.Size(), file.ModTime())
	}

	if len(stgErrors) != 0 {
		return fmt.Errorf(strings.Join(stgErrors, "\n"))
	}
	log.Printf("info: End proces file \n")
	return nil

}

// type todo func(val string) error
type todo func(val string)

func (exeHdl *ExeHdl) readfile(filePath string, handle func(val string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		handle(string(line))
	}
}

func (exehdl *ExeHdl) moveToAchieve(isSuccess bool, fileName string) error {
	source := fmt.Sprintf("%v/%v", exehdl.SourceDir, fileName)
	archiev := ""
	if isSuccess {
		archiev = fmt.Sprintf("%v/success_%v", exehdl.AchieveDir, fileName)
	} else {
		archiev = fmt.Sprintf("%v/fail_%v", exehdl.AchieveDir, fileName)
	}

	err := os.Rename(source, archiev)
	if err != nil {
		return err
	}
	log.Printf("info: Move file %v to %v is success \n", fileName, archiev)

	return nil
}
