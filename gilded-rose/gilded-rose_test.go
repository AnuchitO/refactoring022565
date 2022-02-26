package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

func Test_Foo(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{name: "+5 Dexterity Vest", sellIn: 10, quality: 20}}, outputs: []*Item{{name: "+5 Dexterity Vest", sellIn: 9, quality: 19}}},
		{name: "#2", inputs: []*Item{{name: "Aged Brie", sellIn: 2, quality: 0}}, outputs: []*Item{{name: "Aged Brie", sellIn: 1, quality: 1}}},
		{name: "#3", inputs: []*Item{{name: "Elixir of the Mongoose", sellIn: 5, quality: 7}}, outputs: []*Item{{name: "Elixir of the Mongoose", sellIn: 4, quality: 6}}},
		{name: "#4", inputs: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 80}}, outputs: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 80}}},
		{name: "#5", inputs: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 80}}, outputs: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 80}}},
		{name: "#6", inputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 15, quality: 1}}, outputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 14, quality: 2}}},
		{name: "#7", inputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 49}}, outputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 50}}},
		{name: "#8", inputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 49}}, outputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50}}},
		{name: "#9", inputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 48}}, outputs: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 50}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			UpdateQuality(c.inputs)

			if !reflect.DeepEqual(c.inputs, c.outputs) {
				pretty.Println(c.inputs, c.outputs)
				t.Errorf("Name: Expected %#v but got %#v ", c.outputs, c.inputs)
			}
		})
	}
}
