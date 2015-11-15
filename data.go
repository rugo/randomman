package main

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)

/*
 * Holy moly, what a mess!
 */
func LoadAvailableManpages(basedir string) map[int][]string {
	availablePages := make(map[int][]string)

	sectionDirs, err := ioutil.ReadDir(basedir)
	if err != nil {
		log.Print(err)
	}
	for _, sectionDir := range sectionDirs {
		if sectionDir.IsDir() {
			sectionDirPath := path.Join(basedir, sectionDir.Name())
			files, err := ioutil.ReadDir(sectionDirPath)
			if err != nil {
				log.Print(err)
			}
			sectionNumber, err := strconv.Atoi(sectionDir.Name()[len(sectionDir.Name())-1:])
			if err != nil {
				log.Print(err)
			}
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
					if _, ok := availablePages[sectionNumber]; !ok {
						availablePages[sectionNumber] = []string{}
					}
					availablePages[sectionNumber] = append(
						availablePages[sectionNumber],
						path.Join(sectionDirPath, file.Name()))
				}
			}

		}
	}
	return availablePages
}

func GetRandomManpageFilename(sections []int, available map[int][]string) (string, error) {
	rand.Seed(time.Now().Unix())
	possibleSections := []int{}
	for _, section := range sections {
		if len(available[section]) > 0 {
			possibleSections = append(possibleSections, section)
		}
	}
	if len(possibleSections) == 0 {
		return "", errors.New("None of the given sections has to offer a man page")
	}
	rndSecIndex := rand.Intn(len(possibleSections)) + 1
	rndManPageFile := available[rndSecIndex][rand.Intn(len(available[rndSecIndex]))]
	return rndManPageFile, nil
}

func GetContentOfRandomManPage(section []int, available map[int][]string) (string, error) {
	fileName, err := GetRandomManpageFilename(section, available)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
