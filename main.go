package main

import (
	"fmt"
	"strconv"
)

func main() {
	dir := [][]string{{"/", "-1", "a", "1", "b", "2", "A", "0"}, {"a", "0", "c", "3", "B", "0"}, {"b", "0", "C", "0"}, {"c", "1", "C", "0"}}
	var index = 0

	cmd := ""

	dirName := ""

	for {
		fmt.Println("Please any comand")
		fmt.Scanf("%s", &cmd)

		if cmd == "ls" {
			ls(dir[index])
		} else if cmd == "pwd" {
			pwd(dir, index)
		} else if cmd == "cd" {
			fmt.Scanf("%s", &dirName)

			if i := cd(dir[index], dirName); i == -1 {
				fmt.Println(dirName, "not found.")
			} else if i == -2 {
				fmt.Println(dirName, "is file.")
			} else {
				index = i
			}
			dirName = ""
		} else if cmd == "exit" {
			break
		} else {
			fmt.Println(cmd, "command not found.")
		}
	}
}

func pwd(dir [][]string, index int) {
	nextIndex := index
	pwdArray := []int{}
	for {
		if nextIndex == -1 {
			break
		}
		pwdArray = append(pwdArray, nextIndex)
		nextIndex, _ = strconv.Atoi(dir[nextIndex][1])
	}
	for i := 0; i < len(pwdArray); i++ {
		fmt.Printf("%s", dir[pwdArray[len(pwdArray)-1-i]][0])
		if i != 0 && i != len(pwdArray)-1 {
			fmt.Printf("/")
		}
	}
	fmt.Println("")
}

func ls(files []string) {
	for i, fileName := range files {
		if i%2 == 0 && i != 0 {
			if i != 2 {
				fmt.Printf(",")
			}
			fmt.Printf("%s", fileName)
		}
	}
	fmt.Println("")
}

func cd(files []string, fileName string) int {
	var nextDirIndex int = -1
	if fileName == "" {
		nextDirIndex = 0
	} else if fileName == ".." {
		nextDirIndex, _ = strconv.Atoi(files[1])
	} else {
		for i, file := range files {
			if i%2 == 0 && i != 0 {
				if file == fileName {
					if i, _ := strconv.Atoi(files[i+1]); i == 0 {
						nextDirIndex = -2
					} else {
						nextDirIndex = i
					}
					break
				}
			}
		}
	}
	return nextDirIndex
}
