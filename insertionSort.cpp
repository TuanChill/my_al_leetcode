/*
Insertion Sort

! ý tưởng:
    Giả sử có một dãy a0, a1, ..., an
    Trong đó i phần tử đầu tiên a0, a1, ..., ai-1 đã có thứ tự
    TÌm cách chèn phần tử ai vào vị trí thích hợp của đoạn đã được sắp xếp để có dãy mới a0, a1, ...ai trở nên có thứ tự. Ví trí này chính là vị trí giữa 2 phần tử ak-1 và ak thoả mãn ak-1 < ai < ak

* Độ phức tạp O(n^2)

*/

#include <stdio.h>

int Swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void insertionSort(int a[], int n)
{
    int pos, x;
    for (int i = 1; i < n; i++)
    {
        x = a[i];
        pos = i - 1;
        while (pos >= 0 && a[pos] > x)
        {
            a[pos + 1] = a[pos];
            pos--;
        }
        a[pos + 1] = x;
    }
}

void printArr(int a[], int n)
{
    for (int i = 0; i < n - 1; i++)
        printf("%d\t", a[i]);
}

int main()
{
    int a[] = {12, 2, 8, 5, 1, 6, 4, 15};
    insertionSort(a, 8);
    printArr(a, 8);
    return 0;
}