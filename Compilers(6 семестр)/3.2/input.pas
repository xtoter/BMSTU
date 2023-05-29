var i:integer;
a,b,c:string;
function Add(a,b: real): real; 
begin 
    Add := a + b; 
end;
begin
a := 'a';b:='b';
for i:=1 to 10 do begin
WriteLn(a);
c:=a;a:=b;b:=c+b;
end
end.
