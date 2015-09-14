package main

import (
	"testing"
	"log"
)

func TestExample(t *testing.T){

	val, err := example();

	if err != nil {
		t.Error("error");
	}

	log.Println(val);
}