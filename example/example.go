package main

import (
	"fmt"
	"log"

	"github.com/k-kond/bp3d"
)

func main() {
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Big Bin", 500, 330, 240, 10000))
	//p.AddBin(bp3d.NewBin("Medium Bin", 100, 150, 200, 1000))
	//p.AddBin(bp3d.NewBin("Small Bin", 10, 15, 20, 100))
	
	// Add items.
	// {"volume":0.00143,"length":0.2,"width":0.084,"weight":0.892,"height":0.084}
	// {"volume":0.00143,"length":0.2,"width":0.084,"weight":0.892,"height":0.084}
	p.AddItem(bp3d.NewItem("Шампунь Dove Питающий уход 250 мл", 100, 150, 300, 892))
	p.AddItem(bp3d.NewItem("Маска тканевая для лица Dove успокаивающая", 84, 84, 200, 300))
	p.AddItem(bp3d.NewItem("Маска тканевая для лица Dove успокаивающая", 84, 84, 200, 300))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}

	// Will output:
	//
	// Small Bin(10x15x20, max_weight:100)
	//  packed items:
	//    Item 2(3x3x2, weight: 3) pos(0,0,0) rt(RotationType_WHD (w,h,d))
	//    Item 1(2x2x1, weight: 2) pos(3,0,0) rt(RotationType_WHD (w,h,d))
	// Medium Bin(100x150x200, max_weight:1000)
	//  packed items:
	displayPacked(p.Bins)
}

func displayPacked(bins []*bp3d.Bin) {
	for _, b := range bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
	}
}
