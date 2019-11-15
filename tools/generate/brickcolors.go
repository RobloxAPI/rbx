package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/anaminus/but"
)

type Root struct {
	BrickColors []BrickColor
	Default     uint32
	Palette     []uint32
}

type BrickColor struct {
	Number      uint32
	IndexColor  uint32
	IndexColor3 uint32
	IndexName   uint32
	IndexNumber uint32
	Color       [3]float32
	Color8      [3]uint8
	Name        string
}

func GenerateBrickColors(input, output string) {
	b, err := ioutil.ReadFile(input)
	but.IfFatal(err, "read file")
	var root Root
	but.IfFatal(json.Unmarshal(b, &root), "decode json")

	var buf strings.Builder
	buf.WriteString("// File was generated automatically. DO NOT EDIT!\n\n")
	buf.WriteString("package types\n")
	{
		buf.WriteString("\n")
		fmt.Fprintf(&buf, "const BrickColorPaletteSize = %d\n", len(root.Palette))
		fmt.Fprintf(&buf, "const BrickColorIndexSize = %d\n", len(root.BrickColors))
		fmt.Fprintf(&buf, "const BrickColorDefault BrickColor = %d\n", root.Default)
		var idx int
		for i, bc := range root.BrickColors {
			if bc.Number == root.Default {
				idx = i
				break
			}
		}
		fmt.Fprintf(&buf, "const bcDefaultIndex = %d\n", idx)
	}
	{
		buf.WriteString("\nvar brickColors = [BrickColorIndexSize]BrickColor{")
		wrapList(&buf, 8, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			fmt.Fprintf(buf, "%d", root.BrickColors[i].Number)
			return true
		})
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nvar bcPalette = [BrickColorPaletteSize]BrickColor{")
		wrapList(&buf, 8, len(root.Palette), func(buf *strings.Builder, i int) bool {
			fmt.Fprintf(buf, "%d", root.Palette[i])
			return true
		})
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nvar bcColors = [BrickColorIndexSize]Color3{")
		wrapList(&buf, 0, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			bc := root.BrickColors[i]
			fmt.Fprintf(buf, "{R: %d.0 / 255, G: %d.0 / 255, B: %d.0 / 255}",
				bc.Color8[0],
				bc.Color8[1],
				bc.Color8[2],
			)
			return true
		})
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nvar bcColors8 = [BrickColorIndexSize][3]int{")
		wrapList(&buf, 0, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			bc := root.BrickColors[i]
			fmt.Fprintf(buf, "{%d, %d, %d}",
				bc.Color8[0],
				bc.Color8[1],
				bc.Color8[2],
			)
			return true
		})
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nvar bcIndex = map[BrickColor]int{")
		wrapList(&buf, 6, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			fmt.Fprintf(buf, "%d: %d", root.BrickColors[i].Number, i)
			return true
		})
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nconst bcNames = ")
		var b []byte
		for _, bc := range root.BrickColors {
			b = append(b, []byte(bc.Name)...)
		}
		buf.WriteString(wrapString(72, 1, "", []byte(b)))
	}
	var namesIndex []int
	var idxlen int
	{
		buf.WriteString("\nvar bcNamesIndex = [BrickColorIndexSize + 1]int{")
		n := 0
		wrapList(&buf, 8, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			bc := root.BrickColors[i]
			fmt.Fprintf(buf, "%d", n)
			namesIndex = append(namesIndex, n)
			n += len(bc.Name)
			return true
		})
		fmt.Fprintf(&buf, "\t%d,\n", n)
		namesIndex = append(namesIndex, n)
		idxlen = len(fmt.Sprintf("%d", n))
		buf.WriteString("}\n")
	}
	{
		buf.WriteString("\nvar bcNameIndex = map[string]BrickColor{")
		wrapList(&buf, 0, len(root.BrickColors), func(buf *strings.Builder, i int) bool {
			bc := root.BrickColors[i]
			if bc.IndexName != bc.Number {
				return false
			}
			key := fmt.Sprintf("bcNames[%d:%d]:", namesIndex[i], namesIndex[i+1])
			fmt.Fprintf(buf, "%*s %d", -idxlen-idxlen-11, key, bc.Number)
			return true
		})
		buf.WriteString("}\n")
	}
	but.IfFatal(ioutil.WriteFile(output, []byte(buf.String()), 0666), "write file")
}
