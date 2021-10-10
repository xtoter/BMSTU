assume cs: code, ds: data
;a*b+c/d+5
data segment
	outstring db "00000$"
	vara db 03
	varb db 100
	varc db 06
	vard db 03
	data ends
code segment


numberoutput proc ;Вывод строки
	mov AX, data
	mov DS, AX
	mov AH, 09h
	mov DX, offset outstring
	int 21h
	mov AX,4C00h
	int 21h
	ret
numberoutput endp


start:	mov ax, data
		mov ds, ax
		mov ax, 0000h
		mov bx, 0000h
		mov al, vara  ;al=a
		mov bl, varb  ;bl=b
		mul bl  ;al=al*bl
		mov cx, ax ; cx=ax
		mov ax, 0000h
		mov bx, 0000h
		mov al, varc ;al=c
		mov bl, vard ;bl=d
		div bl ;bl =al/bl
		add ax, cx ;ax=ax+cx
		add al,5 ; al=al+5

		mov si,4 ; Индекс элемента
		mov cx, 5 ;итерации цикла
		MOV BL,10
		goto: 
			DIV BL ;Получаем очередную цифру (делим на 10)
			mov outstring[si], ah
			add outstring[si],"0" ;Для вывода цифры
			mov ah,0
			sub si,1 ;si=si-1
		loop goto


		call numberoutput ;Вывод числа
code ends
end start