package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New() // created new app
	w := a.NewWindow("Calculator")// created new window in app, the header is Hello
	//create buttons for calculator
	input:=""
	output:=widget.NewLabel(input)
	backButton:=widget.NewButton("Back", func(){
		if len(input)>0{
		input=input[0:len(input)-1]; //traverse from 0 to len -1
		output.SetText(input)
		}
	});
	
	clearButton:=widget.NewButton("Clear", func(){
		input = ""
		output.SetText(input)
	});
	openBrac:=widget.NewButton("(", func(){
		input = input+"("
		output.SetText(input)
	});
	closeBrac:=widget.NewButton("(", func(){
		input = input+"("
		output.SetText(input)
	});
	divideButton:=widget.NewButton(")", func(){
		input = input+")"
		output.SetText(input)
	});

	sevenButton:=widget.NewButton("7", func(){
		input = input+"7"
		output.SetText(input)
	});
	eightButton:=widget.NewButton("8", func(){
		input = input+"8"
		output.SetText(input)
	});
	nineButton:=widget.NewButton("9", func(){
		input = input+"9"
		output.SetText(input)
	});
	multiplyButton:=widget.NewButton("*", func(){
		input = input+"*"
		output.SetText(input)
	});

	fourButton:=widget.NewButton("4", func(){
		input = input+"4"
		output.SetText(input)
	});
	fiveButton:=widget.NewButton("5", func(){
		input = input+"5"
		output.SetText(input)
	});
	sixButton:=widget.NewButton("6", func(){
		input = input+"6"
		output.SetText(input)
	});
	minusButton:=widget.NewButton("-", func(){
		input = input+"-"
		output.SetText(input)
	});

	oneButton:=widget.NewButton("1", func(){
		input = input+"1"
		output.SetText(input)
	});
	twoButton:=widget.NewButton("2", func(){
		input = input+"2"
		output.SetText(input)
	});
	threeButton:=widget.NewButton("3", func(){
		input = input+"3"
		output.SetText(input)
	});
	plusButton:=widget.NewButton("+", func(){
		input = input+"+"
		output.SetText(input)
	});

	zeroButton:=widget.NewButton("0", func(){
		input = input+"0"
		output.SetText(input)});
	decimalButton:=widget.NewButton(".", func(){
		input = input+"."
		output.SetText(input)
	});
	//Took ref from : https://github.com/Knetic/govaluate For evaluate functionality. First go get github.com/Knetic/govaluate in terminal to download the packages and dependency from the given path
	equalButton:=widget.NewButton("=", func(){
		expression, err := govaluate.NewEvaluableExpression(input);
		if err == nil {
			result, err := expression.Evaluate(nil); // we have taken as nil as here we dont want to map anything. For ex: 2+3, here we dont need to map anything, just work with the input. 2+3+a+b, here we would have required to map a,b.
				if err == nil {
					//FormatFloat converts the floating-point number to a string, according to the format fmt and precision prec. It rounds the result assuming that the original was obtained from a floating-point value of bitSize bits (32 for float32, 64 for float64).
					
					//float64 is a version of float that stores decimal values using a total of 64-bits of data.
					input= strconv.FormatFloat(result.(float64),'f', -1, 64) // formats the result which is in string to float. Format in ->'f' (-ddd.dddd, no exponent), The special precision -1 uses the smallest number of digits necessary such that ParseFloat will return f exactly. ref : https://pkg.go.dev/strconv#FormatFloat
				}else{
					input="error"
					}
		}else
		{
			input = "error"
		}
   		output.SetText(input)
		
	});
	w.SetContent(container.NewVBox(
		output,
		container.NewGridWithColumns(1, //considering the part below the result as seperate container that has coloumns in it. Considering the below section as one column container which then has sub columns
			container.NewGridWithColumns(1,// this container has 2 columns, history and clear.
				backButton,
			),
			container.NewGridWithColumns(4,// this container has 4 columns.
				clearButton,
				openBrac,
				closeBrac,
				divideButton,
			),
			container.NewGridWithColumns(4,// this container has 4 columns.
				sevenButton,
				eightButton,
				nineButton,
				multiplyButton,
			),
			container.NewGridWithColumns(4,// this container has 4 columns.
				fourButton,
				fiveButton,
				sixButton,
				minusButton,
			),
			container.NewGridWithColumns(4,// this container has 4 columns.
				oneButton,
				twoButton,
				threeButton,
				plusButton,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroButton,
					decimalButton,
				),
				equalButton,),),
	
	
	),)
	
	
	w.ShowAndRun()
	
	//main end
	}
	

