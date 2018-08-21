package clock

type Clock struct {
  Hours int
  Minutes int
}

// DEFINE A "Noon" FUNCTION HERE
func Noon() (Clock) {
  c := Clock{12, 0}
  return c
}