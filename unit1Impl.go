package main

import (
	"github.com/skip2/go-qrcode"
	"github.com/ying32/govcl/vcl"
	"main/internal/base"
)

// ::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.Memo1.SetText("")
	f.Memo2.SetText("")
	f.Memo3.SetText("")
}

func (f *TForm1) OnPageControl1Change(sender vcl.IObject) {
	f.Memo1.SetText("")
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	input := f.Memo1.Text()
	output := base.Encode(input)
	f.Memo1.SetText(output)
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	input := f.Memo1.Text()
	err, output := base.Decode(input)
	if err != nil {
		f.Memo3.SetText(err.Error())
	}
	f.Memo1.SetText(output)
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	input := f.Memo2.Text()
	png, err := qrcode.Encode(input, qrcode.Low, 128)
	if err != nil {
		f.Memo3.SetText(err.Error())
	}
	f.Image1.Picture().LoadFromBytes(png)
}

func (f *TForm1) OnImage1Click(sender vcl.IObject) {

}
