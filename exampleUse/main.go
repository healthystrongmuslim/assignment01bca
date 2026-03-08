package main

import (
	"crypto/rand"
	"image/color"
	"math/big"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/healthystrongmuslim/assignment01bca"
)

func BlockWidget(b *assignment01bca.Block) *fyne.Container {
	transactionEntry := widget.NewEntry()
	transactionEntry.SetText(b.Transaction)
	transactionEntry.Resize(fyne.NewSize(200, transactionEntry.MinSize().Height))
	nonceEntry := widget.NewEntry()
	nonceEntry.SetText(strconv.Itoa(b.Nonce))
	nonceEntry.Resize(fyne.NewSize(300, nonceEntry.MinSize().Height))

	hashLabel := widget.NewLabelWithStyle(b.Hash, fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	hashLabel.Resize(fyne.NewSize(10, hashLabel.MinSize().Height*2))
	prevHashLabel := widget.NewLabelWithStyle(b.PrevHash, fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	prevHashLabel.Resize(fyne.NewSize(10, prevHashLabel.MinSize().Height))

	indexLabel := widget.NewLabelWithStyle(strconv.Itoa(b.Index), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	timeLabel := widget.NewLabelWithStyle(strconv.FormatInt(b.CreationTime.Unix(), 16), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	merkleLabel := widget.NewLabelWithStyle(b.MerkleRoot, fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	merkleLabel.Resize(fyne.NewSize(10, merkleLabel.MinSize().Height))
	updateBlock := func() {
		b.Transaction = transactionEntry.Text
		b.Nonce, _ = strconv.Atoi(nonceEntry.Text)
		b.MerkleRoot = assignment01bca.CalculateHash(b.Transaction)
		header := b.PrevHash + strconv.Itoa(b.Nonce) + b.MerkleRoot + strconv.FormatInt(b.CreationTime.Unix(), 10) + strconv.Itoa(b.Index)
		b.Hash = assignment01bca.CalculateHash(header)
		hashLabel.SetText(b.Hash)
		merkleLabel.SetText(b.MerkleRoot)
	}

	transactionEntry.OnSubmitted = func(string) { updateBlock() }
	nonceEntry.OnSubmitted = func(string) { updateBlock() }
	updateBtn := widget.NewButton("Update Block", updateBlock)
	r0 := container.NewHBox(
		widget.NewLabelWithStyle("Prev", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), prevHashLabel,
	)
	r1 := container.NewHBox(
		widget.NewLabelWithStyle("TID", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), indexLabel,
		widget.NewLabelWithStyle("Nonce", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), nonceEntry,
		widget.NewLabelWithStyle("Tsx", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), transactionEntry,
		widget.NewLabelWithStyle("TTime", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), timeLabel,
		updateBtn,
	)

	// Second row: merkle, hash, time, update button
	r2 := container.NewHBox(widget.NewLabelWithStyle("MRoot", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), merkleLabel)
	r3 := container.NewHBox(widget.NewLabelWithStyle("Hash", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true}), hashLabel)
	bgcolor := color.RGBA{50, 50, 50, 50}
	if b.PrevBlock != nil && b.PrevHash != b.PrevBlock.Hash {
		bgcolor = color.RGBA{150, 100, 100, 255}
	}
	bgRect := canvas.NewRectangle(bgcolor)
	blockVBox := container.NewVBox(r0, r1, r2, r3)
	return container.NewStack(bgRect, blockVBox)
}

func VerifyChain(w fyne.Window, blockchain []*assignment01bca.Block) {
	w.SetContent(ChainWidget(w, blockchain))
}

func AddBlock(w fyne.Window, blockchain []*assignment01bca.Block) {
	prev := blockchain[len(blockchain)-1]
	ranNum, _ := rand.Int(rand.Reader, big.NewInt(10e7))
	newBlock := assignment01bca.NewBlock("New Tx", int(ranNum.Int64()), prev, prev.Index+1)
	blockchain = append(blockchain, newBlock)
	VerifyChain(w, blockchain) // Refresh UI
}

func ChainWidget(w fyne.Window, blockchain []*assignment01bca.Block) *container.Scroll {
	blocks := []fyne.CanvasObject{}
	for _, b := range blockchain {
		blocks = append(blocks, BlockWidget(b))
	}
	addBtn := widget.NewButton("Add Block", func() { AddBlock(w, blockchain) })
	verifyBtn := widget.NewButton("Verify Blockchain (only match hash history)", func() { VerifyChain(w, blockchain) })
	chainBtns := container.NewHBox(addBtn, verifyBtn)
	blocks = append(blocks, chainBtns)
	return container.NewVScroll(container.NewVBox(blocks...))
}

var bchain []*assignment01bca.Block

func main() {
	a := app.New()
	w := a.NewWindow("Blockchain")
	onTypedKey := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyA {
			AddBlock(w, bchain)
		}
		if e.Name == fyne.KeyV {
			VerifyChain(w, bchain)
		}
	}
	w.Canvas().SetOnTypedKey(onTypedKey)
	// Start with genesis block
	bchain = []*assignment01bca.Block{
		assignment01bca.NewBlock("Genesis 5:60", 8, nil, 0),
	}
	bchain = append(bchain, assignment01bca.NewBlock("Asim 10000 BTC → Musab", 311, bchain[0], 1))
	bchain = append(bchain, assignment01bca.NewBlock("Musab 1 PZA → Asim", 32, bchain[1], 2))

	w.SetContent(ChainWidget(w, bchain))
	w.ShowAndRun()
}
