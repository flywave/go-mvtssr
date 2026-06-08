# ANGLE — OpenGL ES 2.0/3.0 implementation for Windows via Direct3D.
#
# Builds static libraries libEGL_static.lib + libGLESv2_static.lib.
# ANGLE uses GN + Ninja natively; this CMake wrapper drives that build.
#
# Prerequisites (Windows only):
#   1. Git (in PATH)
#   2. Python ≥ 3.6  (in PATH)
#   3. Ninja          (in PATH, or downloaded automatically)
#   4. GN             (in PATH, or downloaded automatically)
#   5. Windows SDK    (with D3D11 headers)

if(NOT WIN32)
  return()
endif()

include(ExternalProject)

# ── Version pin ──────────────────────────────────────────────────────
set(ANGLE_TAG "chromium/5981")          # git tag / commit-ish
set(ANGLE_REPO "https://chromium.googlesource.com/angle/angle.git")
set(ANGLE_SRC_DIR  "${CMAKE_CURRENT_SOURCE_DIR}/src")
set(ANGLE_BUILD_DIR "${CMAKE_CURRENT_BINARY_DIR}/_build")
set(ANGLE_OUT_DIR   "${CMAKE_CURRENT_BINARY_DIR}/_out")

# ── Fetch source (shallow clone, no .git in our tree) ────────────────
ExternalProject_Add(angle_fetch
  GIT_REPOSITORY    ${ANGLE_REPO}
  GIT_TAG           ${ANGLE_TAG}
  GIT_SHALLOW       TRUE
  SOURCE_DIR        ${ANGLE_SRC_DIR}
  BINARY_DIR        ${ANGLE_BUILD_DIR}
  CONFIGURE_COMMAND ""
  BUILD_COMMAND     ""
  INSTALL_COMMAND   ""
  LOG_DOWNLOAD      ON
  LOG_CONFIGURE     ON
  LOG_BUILD         ON
)

# ── GN + Ninja configuration ────────────────────────────────────────
# GN args: release build, x64, D3D11 + D3D9 backends, no Vulkan/GL
set(ANGLE_GN_ARGS
  is_debug=false
  is_clang=false
  target_cpu=\"x64\"
  angle_enable_d3d9=true
  angle_enable_d3d11=true
  angle_enable_vulkan=false
  angle_enable_gl=false
  angle_enable_swiftshader=false
  angle_enable_null=false
  angle_enable_essl=false
  angle_build_all=false
)

# ── Determine GN + Ninja binaries ───────────────────────────────────
find_program(GN    gn)
find_program(NINJA ninja)

if(NOT GN)
  message(STATUS "GN not found in PATH — will download")
  set(_gn_url "https://chrome-infra-packages.appspot.com/dl/gn/gn/${CMAKE_VS_PLATFORM_NAME:-win64}/+/latest")
  set(GN "${CMAKE_CURRENT_BINARY_DIR}/tools/gn/gn")
  ExternalProject_Add(gn_binary
    URL               ${_gn_url}
    SOURCE_DIR        "${CMAKE_CURRENT_BINARY_DIR}/tools/gn"
    CONFIGURE_COMMAND ""
    BUILD_COMMAND     ""
    INSTALL_COMMAND   ""
  )
else()
  add_custom_target(gn_binary)   # no-op
endif()

if(NOT NINJA)
  message(STATUS "Ninja not found in PATH — will download")
  set(_ninja_url "https://github.com/ninja-build/ninja/releases/download/v1.12.1/ninja-win.zip")
  set(NINJA "${CMAKE_CURRENT_BINARY_DIR}/tools/ninja/ninja")
  ExternalProject_Add(ninja_binary
    URL               ${_ninja_url}
    SOURCE_DIR        "${CMAKE_CURRENT_BINARY_DIR}/tools/ninja"
    CONFIGURE_COMMAND ""
    BUILD_COMMAND     ""
    INSTALL_COMMAND   ""
  )
else()
  add_custom_target(ninja_binary)   # no-op
endif()

# ── Build step: gn gen → ninja ──────────────────────────────────────
ExternalProject_Add(angle_build
  DEPENDS           angle_fetch gn_binary ninja_binary
  DOWNLOAD_COMMAND  ""
  SOURCE_DIR        ${ANGLE_SRC_DIR}
  BINARY_DIR        ${ANGLE_BUILD_DIR}
  CONFIGURE_COMMAND ${GN} gen ${ANGLE_OUT_DIR} --args="${ANGLE_GN_ARGS}"
  BUILD_COMMAND     ${NINJA} -C ${ANGLE_OUT_DIR} libEGL_static libGLESv2_static
  INSTALL_COMMAND   ""
  BUILD_BYPRODUCTS
    "${ANGLE_OUT_DIR}/libEGL_static.lib"
    "${ANGLE_OUT_DIR}/libGLESv2_static.lib"
)

# ── Imported CMake targets ───────────────────────────────────────────
add_library(angle_egl    STATIC IMPORTED GLOBAL)
add_library(angle_glesv2 STATIC IMPORTED GLOBAL)

set_target_properties(angle_egl PROPERTIES
  IMPORTED_LOCATION             "${ANGLE_OUT_DIR}/libEGL_static.lib"
  INTERFACE_INCLUDE_DIRECTORIES "${ANGLE_SRC_DIR}/include"
)
set_target_properties(angle_glesv2 PROPERTIES
  IMPORTED_LOCATION              "${ANGLE_OUT_DIR}/libGLESv2_static.lib"
  INTERFACE_INCLUDE_DIRECTORIES  "${ANGLE_SRC_DIR}/include"
)

add_dependencies(angle_egl    angle_build)
add_dependencies(angle_glesv2 angle_build)

# ── Install into lib/windows/ ────────────────────────────────────────
install(FILES
    "${ANGLE_OUT_DIR}/libEGL_static.lib"
    "${ANGLE_OUT_DIR}/libGLESv2_static.lib"
  DESTINATION ${FLYWAVE_LIBRARY_OUTPUT_PATH}
)
install(DIRECTORY "${ANGLE_SRC_DIR}/include/"
  DESTINATION ${FLYWAVE_BASE_INSTALL_DIR}/angle_include
  FILES_MATCHING PATTERN "*.h"
)
