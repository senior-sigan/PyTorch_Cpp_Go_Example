cmake_minimum_required(VERSION 3.12)
project(pytorch_cpp_go_example)

set(CMAKE_CXX_STANDARD 17)

# TODO: somehow handle library to use. Maybe cmake flags??
set(CMAKE_PREFIX_PATH ./libs/osx/libtorch)
# set(CMAKE_PREFIX_PATH ./libs/linux_cpu/libtorch)
# set(CMAKE_PREFIX_PATH ./libs/linux_gpu/libtorch)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

find_package(Torch REQUIRED)

add_library(${PROJECT_NAME} library.cpp library.h c_api.h c_library.h c_library.cpp)
target_link_libraries(${PROJECT_NAME} "${TORCH_LIBRARIES}")