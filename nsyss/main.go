package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
	"os/user"
)

func main() {
	parser := argparse.NewParser("nsyss", "New systemctl service")

	serviceName := parser.String("n", "name", &argparse.Options{Required: true, Help: "Name of new service"})
	script := parser.String("s", "script", &argparse.Options{Required: true, Help: "Script that will run on service start (must be bash script)"})
	description := parser.String("d", "description", &argparse.Options{Required: false, Help: "Service description", Default: ""})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf(parser.Usage(err))

		os.Exit(1)
	}

	if os.Getuid() != 0 {
		fmt.Printf("You need to run %v as root\n", parser.GetName())

		os.Exit(1)
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	file, err := os.Create(fmt.Sprintf("/lib/systemd/system/%v.service", *serviceName))
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	_, err = file.Write([]byte(fmt.Sprintf(
		"[Unit]\n" +
			"Description=%v\n" +
			"After=network.target\n" +
			"StartLimitIntervalSec=0\n\n" +
			"[Service]\n" +
			"Type=simple\n" +
			"Restart=always\n" +
			"RestartSec=1\n" +
			"User=%v\n" +
			"ExecStart=/bin/bash %v\n\n" +
			"[Install]\n" +
			"WantedBy=multi-currentUser.target\n",

		*description,
		currentUser.Username,
		*script,
	),
	))
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	err = file.Close()
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	fmt.Printf("%v -> /lib/systemd/system/%v.service\n", *serviceName, *serviceName)
}
