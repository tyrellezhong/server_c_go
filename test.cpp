#include <cstdio>
template<typename... Types>
auto sum(Types... args) {
    return (... + args);
}

int main() {
    printf("sum is %.2f\n", sum(1.0, 2.0, 3.5, 4.0));
}