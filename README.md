Calculation of the number of lines and functions in repositories,
for the golang programming language.

The go_lines_count console utility is designed to calculate and display
number of lines and functions in repositories (directories),
to study the source code.
Displayed:
- names of directories and subdirectories
- number of functions and lines in the catalog

A sample implementation can be found in the examples directory,
example:
```
Name	                Level	Lines count	Functions count
sanek           	1	423385	15884
calendar	        2	674	22
casebook_service	2	7476	46
```

Installation procedure:
1. Compile this repository
>make build
>
the go_lines_count file will appear in the "bin" folder

3. Run the go_lines_count file with the following parameters:
go_lines_count <DIRECTORY_SOURCE> <FILENAME> <FOLDERS_LEVEL>
startup example:
>./go_lines_count ./ ./lines_count.txt 2
>
(or fill out the file bin/settings.txt)

4. After launch, a new filled file will be created or text will be displayed in the console.
```
Settings:
1. You can run it without any settings, there will be default settings.
2. DIRECTORY_SOURCE
the folder where the source code in golang is located,
Calculation will begin from this folder, taking into account subfolders.
3.FILENAME
- for empty - the result will be displayed on the terminal screen
- for .txt files - the result will be saved as a file with text formatting
- for .csv files - the result will be saved to a file according to the CSV standard (MS Excel)
4.FOLDERS_LEVEL
- number of how many nesting levels to display,
default =2
```
Source code in Golang language.
Tested on Linux Ubuntu
Readme from 03/14/2024

License:
Save information about the author and site in this file.

Author: Alexander Nikitin
https://github.com/ManyakRus/go_lines_count
