package models

import "time"

type CommandLineArgument struct {
	BaseSystemDirector string
	ApiUrl             string
	TimeBetweenRequest time.Duration
}
