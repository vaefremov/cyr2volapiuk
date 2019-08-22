package cyr2volapiuk_test

import (
	"fmt"

	"github.com/vaefremov/cyr2volapiuk"
)

func ExampleFileName() {
	fileNameIn := "Это имя файла по-русски?!.hd5"
	fileNameOut := cyr2volapiuk.FileName(fileNameIn)
	fmt.Println(fileNameIn, fileNameOut)
	// Output: Это имя файла по-русски?!.hd5 Eto_imya_fajla_po-russki__.hd5
}
