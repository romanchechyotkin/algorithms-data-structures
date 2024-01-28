# ARRAY

 Массив - ссылка на первый элемент в памяти, для получения последующих элементов => изначальный адрес + размер типа Х индекс искомого элемента    

для получения элемента с индексом 3 в массиве 8 бит чисел в массиве, адрес которого начинается в 0х00

arr[3] = 0x00 + 3 X 8 = 0x24 => O(1)

- Read time - O(1)

- Insert time to end - O(1)
- Insert time to middle - O(n)
- Insert time to start - O(n)

- Remove time from end - O(1)
- Remove time from middle - O(n)
- Remove time from start - O(n)
