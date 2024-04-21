
#include <utility>
#include <vector>

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

extern void TestQsort();
