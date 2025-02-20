#include "algorithm.h"
#include <cstdio>

void TestQsort() {
    std::vector<int> array{4, 5, 6, 1, 2, 3, 7, 8, 9, 5, 56, 78, 0, 0, 99, 20, 1};
    Qsort(array, 0, array.size() - 1);
    for (int val : array) {
        std::printf("%d ", val);
    }
}
function escapeSpecialChars(value) {
    return value.replace(/([ .\:+?^${}()|\[\]\\])/g, '\\\\$1');
  }
  
  var query_string = '*'
  var param_array = []
  if ('$ServiceName' && '${ServiceName}' != '*') {
    param_array.push('(resource.service.name : \\\"$ServiceName\\\")')
  }
  if ('$NodeIP') {
    param_array.push('(resource.net.host.ip : \\\"$NodeIP\\\")')
  }
  if ('$TbusAddr') {
    param_array.push('(resource.tbus_addr : \\\"$TbusAddr\\\")')
  }
  if ('$Env' && '${Env}' != '*') {
    param_array.push('(resource.env : \\\"$Env\\\")')
  }
  if ('${PlayerId}') {
    param_array.push('(attributes.uid : \\\"${PlayerId}\\\")')
  }
  if ('${TaskName}') {
    param_array.push('(attributes.task_name : \\\"${TaskName}\\\")')
  }
  if ('$TraceId') {
    param_array.push('(trace_id : \\\"$TraceId\\\")')
  }
  if ('${Level}' && '${Level}' != '*') {
    strLevel = '${Level:pipe}'
    strLevel = strLevel.replaceAll('|', " OR ")
    param_array.push('(severity_text : ' + strLevel + ')')
  }
  if ('${FileName}') {
    var file_name = '${FileName}'
    if (!file_name.includes(':')) {
      file_name += '*'
      file_name = file_name.replaceAll('/', '\\\\/')
      param_array.push('(attributes.file_name : ' + file_name + ')')
    } else {
      param_array.push('(attributes.file_name : \\\"' + file_name + '\\\")')
    }
  }
  if ('${Event}') {
    param_array.push('(body : \\\"${Event}\\\")')
  }
  // if ('${AttrKey}' || '${AttrValue}') {
  //   var key = '${AttrKey}'
  //   if (!key) {
  //     key = '\\\\*'
  //   }
  //   var value = '${AttrValue}'
  //   if (!value) {
  //     value = '*'
  //   }
  //   param_array.push('attributes.' + key + ' : ' + value)
  // }
  
  if ('${AttrExt}') {
    const attrExt = '${AttrExt}';
    // 将字符串按逗号分割
    const pairs = attrExt.split(',');
    // 创建一个数组来存储键值对
    const result = pairs.map(pair => {
      // 将每个键值对按等号分割
      const [key, value] = pair.split('=');
      // 返回一个对象包含键和值
      return {
        key: key ? key.replace(/^"|"$/g, '').trim() : '',
        value: value ? escapeSpecialChars(value.replace(/^"|"$/g, '').trim()) : ''
      };
    });
    // 遍历结果
    result.forEach(item => {
      key = `${item.key}`
      value = `${item.value}`
      if (key || value) {
        if (!key) {
          key = '\\\\*'
        }
        if (!value) {
          value = '*'
        }
        param_array.push('attributes.' + key + ' : ' + value)
      }
    });
  }
  if (param_array.length > 0) {
    query_string = param_array.join(' AND ')
  }
  
  const result = {
    ...frame,
    fields: frame.fields.map((field) => ({
      ...field,
      values: [query_string]
    }))
  }
  
  return Promise.resolve(result)