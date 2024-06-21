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
Name                     	Level	Lines count	Functions count
.                        	1	505817	18692
calendar                 	2	684	22
casebook_service         	2	7510	47
contract_checker_service 	2	831	13
copy                     	2	0	0
crud_generator           	2	17206	588
debezium_adapter_postgres	2	11060	381
debtors_list             	2	2407	35
go_lines_count           	2	589	23
image_connections        	2	1779	52
image_database           	2	1836	45
image_packages           	2	2379	72
monitor_service          	2	2670	23
stack_exchange_postgres  	2	16720	397
starter                  	2	11718	598
sync_service             	2	427362	16362
telegram_loki            	2	639	22
whatsapp_chatgpt         	2	427	12
```
I wrote 505 thousand lines of code in 2 years, 
including 415 thousand lines with a code generator and 80 thousand lines of code manually.

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
