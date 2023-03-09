# bubble-plot

**very** experimental [Bubbletea](https://github.com/charmbracelet/bubbletea) custom bubble for displaying plots using braille characters and my own fork of [exrook/drawille-go](https://github.com/chriskim06/drawille-go) which takes a lot of inspiration from [termui's implementation](https://github.com/gizak/termui/blob/master/drawille/drawille.go)

## notes

- the font this is displayed on needs to be able to render unicode braille characters (0x28**) and line characters (0x25**) for the x and y axis
- displaying horizontal labels is hard, right now this just tries to display a label at the beginning and end of the graphed lines
- this and my fork of drawille are very much a work in progress
