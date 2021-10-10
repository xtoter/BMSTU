assume CS: code, DS:data
data segment
msgg db "Hello World!$"
msg db 00, 00, 00, 00
vara db 03
varb db 100
varc db 06
vard db 03
data ends
code segment
start:

mov AX, data
mov DS, AX
mov AH, 09h
mov CH,"0"
ADD CH,1

mov BX,2
;SHR CH,1
mov msgg[BX], CH
mov DX, offset msgg
int 21h
mov AX,4C00h
int 21h
code ends
end start