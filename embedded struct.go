package main

import"fmt"
	type Eagle struct {
  Bird 
}
e := Eagle{name: "Baldie"}
e.makeSound()  
e.Bird.makeSound()
}