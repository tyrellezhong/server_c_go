// 编译为example.js
// tsc example.ts --outDir ../js_output/
// 执行
// node ../js_output/example.js
function greet(name) {
    return "\u666E\u901A\u51FD\u6570: Hello, ".concat(name, "!");
}
// 箭头函数
var greet2 = function (name) { return "箭头函数: Hello, " + name; };
// 类
var Person = /** @class */ (function () {
    function Person(name, age) {
        this.name = name;
        this.age = age;
    }
    Person.prototype.greet = function () {
        return "Hello, my name is ".concat(this.name);
    };
    return Person;
}());
// 泛型函数
function identity(arg) {
    return arg;
}
// 类型守卫
function isString(value) {
    return typeof value === 'string';
}
// 类型断言
var value = "hello";
var strLength = value.length;
console.log(greet("World"));
console.log(greet2("World"));
console.log("断言: strlength", strLength);
var person = new Person("Alice", 30);
console.log("类函数: " + person.greet());
// 立即执行函数
(function () {
    var x = "函数自调用";
    console.log(x);
})();
// 定义枚举类型，用于表示用户的角色
var Role;
(function (Role) {
    Role[Role["Admin"] = 0] = "Admin";
    Role[Role["User"] = 1] = "User";
    Role[Role["Guest"] = 2] = "Guest";
})(Role || (Role = {}));
// 创建用户对象，符合 User 接口的结构
var user = {
    id: 1,
    username: "Alice",
    isActive: true,
    role: Role.User,
    hobbies: ["Reading", "Gaming"],
    contactInfo: ["+1", 123456789],
};
// 定义一个返回字符串的函数来获取用户信息
function getUserInfo(user) {
    return "User ".concat(user.username, " is ").concat(user.isActive ? "active" : "inactive", " with role ").concat(Role[user.role]);
}
// 使用 void 类型定义一个函数，专门打印用户信息
function printUserInfo(user) {
    console.log(getUserInfo(user));
}
// 定义一个 union 类型的函数参数，接受用户 ID（number）或用户名（string）
function findUser(query) {
    // 使用 typeof 来判断 query 的类型
    if (typeof query === "number") {
        // 如果是数字，则根据 ID 查找用户
        return query === user.id ? user : undefined;
    }
    else if (typeof query === "string") {
        // 如果是字符串，则根据用户名查找用户
        return query === user.username ? user : undefined;
    }
    return undefined;
}
// 定义一个 never 类型的函数，用于处理程序的异常情况
function throwError(message) {
    throw new Error(message);
}
// 使用 any 类型处理未知类型的数据
var unknownData = "This is a string";
unknownData = 42; // 重新赋值为数字，类型为 any
// 使用 unknown 类型处理不确定的数据，更加安全
var someData = "Possible data";
if (typeof someData === "string") {
    console.log("Length of data: ".concat(someData.length));
}
// 调用各个函数并测试
console.log("打印用户信息:");
printUserInfo(user); // 打印用户信息
console.log("根据 ID 查找用户:", findUser(1)); // 根据 ID 查找用户
console.log("根据用户名查找用户:", findUser("Alice")); // 根据用户名查找用户
// 使用 null 和 undefined 类型的变量
var emptyValue = null;
var uninitializedValue = undefined;
