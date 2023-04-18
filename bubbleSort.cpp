/*
Bubble Sort

! ý tưởng: Xuất phát từ cuối dãy, đổi chỗ các cặp phần tử kế cận để đưa phần tử nhỏ hơn trong cặp phần tử đó về vị trí đứng đầu của dãy hiện hành, sau đó sẽ không xét đến nó ở bước tiếp theo , do vậy ở lần xử lí thứ i sẽ có vị trí đầu dãy là i.
    ? Lặp lại xử lý trên cho đến khi không còn cặp phần tử nào để xét

* Độ phức tạp O(n^2)

*/

#include <stdio.h>

int Swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void printArr(int a[], int n)
{
    for (int i = 0; i < n - 1; i++)
        printf("%d\t", a[i]);
}

void bubbleSort(int a[], int n)
{
    for (int i = 0; i < n - 1; i++)
        for (int j = n - 1; j > i; j--)
            if (a[j] < a[j - 1])
                Swap(&a[j], &a[j - 1]);
}

int main()
{
    int a[] = {12, 2, 8, 5, 1, 6, 4, 15};
    bubbleSort(a, 8);
    printArr(a, 8);
    return 0;
}
