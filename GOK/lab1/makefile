lab1:
	mkdir -p plugin
	g++ -g -Wall -Wextra -std=c++14 -I`gcc -print-file-name=plugin`/include -fPIC -fno-rtti -shared src/lab1/lab1.cpp -o plugin/$@.so

run:
	gcc -fplugin=plugin/lab1.so -O0 src/debug/debug.c -o src/debug/debug
	rm src/debug/debug

dot:
	gcc -fplugin=plugin/lab1.so -O0 -fdump-tree-ssa-graph -g src/debug/debug.c -o src/debug/debug
	dot -Tpng ./src/debug/debug.c.023t.ssa.dot -oout.png
	fim -a out.png
	rm src/debug/debug

.PHONY : clean

clean:
	rm -rf plugin src/debug/debug
