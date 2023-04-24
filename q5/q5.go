package q5
import (
	"strings"
)
//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	dot := "."
	newS := strings.ToLower(s)

	newS2 := strings.Replace(newS, "a", "", 1)
	newS3 := strings.Replace(newS2, "e", "", 1)
	newS4 := strings.Replace(newS3, "i", "", 1)
	newS5 := strings.Replace(newS4, "o", "", 1)
	newS6 := strings.Replace(newS5, "u", "", 1)

	newS7 := dot + newS6

	return newS7

}
