/*
Binary Search

! ý tưởng: Tại mỗi bước ta so sánh X với phần tử ở giữa của dãy tìm kiếm hiện hành. Dựa vào kết quả so sánh mà ta quyết định giới hạn dãy tìm kiếm ở  nửa dưới hay nửa trên của dãy tìm kiếm hiện hành

* Độ phức tạp O(log(n))

*/

#include <stdio.h>

int binarySearch(int a[], int n, int x)
{
    int left = 0;
    int right = n - 1;
    int mid;
    do
    {
        mid = (left + right) / 2;
        if (a[mid] == x)
            return 1;
        if (a[mid] < x)
        {
            left = mid + 1;
        }
        else
        {
            right = mid - 1;
        }
    } while (left <= right);
}

int main()
{
    int a[] = {1, 2, 4, 6, 7, 9, 10};
    int result = binarySearch(a, 7, 2);
    printf("%d", result);
}