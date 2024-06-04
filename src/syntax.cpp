#include <cstdarg>
// 需要提供参数个数
int sum(int n, ...) {
    int total = 0;
    va_list args;
    va_start(args, n);
    for (int i = 0; i < n; ++i) {
        total += va_arg(args, int);
    }
    va_end(args);
    return total;
}
