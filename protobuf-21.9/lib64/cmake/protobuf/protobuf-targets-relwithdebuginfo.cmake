#----------------------------------------------------------------
# Generated CMake target import file for configuration "RelWithDebugInfo".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "protobuf::libprotobuf-lite" for configuration "RelWithDebugInfo"
set_property(TARGET protobuf::libprotobuf-lite APPEND PROPERTY IMPORTED_CONFIGURATIONS RELWITHDEBUGINFO)
set_target_properties(protobuf::libprotobuf-lite PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_RELWITHDEBUGINFO "CXX"
  IMPORTED_LOCATION_RELWITHDEBUGINFO "${_IMPORT_PREFIX}/lib64/libprotobuf-lite.a"
  )

list(APPEND _cmake_import_check_targets protobuf::libprotobuf-lite )
list(APPEND _cmake_import_check_files_for_protobuf::libprotobuf-lite "${_IMPORT_PREFIX}/lib64/libprotobuf-lite.a" )

# Import target "protobuf::libprotobuf" for configuration "RelWithDebugInfo"
set_property(TARGET protobuf::libprotobuf APPEND PROPERTY IMPORTED_CONFIGURATIONS RELWITHDEBUGINFO)
set_target_properties(protobuf::libprotobuf PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_RELWITHDEBUGINFO "CXX"
  IMPORTED_LOCATION_RELWITHDEBUGINFO "${_IMPORT_PREFIX}/lib64/libprotobuf.a"
  )

list(APPEND _cmake_import_check_targets protobuf::libprotobuf )
list(APPEND _cmake_import_check_files_for_protobuf::libprotobuf "${_IMPORT_PREFIX}/lib64/libprotobuf.a" )

# Import target "protobuf::libprotoc" for configuration "RelWithDebugInfo"
set_property(TARGET protobuf::libprotoc APPEND PROPERTY IMPORTED_CONFIGURATIONS RELWITHDEBUGINFO)
set_target_properties(protobuf::libprotoc PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_RELWITHDEBUGINFO "CXX"
  IMPORTED_LOCATION_RELWITHDEBUGINFO "${_IMPORT_PREFIX}/lib64/libprotoc.a"
  )

list(APPEND _cmake_import_check_targets protobuf::libprotoc )
list(APPEND _cmake_import_check_files_for_protobuf::libprotoc "${_IMPORT_PREFIX}/lib64/libprotoc.a" )

# Import target "protobuf::protoc" for configuration "RelWithDebugInfo"
set_property(TARGET protobuf::protoc APPEND PROPERTY IMPORTED_CONFIGURATIONS RELWITHDEBUGINFO)
set_target_properties(protobuf::protoc PROPERTIES
  IMPORTED_LOCATION_RELWITHDEBUGINFO "${_IMPORT_PREFIX}/bin/protoc-3.21.9.0"
  )

list(APPEND _cmake_import_check_targets protobuf::protoc )
list(APPEND _cmake_import_check_files_for_protobuf::protoc "${_IMPORT_PREFIX}/bin/protoc-3.21.9.0" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
