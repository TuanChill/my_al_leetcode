/*
Interchange Sort

! ý tưởng:
    ? Chọn phần tử nhỏ nhất trong n phần tử trong dãy hiện hành ban đầu.
    ? Đưa phần tử này về vị trí ban đầu dãy hiện hành
    ? Xem dãy hiện hành chỉ còn n-1 phần tử của dãy hiện hành ban đầu
        ? Bắt đầu từ vị trí thứ 2;
        ? Lặp lại quá trình trên cho dãy hiện hành đến khi dãy hiện hành chỉ còn 1 phần tử

* Độ phức tạp O(n^2)

*/
#include <stdio.h>

int Swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

int selectionSort(int a[], int n)
{
    int min;
    for (int i = 0; i < n - 1; i++)
    {
        min = i;
        for (int j = i + 1; j < n; j++)
            if (a[j] < a[min])
                min = j;
        Swap(&a[min], &a[i]);
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
    selectionSort(a, 8);
    printArr(a, 8);
    return 0;
}