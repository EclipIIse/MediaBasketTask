package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	root := "/Users/koi/Desktop/Work/WBTech"
	file, err := os.Create("result.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		errF := file.Close()
		if errF != nil {
			log.Println(errF)
		}
	}()

	var wg sync.WaitGroup

	err = catalogFileList(root, &wg, *file)

	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}

func catalogFileList(rootDir string, wg *sync.WaitGroup, file os.File) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err = file.WriteString(path + "\n")
				if err != nil {
					log.Println(err)
				}
			}()
		}
		return nil
	})
}
