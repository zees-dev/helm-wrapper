package main

import (
	"log"

	"github.com/zees-dev/helm-wrapper/pkg/action"
)

func main() {
	log.Println("running helm list...")

	releaseStrings := action.GetReleaseStrings()
	for _, rs := range releaseStrings {
		log.Println(rs)
	}
}
