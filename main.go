package main

import (
	"os"
	"os/exec"
	"path/filepath"	
	"log"
	"strings"
	"fmt"
)

func main(){
	currentDirectory, err := os.Getwd()
	currDir := getCurrDirName(currentDirectory)
	if err != nil {
		log.Fatal(err)
	}
	iterate(currentDirectory, currDir)
}

func getCurrDirName(path string) string{
	dirName := strings.Split(path, "/")
	return dirName[len(dirName)-1]
}

func iterate(path string, currDir string){
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			log.Fatalf(err.Error())
		}
		if(info.Name() != currDir){
			if(info.IsDir()){
				//Dentro de las carpetas de cada dia 
				cmd := exec.Command("/bin/sh", "-c", "(cd " + path + ";" + "../import-cv.sh lemus-maps.disca.upv.es)")
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatalf(err.Error())
				}
				fmt.Printf(string(out))
			}
		}
		return nil
	})
}
