# go_exam_2

## Step to build and binary file for test
1. cd go_exam_2/function
2. go test -c
3. run with ./function.test -test.v

```shell
# example result

$ ./function.test -test.v

=== RUN   TestValidateThailandIDLength
=== RUN   TestValidateThailandIDLength/Length_<_13
=== RUN   TestValidateThailandIDLength/Length_>_13
=== RUN   TestValidateThailandIDLength/Length_==_13
--- PASS: TestValidateThailandIDLength (0.00s)
    --- PASS: TestValidateThailandIDLength/Length_<_13 (0.00s)
    --- PASS: TestValidateThailandIDLength/Length_>_13 (0.00s)
    --- PASS: TestValidateThailandIDLength/Length_==_13 (0.00s)
=== RUN   TestValidateThailandIDNumericOnly
=== RUN   TestValidateThailandIDNumericOnly/Alphabet_Only
=== RUN   TestValidateThailandIDNumericOnly/Alphabet_Numeric
=== RUN   TestValidateThailandIDNumericOnly/Numeric_Only
--- PASS: TestValidateThailandIDNumericOnly (0.00s)
    --- PASS: TestValidateThailandIDNumericOnly/Alphabet_Only (0.00s)
    --- PASS: TestValidateThailandIDNumericOnly/Alphabet_Numeric (0.00s)
    --- PASS: TestValidateThailandIDNumericOnly/Numeric_Only (0.00s)
=== RUN   TestValidateThailandIDPattern
=== RUN   TestValidateThailandIDPattern/Logic_Incorrect
calculate identification id : 9999999999994
    input identification id : 9999999999999
=== RUN   TestValidateThailandIDPattern/Logic_Correct
calculate identification id : 3103109929515
    input identification id : 3103109929515
--- PASS: TestValidateThailandIDPattern (0.00s)
    --- PASS: TestValidateThailandIDPattern/Logic_Incorrect (0.00s)
    --- PASS: TestValidateThailandIDPattern/Logic_Correct (0.00s)
=== RUN   TestValidateThailandCitizenID
=== RUN   TestValidateThailandCitizenID/10_Digit
=== RUN   TestValidateThailandCitizenID/14_Digit
=== RUN   TestValidateThailandCitizenID/Text_and_Numeric
=== RUN   TestValidateThailandCitizenID/Text
=== RUN   TestValidateThailandCitizenID/13_Digit/_Correct_format#1
calculate identification id : 4953714647109
    input identification id : 4953714647109
=== RUN   TestValidateThailandCitizenID/13_Digit/_Correct_format#2
calculate identification id : 6226429459176
    input identification id : 6226429459176
--- PASS: TestValidateThailandCitizenID (0.00s)
    --- PASS: TestValidateThailandCitizenID/10_Digit (0.00s)
    --- PASS: TestValidateThailandCitizenID/14_Digit (0.00s)
    --- PASS: TestValidateThailandCitizenID/Text_and_Numeric (0.00s)
    --- PASS: TestValidateThailandCitizenID/Text (0.00s)
    --- PASS: TestValidateThailandCitizenID/13_Digit/_Correct_format#1 (0.00s)
    --- PASS: TestValidateThailandCitizenID/13_Digit/_Correct_format#2 (0.00s)
PASS
```

## Step to build go code
1. cd go_exam_2
2. go build .   
3. run with ./go_exam_2
4. Note that: Current platform arch. MacOS

```shell
# example result

$ ./go_exam_2
calculate identification id : 1516712728928
    input identification id : 1516712728928
validate thai citizen id success
```