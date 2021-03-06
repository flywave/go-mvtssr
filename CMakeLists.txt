CMAKE_MINIMUM_REQUIRED(VERSION 3.1.3)

PROJECT(go-mvtssr C CXX)

# Set build type
IF(NOT CMAKE_BUILD_TYPE)
  SET(CMAKE_BUILD_TYPE Release CACHE STRING
      "Choose the type of build, options are: None Debug Release RelWithDebInfo MinSizeRel."
      FORCE)
ENDIF(NOT CMAKE_BUILD_TYPE)

OPTION(FLYWAVE_ENABLE_SOLUTION_FOLDERS "Build with UBSAN" ON)

if(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
  SET_PROPERTY(GLOBAL PROPERTY USE_FOLDERS ON)
  SET_PROPERTY(GLOBAL PROPERTY PREDEFINED_TARGETS_FOLDER "CMakeTargets")
endif()

# Unix specific options
IF(UNIX OR APPLE)
  SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC -Wall")
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC -Wall")
ENDIF()

# CXX_STANDARD
SET(CMAKE_CXX_STANDARD 17)
SET(CMAKE_CXX17_STANDARD_COMPILE_OPTION "-std=c++17")
set(CMAKE_CXX_STANDARD_REQUIRED ON)

IF(UNIX)
IF(CMAKE_BUILD_TYPE STREQUAL "Debug")
  SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g")
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -g")
ENDIF()
ENDIF()
 
# Set output directories
SET(FLYWAVE_BASE_INSTALL_DIR ${CMAKE_CURRENT_SOURCE_DIR}/lib)
IF(APPLE)
  if (${CMAKE_SYSTEM_PROCESSOR} MATCHES "^arm")
    SET(FLYWAVE_LIBRARY_OUTPUT_PATH "${FLYWAVE_BASE_INSTALL_DIR}/darwin_arm")
  else()
    SET(FLYWAVE_LIBRARY_OUTPUT_PATH "${FLYWAVE_BASE_INSTALL_DIR}/darwin")
  endif()
  SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++14")
ELSEIF(WIN32)
SET(FLYWAVE_LIBRARY_OUTPUT_PATH "${FLYWAVE_BASE_INSTALL_DIR}/windows")
ELSE()
    SET(FLYWAVE_LIBRARY_OUTPUT_PATH "${FLYWAVE_BASE_INSTALL_DIR}/linux")
ENDIF()
MAKE_DIRECTORY(${FLYWAVE_LIBRARY_OUTPUT_PATH})

SET(FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}")
SET(FLYWAVE_LIBRARY_DIRS "")

IF ("${CMAKE_SYSTEM_NAME}" STREQUAL "Linux")
# libuv
IF(NOT LIBUV_INCLUDE)
  ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/external/libuv")
  
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/libuv/include")
  LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/external/libuv/")
  LIST(APPEND FLYWAVE_LIBRARY_DEPES "uv")
  SET(LIBUV_INCLUDE YES)
ENDIF()

ENDIF()

# libjpeg
IF(NOT JPEG_INCLUDE)
  ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/external/libjpeg")

  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/libjpeg")
  LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/external/libjpeg/")
  LIST(APPEND FLYWAVE_LIBRARY_DEPES "jpeg")
  SET(JPEG_INCLUDE YES)
ENDIF()

# libpng
IF(NOT PNG_INCLUDE)
  ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/external/libpng")

  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/libpng")
  LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/external/libpng/")
  LIST(APPEND FLYWAVE_LIBRARY_DEPES "png")
  SET(PNG_INCLUDE YES)
ENDIF()

# zlib
IF(NOT ZLIB_INCLUDE)
  ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/external/zlib")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/zlib/")
  LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/external/zlib/")
  LIST(APPEND FLYWAVE_LIBRARY_DEPES "zlib")
  SET(ZLIB_INCLUDE YES)
ENDIF()

# mbgl-core
IF(NOT MBGLC_INCLUDE)
  ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln")
  
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/platform/default/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/deps/geometry.hpp/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/deps/geojson.hpp/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/deps/optional")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/deps/variant/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/extras/rapidjson/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/extras/expected-lite/include")
  LIST(APPEND FLYWAVE_INCLUDE_DIRS "${CMAKE_CURRENT_SOURCE_DIR}/external/mbgln/vendor/mapbox-base/include")

  LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/external/mbgln/")
  LIST(APPEND FLYWAVE_LIBRARY_DEPES "mbgl-core" "mbgl-vendor-nunicode" "mbgl-vendor-parsedate" "mbgl-vendor-sqlite" "mbgl-vendor-icu" "mbgl-vendor-csscolorparse")
  SET(MBGLC_INCLUDE YES)
ENDIF()

LIST(APPEND FLYWAVE_LIBRARY_DIRS "${CMAKE_CURRENT_BINARY_DIR}/src/")

ADD_SUBDIRECTORY("${CMAKE_CURRENT_SOURCE_DIR}/src")

foreach(LIBDIR ${FLYWAVE_LIBRARY_DIRS})
  INSTALL(
        DIRECTORY ${LIBDIR}
        DESTINATION ${FLYWAVE_LIBRARY_OUTPUT_PATH}
        USE_SOURCE_PERMISSIONS FILE_PERMISSIONS
        FILES_MATCHING PATTERN "*.a"
        )
  endforeach()

INSTALL(
    DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/src/"
    DESTINATION ${FLYWAVE_BASE_INSTALL_DIR}
    USE_SOURCE_PERMISSIONS FILE_PERMISSIONS
    FILES_MATCHING PATTERN "*.h"
    )
  
