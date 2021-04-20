package glog

import "log"

func metalog(message ...interface{}) {
	if EnableMetaLogging {
		log.Println(message...)
	}
}
