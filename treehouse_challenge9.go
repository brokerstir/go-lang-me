package clock

type Hours int

func (h *Hours) Increment() {
  *h = (*h + 1) % 24
}
