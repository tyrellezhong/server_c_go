set(APP main_go)
set(EXECUTABLE_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/build)

add_custom_target(${APP} ALL
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMAND go build "-gcflags=all=-N -l" -o ${EXECUTABLE_OUTPUT_PATH}/${APP}
        VERBATIM
        )

# 将部署可执行程序
# install(PROGRAMS ${EXECUTABLE_OUTPUT_PATH}/${APP} DESTINATION ${CMAKE_INSTALL_PREFIX}/${APP}/bin)
