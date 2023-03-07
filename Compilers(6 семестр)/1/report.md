% Лабораторная работа № 1.1. Раскрутка самоприменимого компилятора
% 7 февраля 2023 г.
% Терюха Михаил, ИУ9-62Б

# Цель работы
Целью данной работы является ознакомление с раскруткой самоприменимых компиляторов на примере модельного компилятора.

# Индивидуальный вариант
Компилятор BeRo. Обеспечить возможность использования символа @ в идентификаторах.


# Реализация

Различие между файлами `pcom.pas` и `pcom2.pas`:

```diff
627c627
<  if (('a'<=CurrentChar) and (CurrentChar<='z')) or (('A'<=CurrentChar) and (CurrentChar<='Z')) then begin
---
>  if (('a'<=CurrentChar) and (CurrentChar<='z')) or (('A'<=CurrentChar) and (CurrentChar<='Z')) or (CurrentChar='@') then begin
632c632
<         (CurrentChar='_') do begin
---
>         (CurrentChar='_') or (CurrentChar='@') do begin
635c635
<     if ('a'<=CurrentChar) and (CurrentChar<='z') then begin
---
>     if ('a'<=CurrentChar) and (CurrentChar<='z') or (CurrentChar='@') then begin
```

Различие между файлами `pcom2.pas` и `pcom3.pas`:

```diff
258c258
< procedure StringCopy(var Dest:TAlfa;Src:TAlfa);
---
> procedure String@Copy(var Dest:TAlfa;Src:TAlfa);
2957,2986c2957,2986
<  StringCopy(Keywords[SymBEGIN],'BEGIN               ');
<  StringCopy(Keywords[SymEND],'END                 ');
<  StringCopy(Keywords[SymIF],'IF                  ');
<  StringCopy(Keywords[SymTHEN],'THEN                ');
<  StringCopy(Keywords[SymELSE],'ELSE                ');
<  StringCopy(Keywords[SymWHILE],'WHILE               ');
<  StringCopy(Keywords[SymDO],'DO                  ');
<  StringCopy(Keywords[SymCASE],'CASE                ');
<  StringCopy(Keywords[SymREPEAT],'REPEAT              ');
<  StringCopy(Keywords[SymUNTIL],'UNTIL               ');
<  StringCopy(Keywords[SymFOR],'FOR                 ');
<  StringCopy(Keywords[SymTO],'TO                  ');
<  StringCopy(Keywords[SymDOWNTO],'DOWNTO              ');
<  StringCopy(Keywords[SymNOT],'NOT                 ');
<  StringCopy(Keywords[SymDIV],'DIV                 ');
<  StringCopy(Keywords[SymMOD],'MOD                 ');
<  StringCopy(Keywords[SymAND],'AND                 ');
<  StringCopy(Keywords[SymOR],'OR                  ');
<  StringCopy(Keywords[SymCONST],'CONST               ');
<  StringCopy(Keywords[SymVAR],'VAR                 ');
<  StringCopy(Keywords[SymTYPE],'TYPE                ');
<  StringCopy(Keywords[SymARRAY],'ARRAY               ');
<  StringCopy(Keywords[SymOF],'OF                  ');
<  StringCopy(Keywords[SymPACKED],'PACKED              ');
<  StringCopy(Keywords[SymRECORD],'RECORD              ');
<  StringCopy(Keywords[SymPROGRAM],'PROGRAM             ');
<  StringCopy(Keywords[SymFORWARD],'FORWARD             ');
<  StringCopy(Keywords[SymHALT],'HALT                ');
<  StringCopy(Keywords[SymFUNC],'FUNCTION            ');
<  StringCopy(Keywords[SymPROC],'PROCEDURE           ');
---
>  String@Copy(Keywords[SymBEGIN],'BEGIN               ');
>  String@Copy(Keywords[SymEND],'END                 ');
>  String@Copy(Keywords[SymIF],'IF                  ');
>  String@Copy(Keywords[SymTHEN],'THEN                ');
>  String@Copy(Keywords[SymELSE],'ELSE                ');
>  String@Copy(Keywords[SymWHILE],'WHILE               ');
>  String@Copy(Keywords[SymDO],'DO                  ');
>  String@Copy(Keywords[SymCASE],'CASE                ');
>  String@Copy(Keywords[SymREPEAT],'REPEAT              ');
>  String@Copy(Keywords[SymUNTIL],'UNTIL               ');
>  String@Copy(Keywords[SymFOR],'FOR                 ');
>  String@Copy(Keywords[SymTO],'TO                  ');
>  String@Copy(Keywords[SymDOWNTO],'DOWNTO              ');
>  String@Copy(Keywords[SymNOT],'NOT                 ');
>  String@Copy(Keywords[SymDIV],'DIV                 ');
>  String@Copy(Keywords[SymMOD],'MOD                 ');
>  String@Copy(Keywords[SymAND],'AND                 ');
>  String@Copy(Keywords[SymOR],'OR                  ');
>  String@Copy(Keywords[SymCONST],'CONST               ');
>  String@Copy(Keywords[SymVAR],'VAR                 ');
>  String@Copy(Keywords[SymTYPE],'TYPE                ');
>  String@Copy(Keywords[SymARRAY],'ARRAY               ');
>  String@Copy(Keywords[SymOF],'OF                  ');
>  String@Copy(Keywords[SymPACKED],'PACKED              ');
>  String@Copy(Keywords[SymRECORD],'RECORD              ');
>  String@Copy(Keywords[SymPROGRAM],'PROGRAM             ');
>  String@Copy(Keywords[SymFORWARD],'FORWARD             ');
>  String@Copy(Keywords[SymHALT],'HALT                ');
>  String@Copy(Keywords[SymFUNC],'FUNCTION            ');
>  String@Copy(Keywords[SymPROC],'PROCEDURE           ');

```

# Тестирование

Тестовый пример:

```pascal
program Hello;
procedure hell@oz;
 begin
  writeln('hellox');
end;
begin
   hell@oz;
end.
```

Вывод тестового примера на `stdout`

```
hellox
```

# Вывод
Я ознакомился с понятием раскрутки самоприменимого компилятора, изучил устройство компилятора BeRo, расширил на восприятия нового символа в индентификаторах.
