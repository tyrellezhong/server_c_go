import gdb
import re
import sys

class ProtobufPrinter:
    def __init__(self, val):
        self.val = val

    def to_string(self):
        # 这里需要根据你的 protobuf 类型调整
        return gdb.execute("call " + str(self.val) + ".ShortDebugString()", to_string=True)

def protobuf_lookup(val):
    type_name = str(val.type)
    if re.match(r'^google::protobuf::Message$', type_name) or re.match(r'^.*::[A-Za-z0-9_]+_Message$', type_name):
        return ProtobufPrinter(val)
    return None

gdb.pretty_printers.append(protobuf_lookup)
