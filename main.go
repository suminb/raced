package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ParseOpts() {
	// Parse command line options
	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	fmt.Printf("word = %s\n", *wordPtr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Split sections by newline characters
	metadataSection := ReadAimCsvSection(scanner)
	headerSection := ReadAimCsvSection(scanner)
	recordSection := ReadAimCsvSection(scanner)

	// Parse the file
	// csvReader := csv.NewReader(stdinReader)
	csvReader := csv.NewReader(strings.NewReader(metadataSection))
	csvReader.FieldsPerRecord = -1
	ReadAimMetadata(csvReader)

	csvReader = csv.NewReader(strings.NewReader(headerSection))
	csvReader.FieldsPerRecord = -1
	ReadAimHeaders(csvReader)

	csvReader = csv.NewReader(strings.NewReader(recordSection))
	csvReader.FieldsPerRecord = -1
	ReadAimRecords(csvReader)
}

func ReadAimCsvSection(scanner *bufio.Scanner) string {
	builder := strings.Builder{}
	builder.Reset()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		builder.WriteString(line)
		builder.WriteByte('\n')
	}
	return builder.String()
}

// Reads in a CSV file exported from Race Studio, which is consist of three
// major sections: 1. file metadata, 2. headers, and 3. records.
func ReadAimCsv(reader *csv.Reader) {
	// ReadAimMetadata(reader)
	// ReadAimHeaders(reader)
}

func ReadAimMetadata(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) <= 0 {
			break
		}
		fmt.Printf("metadata: %s\n", record)
	}
}

func ReadAimHeaders(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) <= 1 {
			break
		}
		fmt.Printf("header: %s\n", record)
	}
}

func ReadAimRecords(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) <= 1 {
			break
		}
		fmt.Printf("record: %s\n", record)
	}
}
