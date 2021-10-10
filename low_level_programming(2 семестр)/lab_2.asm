assume cs: code, ds: data
;a*b+c/d+5
data segment
	outstring db "0000000$"
    
    arr	Dw  200,  233,  206,  240,  230, 200
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
    mov ax,0000h
	ret
numberoutput endp

tostring proc
        mov si,4 ; Индекс элемента
         mov DI,cx
		mov cx, 5 ;итерации цикла
        mov bp,bx
		MOV BL,10
        mov outstring[5], 10
        mov outstring[6], 13
		goto: 
			DIV BL ;Получаем очередную цифру (делим на 10)
			mov outstring[si], ah
			add outstring[si],"0" ;Для вывода цифры
			mov ah,0
			sub si,1 ;si=si-1
		loop goto
        mov cx,DI
         mov bx,bp
        mov ax,0000h
        ret
tostring endp

start:	
    mov ax, data
	mov ds, ax
    mov cx, 5 ;итерации цикла      
    mov dx,arr[bx] ;0 элемент кладем в dx
    mov bx,2 ;идем с 1 элемента
	gotoo: 
        mov ax,arr[bx]
        CMP	 ax,dx    
        JA	SKIP
		mov dx,ax
        SKIP: 
        add bx,2
    loop gotoo
    mov ax,dx ;Кладем результат в ax для вывода
    call tostring
    call numberoutput ;Вывод числа
	mov AX,4C00h
	int 21h
code ends
end start