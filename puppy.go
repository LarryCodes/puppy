package puppy

import (
  "github.com/LarryCodes/dog"
)

func Bark() string {
  return "Woof!"
}

func Barks() string {
  return "Woof! Woof! Woof!"
}

func BigBark() string {
  return dog.WhenGrownUp("Woof!")
}

func BigBarks() string {
  return dog.WhenGrownUp("Woof! Woof! Woof!")
}
