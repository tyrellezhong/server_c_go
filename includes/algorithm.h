
#include <utility>
#include <vector>


extern void TestQsort();
// 快速排序
template<class T>
int FindMid(std::vector<T>& array, int left, int right) {
    int mid = (left + right) / 2;
    if (array[left] > array[right]) {
        std::swap(array[left], array[right]);
    }
    if (array[left] > array[mid]) {
        std::swap(array[left], array[mid]);
    }
    if (array[mid] > array[right]) {
        std::swap(array[mid], array[right]);
    }
    int ret = array[mid];
    if (right > left)
    {
        std::swap(array[mid], array[right - 1]);
    }
    return ret;
};


template <class Tp>
void Qsort(std::vector<Tp>& array, int left, int right) {
    int pivot = FindMid(array, left, right);
    if (right - left > 2) {
        int i = left;
        int j = right - 1;
        while (true) {
            while (array[++i] < pivot) {}
            while (array[--j] > pivot) {}
            if (i < j) {
                std::swap(array[i], array[j]);
            } else {
                break;
            }
        }
        std::swap(array[i], array[right - 1]);
        Qsort(array, left, i - 1);
        Qsort(array, i + 1, right);
    }
}

// 归并排序
template <class T>
void Merge(std::vector<T>& array, int left, int mid, int right) {
    std::vector<T> temp(right - left + 1);
    int i = left, j = mid + 1, k = 0;
    
    while (i <= mid && j <= right) {
        if (array[i] <= array[j]) {
            temp[k++] = array[i++];
        } else {
            temp[k++] = array[j++];
        }
    }
    
    while (i <= mid) {
        temp[k++] = array[i++];
    }
    
    while (j <= right) {
        temp[k++] = array[j++];
    }
    
    for (int i = 0; i < k; i++) {
        array[left + i] = temp[i];
    }
}

template <class T>
void MergeSort(std::vector<T>& array, int left, int right) {
    if (left < right) {
        int mid = left + (right - left) / 2;
        MergeSort(array, left, mid);
        MergeSort(array, mid + 1, right);
        Merge(array, left, mid, right);
    }
}

extern void TestMergeSort();
