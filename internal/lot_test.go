package internal

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return strings.TrimSpace(buf.String())
}

func TestParkingLotBasicFlow(t *testing.T) {
	lot := NewParkingLot(2)

	output1 := captureOutput(func() {
		lot.Park("KA-01-HH-1234")
	})
	if !strings.Contains(output1, "Allocated slot number: 1") {
		t.Errorf("Expected allocation to slot 1, got: %s", output1)
	}

	output2 := captureOutput(func() {
		lot.Park("KA-01-HH-9999")
	})
	if !strings.Contains(output2, "Allocated slot number: 2") {
		t.Errorf("Expected allocation to slot 2, got: %s", output2)
	}

	output3 := captureOutput(func() {
		lot.Park("KA-01-BB-0001")
	})
	if !strings.Contains(output3, "Sorry, parking lot is full") {
		t.Errorf("Expected full lot message, got: %s", output3)
	}
}

func TestLeaveAndCharges(t *testing.T) {
	lot := NewParkingLot(1)
	lot.Park("KA-01-HH-1234")

	output := captureOutput(func() {
		lot.Leave("KA-01-HH-1234", 4)
	})
	if !strings.Contains(output, "Charge $30") {
		t.Errorf("Expected charge of $30, got: %s", output)
	}
}

func TestLeaveNonExistentCar(t *testing.T) {
	lot := NewParkingLot(1)

	output := captureOutput(func() {
		lot.Leave("DL-12-AA-9999", 2)
	})
	if !strings.Contains(output, "not found") {
		t.Errorf("Expected car not found message, got: %s", output)
	}
}

func TestStatus(t *testing.T) {
	lot := NewParkingLot(2)
	lot.Park("KA-01-HH-1234")

	output := captureOutput(func() {
		lot.Status()
	})

	if !strings.Contains(output, "Slot No. Registration No.") ||
		!strings.Contains(output, "1 KA-01-HH-1234") {
		t.Errorf("Unexpected status output:\n%s", output)
	}
}