/*
Interchange Sort

! ý tưởng: Xuất phát từ đầu dãy, tìm tất cả các nghịch thế chứa phần tử này, triệt tiêu chúng bằng cách đổi chỗ 2 phần tử trong cặp nghịch thế. Lặp lại xử lí trên với phần tử kế trong dãy

* Độ phức tạp O(n^2)

*/

#include <stdio.h>

int Swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

int interchangeSort(int a[], int n)
{
    for (int i = 0; i < n - 2; i++)
        for (int j = i + 1; j < n - 1; j++)
            if (a[j] < a[i])
                Swap(&a[i], &a[j]);
}

void printArr(int a[], int n)
{
    for (int i = 0; i < n - 1; i++)
        printf("%d\t", a[i]);
}

int main()
{
    int a[] = {12, 2, 8, 5, 1, 6, 4, 15};
    interchangeSort(a, 8);
    printArr(a, 8);
}