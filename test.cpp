#include <cstdio>
template<typename... Types>
auto sum(Types... args) {
    return (... + args);
}

// int main() {
//     printf("sum is %.2f\n", sum(1.0, 2.0, 3.5, 4.0));
// }

#include <iostream>
#include <string>
#include <vector>

// 假设这些是 visit_struct 库的一部分
namespace visit_struct {
    namespace detail {
        template<typename T, typename U, U T::*Member>
        struct member_ptr_helper {
            static U T::* get_member_ptr() { return Member; }
        };

        template<typename... Members>
        struct Append_t;

        template<typename... Members, typename NewMember>
        struct Append_t<std::tuple<Members...>, NewMember> {
            using type = std::tuple<Members..., NewMember>;
        };

        template<int N>
        struct Rank : Rank<N - 1> {};

        template<>
        struct Rank<0> {};

        template<size_t N>
        struct char_array {
            char data[N];
        };

        constexpr size_t max_visitable_members_intrusive = 10;
    }
}

#define VISIT_STRUCT_CURRENT_TYPE MyStruct
#define VISIT_STRUCT_GET_REGISTERED_MEMBERS std::tuple<>

#define VISIT_STRUCT_MAKE_MEMBER_NAME(NAME) Visit_Struct_Member_Record__##NAME
#define VISIT_STRUCT_CONSTEXPR constexpr

#define CONFIGABLE_VISIT_AUX(TYPE, NAME, CFG_NAME)                                                               \
struct VISIT_STRUCT_MAKE_MEMBER_NAME(NAME) :                                                                     \
  visit_struct::detail::member_ptr_helper<VISIT_STRUCT_CURRENT_TYPE,                                             \
                                          TYPE,                                                                  \
                                          &VISIT_STRUCT_CURRENT_TYPE::NAME>                                      \
{                                                                                                                \
  static VISIT_STRUCT_CONSTEXPR const ::visit_struct::detail::char_array<sizeof(#CFG_NAME)> & member_name() {    \
    return #CFG_NAME;                                                                                            \
  }                                                                                                              \
};                                                                                                               \
static inline ::visit_struct::detail::Append_t<VISIT_STRUCT_GET_REGISTERED_MEMBERS,                              \
                                               VISIT_STRUCT_MAKE_MEMBER_NAME(NAME)>                              \
  Visit_Struct_Get_Visitables__(::visit_struct::detail::Rank<VISIT_STRUCT_GET_REGISTERED_MEMBERS::size + 1>);    \
static_assert(true, "")

struct MyStruct {
    int a;
    double b;
    std::string c;

    CONFIGABLE_VISIT_AUX(int, a, "a");
    CONFIGABLE_VISIT_AUX(double, b, "b");
    CONFIGABLE_VISIT_AUX(std::string, c, "c");
};

// 获取所有成员的元组类型
using MyStructMembers = decltype(Visit_Struct_Get_Visitables__(visit_struct::detail::Rank<visit_struct::detail::max_visitable_members_intrusive>{}));

// 遍历成员并打印
template<typename T, typename Tuple, std::size_t... I>
void print_members(const T& obj, const Tuple& t, std::index_sequence<I...>) {
    (..., (std::cout << std::get<I>(t).member_name().data << ": " << obj.*(std::get<I>(t).get_member_ptr()) << std::endl));
}

template<typename T>
void print_struct(const T& obj) {
    print_members(obj, MyStructMembers{}, std::make_index_sequence<std::tuple_size<MyStructMembers>::value>{});
}

int main() {
    MyStruct s{42, 3.14, "Hello"};
    print_struct(s);
    return 0;
}