# Statistical Analysis Program

This performs statistical analysis on a dataset of numbers provided int the text file. It calculates the average, median, variance, and standard deviation of the numbers and prints out the results

## Prerequisites

- [Go Programming Language](https://golang.org/dl/) installed on your system. 

## Installation

Run these commands on your terminal window to clone the repository and install the program 
```
git clone https://learn.zone01kisumu.ke/git/aadero/math-skills.git

cd math-skills
```

## Usage.
To use this program, follow these stesps:
1. Place your dataset in a text file (data.txt) each number is on a separate line. 
For example:

```shell
189
113
121
114
145
110
...
```
2. Use the terminal (command line) to navigate the folder whre you Go prgram and data file is located.

3. Run your program with the following command, replacing `data.txt` with the path to your data file:

```shell
go run . data.txt
```
Make sure you replace `data.txt` with the actual filename of the data set  if it's different. 

## Expected Output.

The program will output the calculations as rounded integers in the terminal:

```
Average: 35 
Median: 4 
Variance:5  
Standard Deviation: 65
```
*Note*: The Numbers shown above are only an example. Your results will vary depending on the data contained in your file. 

## Troubleshooting
- Ensure tha you're in the correct directory where your Go Program exists. 
-Double-check that `data.txt` exists in the same directory and is corectly formarted with numbers, one per line. 
- If you  encounter permission issues, make sure the Go program file is executable. You may set this with `chmod +x your-program.go` on Unix systems.

## Author

[mvuvi](https://learn.zone01kisumu.ke/git/aadero/) - Zone01 Kisumu Aspiring developer. 

## Acknowledgements

I would like to Thank Zone01 and Cohort-1 peers that helped me learn and inspired me to create this program.
