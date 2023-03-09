package main

import (
	"log"
	"os"
)

var (
	poderRootDir      = "/var/lib/poder"
	poderTmpImageDir  = "/var/lib/poder/tmp"
	poderImageDir     = "/var/lib/poder/images"
	poderContainerDir = "/var/run/poder/containers"
)

// poder 启动时预先创建使用过程中需要用到的文件夹
func createRequiredPoderDirs() (err error) {
	dirs := []string{poderRootDir, poderTmpImageDir, poderImageDir, poderContainerDir}

	return createDirIfNotExist(dirs)
}

func createDirIfNotExist(dirs []string) error {
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				log.Printf("Error creating directory: %v\n", err)
				return err
			}
		}
	}
	return nil
}
