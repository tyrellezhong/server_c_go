#include "time_test.h"
#include <chrono>
#include <iostream>
#include <ostream>
#include "Log.h"
#include <sys/time.h>
#include <iomanip>

void TimeTest::ChronoTimeTest() {
    DebugBegin("ChronoTimeTest");
    using namespace std::chrono;
    // 获取当前时间点
    auto now = std::chrono::system_clock::now();

    // 转换为时间戳（自 Unix 纪元以来的秒数）
    auto timestamp = std::chrono::system_clock::to_time_t(now);
    auto time_point = std::chrono::system_clock::from_time_t(timestamp);

    // 打印当前时间戳
    LogInfo("当前时间戳：%ld", timestamp);

    // 将时间戳转换为本地时间
    std::tm* local_time = std::localtime(&timestamp);

    // 格式化时间输出
    char buffer[100];
    std::strftime(buffer, sizeof(buffer), "%Y-%m-%d %H:%M:%S", local_time);

    // 打印格式化后的本地时间
    LogInfo("当前时间：%s", buffer)

     // 转换为时间戳（秒）
    auto duration = now.time_since_epoch();
    auto seconds = std::chrono::duration_cast<std::chrono::seconds>(duration).count();
    auto nanoseconds = std::chrono::duration_cast<std::chrono::nanoseconds>(duration).count() % 1000000000;

    // 打印当前时间戳（秒）
    std::cout << "当前时间戳: " << seconds << " 秒  " << "纳秒 " << std::chrono::duration_cast<std::chrono::nanoseconds>(duration).count() << std::endl;

    // 打印格式化后的本地时间和纳秒
    std::cout << "格式化后的本地时间: " << buffer << "." << std::setw(9) << std::setfill('0') << nanoseconds << std::endl;

    // 定义一个 time_point 类型，使用系统时钟
    using TimePoint = std::chrono::time_point<std::chrono::system_clock>;

    // 获取 time_point 的最小值和最大值
    TimePoint min_time = TimePoint::min();
    TimePoint max_time = TimePoint::max();

    // 将 time_point 转换为 time_t 类型
    std::time_t min_time_t = std::chrono::system_clock::to_time_t(min_time); // 负数，后文显示可能会出问题
    std::time_t max_time_t = std::chrono::system_clock::to_time_t(max_time);

    // 转换为本地时间
    std::tm* min_local_time = std::localtime(&min_time_t);
    std::tm* max_local_time = std::localtime(&max_time_t);

    // 格式化时间输出
    char min_buffer[100];
    char max_buffer[100];
    std::strftime(min_buffer, sizeof(min_buffer), "%Y-%m-%d %H:%M:%S", min_local_time);
    std::strftime(max_buffer, sizeof(max_buffer), "%Y-%m-%d %H:%M:%S", max_local_time);

    // 打印最小值和最大值
    std::cout << "time_point 最小值: " << min_buffer << std::endl;
    std::cout << "time_point 最大值: " << max_buffer << std::endl;

    // 定义一个 duration 类型，使用秒为单位
    using Duration = std::chrono::duration<int>;

    // 获取 duration 的最小值、最大值和零值
    Duration min_duration = Duration::min();
    Duration max_duration = Duration::max();
    Duration zero_duration = Duration::zero();

    // 打印最小值、最大值和零值
    std::cout << "duration 最小值: " << min_duration.count() << " 秒" << std::endl;
    std::cout << "duration 最大值: " << max_duration.count() << " 秒" << std::endl;
    std::cout << "duration 零值: " << zero_duration.count() << " 秒" << std::endl;


    // 定义两个 duration 对象
    std::chrono::duration<int> dur1 = std::chrono::seconds(10);
    std::chrono::duration<int> dur2 = std::chrono::minutes(1);

    // 将两个 duration 对象相加
    std::chrono::duration<int> result = dur1 + dur2;

    // 打印结果
    std::cout << "结果: " << result.count() << " 秒" << std::endl;

    // 定义一个 duration 对象
    std::chrono::duration<int> dur = std::chrono::hours(3);

    // 将 duration 对象加到 time_point 对象上
    std::chrono::time_point<std::chrono::system_clock> future_time = now + dur;

    // 打印结果
    std::time_t future_time_t = std::chrono::system_clock::to_time_t(future_time);
    std::cout << "未来时间: " << std::ctime(&future_time_t);

    // 比较两个 duration 对象
    if (dur1 < dur2) {
        std::cout << "dur1 小于 dur2" << std::endl;
    } else {
        std::cout << "dur1 不小于 dur2" << std::endl;
    }

    // 比较两个 time_point 对象
    if (now < future_time) {
        std::cout << "现在时间小于未来时间" << std::endl;
    } else {
        std::cout << "现在时间不小于未来时间" << std::endl;
    }

    auto time_now = std::chrono::time_point_cast<std::chrono::duration<int64_t, std::ratio<24 * 60 * 60, 1>>>(std::chrono::system_clock::now());

    std::cout << "form epoch days: " << time_now.time_since_epoch().count() << std::endl;
 
  
  

}

void TimeTest::CTimeTest() {
    DebugBegin("CTimeTest");
    // 获取当前时间戳
    time_t timestamp = time(NULL);

    struct timeval tv;
    gettimeofday(&tv, NULL);

    // 打印当前时间戳
    LogInfo("当前时间戳: %ld", timestamp);

     // 打印当前时间戳（秒和微秒）
    LogInfo("当前时间戳: %ld 秒, %ld 微秒", tv.tv_sec, tv.tv_usec);

    // 将时间戳转换为本地时间
    struct tm *local_time = localtime(&timestamp);

    // 格式化时间输出
    char buffer[100];
    strftime(buffer, sizeof(buffer), "%Y-%m-%d %H:%M:%S", local_time);

    // 打印格式化后的本地时间
    LogInfo("格式化后的本地时间: %s", buffer);

}

