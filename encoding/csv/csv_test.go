package csv

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	// 写到数据库
	csvFile, err := os.Create("tmp.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer csvFile.Close()

	// 定义 csv 数据
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"1", "4", "ns1"},
		{"2", "5", "ns2"},
		{"3", "6", "ns3"},
	}
	w := csv.NewWriter(csvFile)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			t.Fatalf("error writing record to csv: %s", err)
		}
	}
	//
	w.Flush()

	if err := w.Error(); err != nil {
		t.Fatal(err)
	}

	// 删除
	os.Remove("tmp.csv")
}

func TestReaderAll(t *testing.T) {
	csvFile, err := os.Open("testdata/tmp.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)

	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, record := range records {
		t.Log(record)
	}
}

func TestReader(t *testing.T) {
	csvFile, err := os.Open("testdata/tmp.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Log(record)
	}
}
