// Cia jei kam reiketu slaptozodi uzsisifroti ir duombaze isirazyti
// seip nereikalingas sitas!
package main

import (
	"crypto/sha256"
	"encoding/hex"
	//"fmt"
)

//cia jei noretume susikurti daugiau useriu admin tai kad tinkamus duomenis i db irasyti per sita galim uzsifroti su atitinkamasi parametrais
/*func main() {
	slaptazodis := "Rastuvas123"
	druska := "asdf124!£!" // Jūsų pasirinkta druska

	hashedPassword := sukurtiHash(slaptazodis, druska)
	fmt.Print("asd")
	fmt.Println("Hashed password:", hashedPassword)
}*/

func sukurtiHash(slaptazodis, druska string) string {
	// Pridedame druską prie slaptažodžio
	sujungtasSlaptazodis := slaptazodis + druska

	// Sukuriame SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(sujungtasSlaptazodis))
	hashed := hasher.Sum(nil)

	// Konvertuojame į heksadecimales formatą
	hashedPassword := hex.EncodeToString(hashed)

	return hashedPassword
}
