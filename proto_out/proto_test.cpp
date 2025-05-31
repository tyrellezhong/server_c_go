#include "Log.h"
#include "proto_option.pb.h"
void TestProtoOption() {
    // 创建 MyMessage 实例
    msg::MyMessage message;
    message.set_name("John");
    message.set_age(42);
    message.mutable_person()->mutable_map()->insert({1, 2});
    // message.PrintDebugString();
    // LogInfo("ShotDebugString=%s", message.ShortDebugString().c_str());
    // LogInfo("DebugString=%s", message.DebugString().c_str());
    // LogInfo("Utf8DebugString=%s", message.Utf8DebugString().c_str());

    // 访问字段的自定义选项
    const google::protobuf::FieldDescriptor* name_field = message.GetDescriptor()->FindFieldByName("name");
    const google::protobuf::FieldDescriptor* age_field = message.GetDescriptor()->FindFieldByName("age");

    // 获取自定义选项
    if (name_field->options().HasExtension(msg::my_option)) {
        std::string name_option = name_field->options().GetExtension(msg::my_option);
        std::cout << "Custom option for 'name': " << name_option << std::endl;
    }

    if (age_field->options().HasExtension(msg::my_option)) {
        std::string age_option = age_field->options().GetExtension(msg::my_option);
        std::cout << "Custom option for 'age': " << age_option << std::endl;
    }

    std::cout << "debug string: " << message.DebugString() << std::endl;
    std::cout << "short debug string: " << message.ShortDebugString() << std::endl;

}