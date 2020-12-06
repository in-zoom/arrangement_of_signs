package main

import (
	"fmt"
	"github.com/apaxa-go/eval"
	"reflect"
	"strings"
)

const desiredValue = 200

func main() {
	act := make(map[string]string)
	act["a"] = ""
	act["b"] = "+"
	act["c"] = "-"

	str := "9876543210"

	enumerationOfOptions(act, str)
}

func enumerationOfOptions(act map[string]string, str string) {

	numerics := strings.Split(str, "")

	for i, _ := range act {
		for j, _ := range act {
			for k, _ := range act {
				for l, _ := range act {
					for w, _ := range act {
						for u, _ := range act {
							for o, _ := range act {
								for s, _ := range act {
									for p, _ := range act {
										command := []string{i, j, k, l, w, u, o, s, p}
										line := strings.Join(command, "")
										if strings.Contains(line, "aaaa") && strings.Contains(line, "bbbb") && strings.Contains(line, "cccc") {
											continue
										}
										res := arrangementOfSigns(numerics, command)

										for key, _ := range act {
											res = strings.ReplaceAll(res, key, act[key])
										}
										result := calculationOfTheResult(res)
										if result == desiredValue {
											fmt.Println(res + "=200")
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func arrangementOfSigns(numerics, command []string) string {

	res := "" + numerics[0]
	for j, _ := range command {
		res = res + command[j] + numerics[j+1]
	}
	return res
}

func calculationOfTheResult(res string) int64 {

	expr, err := eval.ParseString(res, "")
	if err != nil {
		fmt.Println(err)
	}

	r, err := expr.EvalToInterface(nil)
	if err != nil {
		fmt.Println(err)
	}
	val := reflect.ValueOf(r)
	result := val.Int()
	return result
}
