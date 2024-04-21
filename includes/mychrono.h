#include <chrono>

typedef std::chrono::duration<int> seconds_type;
typedef std::chrono::duration<int,std::milli> milliseconds_type;
typedef std::chrono::duration<int,std::ratio<60*60>> hours_type;