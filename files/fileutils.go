package files

import (
	"os"
	"log"
	"fmt"
)

/*
1a
Les text1.txt og text2.txt vha. fileutils inn i “byte slices” og sammenlign disse.
Filene har tilsynelatende lik innhold, men hvorfor er den ene filen større enn den
andre? Bortsett fra forskjeller i størrelsen, hvorfor må man være oppmerksom på
dette “fenomenet”?

1b
Skriv en funksjon, som finner ut hvilken kode for linjeskift en tekst-fil bruker. Dere
skal selv finne eget navn for deres funksjoner, men programmet skal hete
“lineshift.go”. Programmet skal kunne bli utført fra kommandolinje og returnere en
eller annen type konklusjon. Programmet skal ta et <filnavn> som in-data.

 */




func ReadFile(filename string) {
	// Read a file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Put file into a byte slice variable.
	fileByte := make([]byte, 1024)
	// read the file for the number of bytes specified in fileByte.
	readFile, err := file.Read(fileByte)
	if err != nil {
		log.Fatal(err)
	}
	// print the result.
	fmt.Printf("%d bytes: %\n", readFile, string(fileByte))
}



