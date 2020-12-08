t='''
"0"	Binary	Hex
****
*  *
*  *
*  *
****
	11110000
10010000
10010000
10010000
11110000
	0xF0
0x90
0x90
0x90
0xF0
	
"1"	Binary	Hex
  * 
 ** 
  * 
  * 
 ***
	00100000
01100000
00100000
00100000
01110000
	0x20
0x60
0x20
0x20
0x70
"2"	Binary	Hex
****
   *
****
*   
****
	11110000
00010000
11110000
10000000
11110000
	0xF0
0x10
0xF0
0x80
0xF0
	
"3"	Binary	Hex
****
   *
****
   *
****
	11110000
00010000
11110000
00010000
11110000
	0xF0
0x10
0xF0
0x10
0xF0
"4"	Binary	Hex
*  *
*  *
****
   *
   *
	10010000
10010000
11110000
00010000
00010000
	0x90
0x90
0xF0
0x10
0x10
	
"5"	Binary	Hex
****
*   
****
   *
****
	11110000
10000000
11110000
00010000
11110000
	0xF0
0x80
0xF0
0x10
0xF0
"6"	Binary	Hex
****
*   
****
*  *
****
	11110000
10000000
11110000
10010000
11110000
	0xF0
0x80
0xF0
0x90
0xF0
	
"7"	Binary	Hex
****
   *
  * 
 *  
 *  
	11110000
00010000
00100000
01000000
01000000
	0xF0
0x10
0x20
0x40
0x40
"8"	Binary	Hex
****
*  *
****
*  *
****
	11110000
10010000
11110000
10010000
11110000
	0xF0
0x90
0xF0
0x90
0xF0
	
"9"	Binary	Hex
****
*  *
****
   *
****
	11110000
10010000
11110000
00010000
11110000
	0xF0
0x90
0xF0
0x10
0xF0
"A"	Binary	Hex
****
*  *
****
*  *
*  *
	11110000
10010000
11110000
10010000
10010000
	0xF0
0x90
0xF0
0x90
0x90
	
"B"	Binary	Hex
*** 
*  *
*** 
*  *
*** 
	11100000
10010000
11100000
10010000
11100000
	0xE0
0x90
0xE0
0x90
0xE0
"C"	Binary	Hex
****
*   
*   
*   
****
	11110000
10000000
10000000
10000000
11110000
	0xF0
0x80
0x80
0x80
0xF0
	
"D"	Binary	Hex
*** 
*  *
*  *
*  *
*** 
	11100000
10010000
10010000
10010000
11100000
	0xE0
0x90
0x90
0x90
0xE0
"E"	Binary	Hex
****
*   
****
*   
****
	11110000
10000000
11110000
10000000
11110000
	0xF0
0x80
0xF0
0x80
0xF0
	
"F"	Binary	Hex
****
*   
****
*   
*   
	11110000
10000000
11110000
10000000
10000000
	0xF0
0x80
0xF0
0x80
0x80
'''
def TestprintDig(x,hexs):
	for i in range(5*x,5*x+5):
		print("{0:08b}".format(int(hexs[i], 16)))
lines=t.split("\n")
hexs=[]
for i in range(len(lines)):
	line=lines[i].strip()
	if line.startswith("0x"):
		hexs.append(line)
output=""
for i in range(0,80):
	if i%5==0:
		output=output+"//{}\n".format(i//5)
	output=output+"ram.mem[{}]={}\n".format(i,hexs[i])
	#print("{0:08b}".format(int(hexs[i], 16)))
#TestprintDig(15,hexs)
f=open('output.txt','w+')
f.write(output)
f.close()
