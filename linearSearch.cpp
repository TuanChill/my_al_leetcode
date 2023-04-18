/*
LinearSearch

# ý tưởng:
    * So sánh x lần lượt với phần tử thứ 1, thứ 2,... của mảng cho đến khi tìm được khoá cần tìm hoặc không tìm thấy
    * Độ phức tạp O(n)

*/

#include <stdio.h>

int linearSearch(int a[], int n, int x)
{
    int i = 0;
    while (a[i] != x && i < n)
    {
        i++;
    }
    if (i == n)
        return 0;
    return 1;
}

int main()
{
    int a[] = {2, 8, 5, 1, 6, 4, 6};
    int result = linearSearch(a, 7, 9);
    return 0;
}