package main

import( "fmt" )

type person struct { 
    name string
}

type human interface { 
    saySomething()
}

func(p *person) saySomething() { 
    sentence :=  "Hallo" + p.name
    fmt.Println(sentence)
}

func speak(h human) { 
    h.saySomething()
}

func main() { 
    rafi := person{"Rafi"}

    speak(&rafi)
}
