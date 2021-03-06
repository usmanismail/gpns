package gpnsconfig

import (
	"flag"
	"github.com/msbranco/goconfig"
	"log"
)

type APPLICATION_MODE int

const (
	SERVER_MODE   APPLICATION_MODE = iota
	REGISTER_MODE APPLICATION_MODE = iota
	SEND_MODE     APPLICATION_MODE = iota
)

func ParseConfig() APPLICATION_MODE {
	var aws_config_file string
	var base_config_file string
	var register bool
	var send bool
	var input_file string
	var output_file string
	var message_file string

	flag.StringVar(&base_config_file, "baseConfig", "./config/base.conf", "The path to the base configuration file")
	flag.StringVar(&aws_config_file, "awsConfig", "./config/aws.conf", "The path to the aws configuration file")

	flag.BoolVar(&register, "register", false, "Set flag to run in client mode and register a set of devices. If true inputFile and outputFile must be set.")
	flag.BoolVar(&send, "send", false, "Set flag to run in client mode and send push notifications to a set of arns. If true inputFile and outputFile must be set.")

	flag.StringVar(&input_file, "inputFile", "", "The path to the Device IDs or Arns file")
	flag.StringVar(&output_file, "outputFile", "", "The path to the Device IDs or Arns file")
	flag.StringVar(&message_file, "messageFile", "", "The path to the file containing the notificaito message to be sent out")

	flag.Parse()

	log.Printf("Using base configuration file: %s", base_config_file)
	baseConfig, err := goconfig.ReadConfigFile(base_config_file)
	checkError("Unable to parse base config", err)

	log.Printf("Using aws configuration file: %s", aws_config_file)
	awsConfig, err := goconfig.ReadConfigFile(aws_config_file)
	checkError("Unable to parse AWS config", err)

	parseBaseConfig(baseConfig)
	parseAwsConfig(awsConfig)

	if register {
		log.Printf("Running in client mode, registering devices listed in %s, and printing arns in %s", input_file, output_file)
		return REGISTER_MODE
	} else if send {
		log.Printf("Running in client mode, sending pusn notes to ARNs listed in %s, and printing results in %s", input_file, output_file)
		return SEND_MODE
	} else {
		log.Printf("Running in server mode")
		return SERVER_MODE
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
