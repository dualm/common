package common

import (
	"testing"
)

func TestEquipment_Clear(t *testing.T) {
	eqp, err := NewEquipment("", "TEST", "-*")
	if err != nil {
		t.Fatal(err)
	}

	eqp.SetMachineRecipeName("recipe")

	if err := eqp.Clear(); err != nil {
		t.Fatal(err)
	}

	t.Log(eqp.MachineName, eqp.MachineRecipeName())
}
