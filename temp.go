FATIMA && REEM PLS IGNOREEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE
//package Methods

// import "fmt"

// const (
//     Reset = "\033[0m"
// )

// // PrintAscii prints a text-based representation of characters with optional coloring for a substring.
// func PrintAscii(argLettersDivided, fileLines []string, subString string, color *string) {
//     desiredColor := *color
//     tempColor := Reset
//     subStringLen := len(subString)
//     subStringIndex := 0
//     inSubString := false

//     for i := 0; i < len(argLettersDivided); i++ {
//         if argLettersDivided[i] == "\\n" {
//             fmt.Println()
//             continue
//         }

//         for j := 0; j < 8; j++ {
//             for k := 0; k < len(argLettersDivided[i]); k++ {
//                 start := ((int(argLettersDivided[i][k]) - 32) * 9) + (j + 1)

//                 // Check if we are in the middle of the substring
//                 if inSubString {
//                     if subStringIndex < subStringLen && argLettersDivided[i][k] == subString[subStringIndex] {
//                         fmt.Print(tempColor + fileLines[start])
//                         subStringIndex++
//                     } else {
//                         // If the current character does not match the expected substring character, reset
//                         inSubString = false
//                         subStringIndex = 0
//                         fmt.Print(Reset + fileLines[start])
//                     }
//                     // If we have reached the end of the substring, reset
//                     if subStringIndex == subStringLen {
//                         inSubString = false
//                         subStringIndex = 0
//                     }
//                 } else {
//                     if k <= len(argLettersDivided[i])-subStringLen && argLettersDivided[i][k:k+subStringLen] == subString {
//                         inSubString = true
//                         subStringIndex = 0
//                         tempColor = desiredColor
//                         fmt.Print(tempColor + fileLines[start])
//                     } else {
//                         fmt.Print(Reset + fileLines[start])
//                     }
//                 }
//             }
//             fmt.Print("\n")
//         }
//     }
//     // Make sure to reset color after printing
//     fmt.Print(Reset)
// }
