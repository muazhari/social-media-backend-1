package test

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	chdirErr := os.Chdir("../.")
	if chdirErr != nil {
		t.Fatal(chdirErr)
	}
	fmt.Println("TestWeb started.")
	fmt.Println("TestWeb finished.")
}
