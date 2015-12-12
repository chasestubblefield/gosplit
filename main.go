package main

import "fmt"

func main() {
	testRandomStuff()
	// doFakeSplits()
}

func testRandomStuff() {
	inYAML :=
`game:
  name: 'The Legend of Zelda: Ocarina of Time'
  category: Any%
splits:
- name: Escape
  pb: 5m9.6687865s
  best_segment: 4m59.6438477s
- name: Kakariko
  pb: 6m28.0898338s
  best_segment: 1m5.2248149s
- name: Bottle
  pb: 8m37.0005333s
  best_segment: 2m8.9106995s
- name: Deku
  pb: 10m52.8010015s
  best_segment: 2m13.8767133s
- name: Scrubs
  pb: 12m22.9038024s
  best_segment: 1m27.9466357s
- name: Gohma
  pb: 13m24.3718981s
  best_segment: 48.7743584s
- name: Warp
  pb: 14m35.1288565s
  best_segment: 1m8.2604519s
- name: Collapse
  pb: 18m40.2788326s
  best_segment: 3m53.466028s
- name: Ganon
  pb: 22m21.0323713s
  best_segment: 3m39.0509277s
`
	inYAMLbytes := []byte(inYAML)
	splits, err := importFromYAML(inYAMLbytes)
	fmt.Printf("importFromYAML::\ninput:\n%voutput:\n%v, %v\n", inYAML, splits, err)

	outYAMLbytes, err := exportToYAML(splits)
	outYAML := string(outYAMLbytes)
	fmt.Printf("exportToYAML::\ninput:\n%v\noutput:\n%v%v\n", splits, outYAML, err)

	fmt.Printf("inYAML == outYAML: %v\n", inYAML == outYAML)
}
