[![Build Status](https://travis-ci.org/chasetopher/gosplit.svg)](https://travis-ci.org/chasetopher/gosplit)

# GoSplit

Terminal-based speedrun timer written in [Go](https://golang.org).

## Installation

Ensure that Go is [installed](https://golang.org/doc/install) and `$GOPATH` is set in your profile.

```bash
go get -u github.com/chasetopher/gosplit
```

## Usage

```bash
$GOPATH/bin/gosplit
```

Or, if `$GOPATH/bin` is in your `$PATH`, simply:

```bash
gosplit
```

## Example `game.yml`

```yml
game:
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
```
