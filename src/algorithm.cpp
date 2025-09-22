#include "algorithm.h"
#include <cstdio>

void TestQsort() {
    std::vector<int> array{4, 5, 6, 1, 2, 3, 7, 8, 9, 5, 56, 78, 0, 0, 99, 20, 1};
    Qsort(array, 0, array.size() - 1);
    for (int val : array) {
        std::printf("%d ", val);
    }
}

void TestMergeSort() {
    std::vector<int> array{4, 5, 6, 1, 2, 3, 7, 8, 9, 5, 56, 78, 0, 0, 99, 20, 1};
    printf("Original array: ");
    for (int val : array) {
        printf("%d ", val);
    }
    printf("\n");
    
    MergeSort(array, 0, array.size() - 1);
    
    printf("Sorted array: ");
    for (int val : array) {
        printf("%d ", val);
    }
    printf("\n");
}
