package avax_utils

import (
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"log"
	"testing"
)

func TestDecodeUtxo(t *testing.T) {
	utxoBytes, err := formatting.Decode(formatting.CB58, "11PQ1sNw9tcXjVki7261souJnr1TPFrdVCu5JGZC7Shedq3a7xvnTXkBQ162qMYxoerMdwzCM2iM1wEQPwTxZbtkPASf2tWvddnsxPEYndVSxLv8PDFMwBGp6UoL35gd9MQW3UitpfmFsLnAUCSAZHWCgqft2iHKnKRQRz")
	if err != nil {
		t.Fatal("formatting.Decode error", err)
	}

	utxo := avax.UTXO{}
	_, err = platformvm.Codec.Unmarshal(utxoBytes, &utxo)
	
	if err != nil {
		t.Fatal(err)
	}
	t.Log("utxo:", utxo)

	txid := utxo.TxID
	t.Log("txid:", txid)

	outputIndex := utxo.OutputIndex
	t.Log("outputIndex:", outputIndex)

	log.Printf("%+v", utxo.Out)
}
