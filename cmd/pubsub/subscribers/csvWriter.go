package subscribers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

type CsvWriter struct {
	mutex     *sync.Mutex
	csvWriter *csv.Writer
	file      *os.File
}

func PrepareFile(filename string) *CsvWriter {
	csvFile, err := os.OpenFile("data/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(csvFile)
	return &CsvWriter{csvWriter: w, mutex: &sync.Mutex{}, file: csvFile}
}

func (w *CsvWriter) Write(row string) {
	w.mutex.Lock()
	err := w.csvWriter.Write([]string{row})
	fmt.Println(fmt.Sprintf("{ row: %s }", row))

	if err != nil {
		log.Fatal(err)
	}

	w.csvWriter.Flush()

	err = w.csvWriter.Error()

	if err != nil {
		log.Fatal(err)
	}

	w.csvWriter.Flush()
	_ = w.file.Close()
	w.mutex.Unlock()
}
