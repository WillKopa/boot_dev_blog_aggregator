package main

import (
	"fmt"
	"log"
	"github.com/WillKopa/boot_dev_blog_aggregator/internal/config"
)

func main() {
	gator_config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gator_config)
	gator_config.SetUser("William")

	gator_config, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gator_config)

	gator_config.SetUser("")

	gator_config, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gator_config)

}