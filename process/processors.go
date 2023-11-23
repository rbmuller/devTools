package DevProcess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	php "github.com/kovetskiy/go-php-serialize"

	"github.com/devtools/models"
	null "gopkg.in/guregu/null.v3"
)

func ChunkArray[T any](items []T, chunkSize int) (chunks [][]T) {
	//Chunk an array for batch processing

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	chunked_data := append(chunks, items)

	return chunked_data
}

func CopyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func GenerateMap(data []byte) (models.Maps, error) {
	var dataMap models.Maps
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return dataMap, nil
}

func EpochToTimestamp(epochTime int) (string, bool) {

	if epochTime == 0 {
		return "", false
	}

	Epochint64 := int64(epochTime)

	// Convert milliseconds to seconds
	epochTimeSecs := Epochint64 / 1000

	// Convert the epoch time to a Go time object
	t := time.Unix(epochTimeSecs, 0)

	// Format the time as a string
	timestamp := t.Format("2006-01-02 15:04:05")

	if timestamp == "" {
		return "", false
	}

	return timestamp, true

}

func RemoveDuplicates(slice []string) []string {
	//Remove duplicates from a slice
	seen := make(map[string]bool)
	result := make([]string, 0)

	for _, element := range slice {
		if !seen[element] {
			seen[element] = true
			result = append(result, element)
		}
	}

	return result
}

func RemoveSliceDuplicates(slices [][]string) [][]string {
	seen := make(map[string]bool)
	result := make([][]string, 0)

	for _, slice := range slices {
		if !seen[slice[0]] {
			seen[slice[0]] = true
			result = append(result, slice)
		}

	}

	return result
}

func DaysDifference(dateStart, dateEnd time.Time) int {
	return int(math.Round(dateEnd.Sub(dateStart).Hours() / 24))
}

func AppendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func SelfGetByIndex(list []string, index int) string {
	if index >= len(list) {
		return ""
	}

	return list[index]
}

func MapArrayToString(strs []string) string {
	var result bytes.Buffer

	for _, str := range strs {
		result.WriteString(fmt.Sprintf(`%s,`, str))
	}

	result.Truncate(result.Len() - 1)

	return result.String()
}

func NumbersToString(nums []uint, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(nums), " "), delim), "[]")
}

func AnyToString(any interface{}) string {
	var s string

	kind := reflect.TypeOf(any).Kind()

	switch kind {
	case reflect.Float64:
		return fmt.Sprintf("%f", any)
	case reflect.Int:
		return strconv.Itoa(any.(int))
	default:
		return s
	}
}

func StringInArray(array []string, word string) bool {
	inArray := false

	for _, item := range array {
		if item == word {
			inArray = true
			break
		}
	}

	return inArray
}

func IntInArray(array []int, number int) bool {
	inArray := false

	for _, item := range array {
		if item == number {
			inArray = true
			break
		}
	}

	return inArray
}

func UintInArray(array []uint, number uint) bool {
	inArray := false

	for _, item := range array {
		if item == number {
			inArray = true
			break
		}
	}

	return inArray
}

func Floatify(f interface{}) null.Float {
	var floatValue null.Float
	switch i := f.(type) {
	case float64:
		floatValue = null.FloatFrom(i)
	case string:
		fl, _ := strconv.Atoi(i)
		floatValue = null.FloatFrom(float64(fl))
	default:
		floatValue = null.Float{}
	}
	return floatValue
}

func Integerify(f interface{}) null.Int {
	var intValue null.Int
	switch i := f.(type) {
	case float64:
		intValue = null.IntFrom(int64(i))
	case string:
		integer, _ := strconv.Atoi(i)
		intValue = null.IntFrom(int64(integer))
	default:
		intValue = null.Int{}
	}
	return intValue
}

func Boolify(f interface{}) null.Bool {
	var boolValue null.Bool
	if value, ok := f.(bool); ok {
		boolValue = null.BoolFrom(value)
	} else {
		boolValue = null.Bool{}
	}
	return boolValue
}

func Stringify(f interface{}) null.String {
	if f == nil {
		return null.String{}
	}
	var stringValue null.String
	if value, ok := f.(string); ok {
		stringValue = null.StringFrom(value)
	} else {
		stringValue = null.String{}
	}
	return stringValue
}

func TrimQuotes(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func UniqueStringArray(list []string) []string {
	unique := make([]string, 0)

	for _, item := range list {
		if !StringInArray(unique, item) {
			unique = append(unique, item)
		}
	}

	return unique
}

func StringifyWithQuotes(items []string) string {
	if len(items) == 0 {
		return ""
	}
	var result bytes.Buffer

	for _, item := range items {
		result.Write([]byte(fmt.Sprintf("'%s',", item)))
	}

	result.Truncate(result.Len() - 1)

	return result.String()
}

func CompleteDateWithDays(dateWithoutDay string) string {
	return fmt.Sprintf("%s01", dateWithoutDay)
}

func SplitByRegexp(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)

	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}

	result[len(indexes)] = text[laststart:]

	return result
}

func MillisecondsToTime(millis int) time.Time {
	return time.Unix(0, int64(millis*int(time.Millisecond)))
}

func Unserialize(valueInBytes []byte) ([]int, error) {
	decoded, err := php.Decode(string(valueInBytes))
	if err != nil {
		return nil, err
	}

	array := make([]int, 0)

	for _, id := range decoded.(map[interface{}]interface{}) {
		if id != nil {
			//nolint
			switch id.(interface{}).(type) {
			case string:
				intID, _ := strconv.Atoi(id.(string))
				array = append(array, intID)
			case int64:
				array = append(array, int(id.(int64)))
			case int:
				array = append(array, id.(int))
			default:
				return nil, fmt.Errorf("Unsupported type from serialized value")
			}
		}
	}

	return array, nil
}

func StringToDate(layout string, dateStr string) (null.Time, error) {
	if dateStr == "" {
		return null.Time{}, nil
	}

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return null.Time{}, err
	}

	return null.TimeFrom(date), nil
}

func Divide(left, right float64) float64 {
	if right == 0 {
		return 0
	}
	return left / right
}

func SplitEOLDataByTokens(dateStr null.String) (day, month, year string) {
	if !dateStr.Valid {
		return "0", "0", "0"
	}
	dateTokens := strings.Split(dateStr.String, "/")
	result := make([]string, 0, len(dateTokens))

	for _, token := range dateTokens {
		if len(token) == 1 {
			result = append(result, fmt.Sprintf("0%s", token))
		} else {
			result = append(result, token)
		}
	}

	return result[1], result[0], result[2]
}

func CompleteWithLeadingZeros(dateStr null.String) null.String {
	if !dateStr.Valid {
		return null.String{}
	}

	day, month, year := SplitEOLDataByTokens(dateStr)

	finalDateString := null.StringFrom(
		fmt.Sprintf("%s/%s/%s", day, month, year),
	)

	return finalDateString
}
