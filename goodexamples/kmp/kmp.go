/*
Алгоритм Кнута Морриса Прата по нахождению в строке подстроки
*/

package kmp

var pi []int

//Преобразуем строку в массив рун
func toRune(str string) []rune {
	return []rune(str)
}

//Ищем вхождение подстроки в строку
func FindRepeateds(subStr, str string) int {
	//Количество найденых совпадений в подстроке
	var count int
	//Массив значений префиксов и суффиксов в строке,
	//которую надо найти в основной строке
	pi := CreatePi(subStr)
	//Индекс для прохождения по массиву рун подстроки
	var iSubStr int
	//Индекс для массиву рун строки
	var iStr int

	strRune := toRune(str)
	subStrRune := toRune(subStr)

	for {
		if iStr == len(strRune) {
			break
		}
		switch {
		case subStrRune[iSubStr] == strRune[iStr]:
			if iSubStr == len(subStrRune)-1 {
				iSubStr = -1
				count++
			}
			iSubStr++
			iStr++
		case iSubStr == 0:
			iStr++
		default:
			iSubStr = pi[iSubStr-1]
		}
	}
	return count
}

//Создаем массив Pi для определния индека для подстроки
func CreatePi(str string) []int {
	var i, iPrefics int
	strRune := toRune(str)
	pi = make([]int, len(strRune))
	i = 1
	for {
		if i == len(strRune) {
			break
		}
		switch {
		case strRune[i] == strRune[iPrefics]:
			pi[i] = iPrefics + 1
			iPrefics++
			i++
		case iPrefics == 0:
			pi[i] = 0
			i++
		default:
			iPrefics = pi[iPrefics-1]
		}

	}
	return pi
}
