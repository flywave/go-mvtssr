set(CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}/CMake;${CMAKE_MODULE_PATH}")
include(Utilities)
include(Macros)
include(CMakeDependentOption)
include(CheckCCompilerFlag)

project( CURL C )

message(WARNING "the curl cmake build system is poorly maintained. Be aware")

file (READ ${CURL_SOURCE_DIR}/include/curl/curlver.h CURL_VERSION_H_CONTENTS)
string (REGEX MATCH "#define LIBCURL_VERSION \"[^\"]*"
  CURL_VERSION ${CURL_VERSION_H_CONTENTS})
string (REGEX REPLACE "[^\"]+\"" "" CURL_VERSION ${CURL_VERSION})
string (REGEX MATCH "#define LIBCURL_VERSION_NUM 0x[0-9a-fA-F]+"
  CURL_VERSION_NUM ${CURL_VERSION_H_CONTENTS})
string (REGEX REPLACE "[^0]+0x" "" CURL_VERSION_NUM ${CURL_VERSION_NUM})

include_regular_expression("^.*$")    # Sukender: Is it necessary?

# Setup package meta-data
# SET(PACKAGE "curl")
message(STATUS "curl version=[${CURL_VERSION}]")
# SET(PACKAGE_TARNAME "curl")
# SET(PACKAGE_NAME "curl")
# SET(PACKAGE_VERSION "-")
# SET(PACKAGE_STRING "curl-")
# SET(PACKAGE_BUGREPORT "a suitable curl mailing list => https://curl.haxx.se/mail/")
set(OPERATING_SYSTEM "${CMAKE_SYSTEM_NAME}")
set(OS "\"${CMAKE_SYSTEM_NAME}\"")

include_directories(${PROJECT_BINARY_DIR}/include/curl)
include_directories( ${CURL_SOURCE_DIR}/include )

option(CURL_WERROR "Turn compiler warnings into errors" OFF)
option(PICKY_COMPILER "Enable picky compiler options" ON)
option(BUILD_CURL_EXE "Set to ON to build curl executable." ON)
option(CURL_STATICLIB "Set to ON to build libcurl with static linking." ON)
option(ENABLE_ARES "Set to ON to enable c-ares support" OFF)
if(WIN32)
  option(CURL_STATIC_CRT "Set to ON to build libcurl with static CRT on Windows (/MT)." OFF)
  option(ENABLE_INET_PTON "Set to OFF to prevent usage of inet_pton when building against modern SDKs while still requiring compatibility with older Windows versions, such as Windows XP, Windows Server 2003 etc." ON)
endif()

CMAKE_DEPENDENT_OPTION(ENABLE_THREADED_RESOLVER "Set to ON to enable threaded DNS lookup"
        ON "NOT ENABLE_ARES"
        OFF)

option(ENABLE_DEBUG "Set to ON to enable curl debug features" OFF)
option(ENABLE_CURLDEBUG "Set to ON to build with TrackMemory feature enabled" OFF)

if(CMAKE_COMPILER_IS_GNUCC OR CMAKE_COMPILER_IS_CLANG)
  if (PICKY_COMPILER)
    foreach (_CCOPT -pedantic -Wall -W -Wpointer-arith -Wwrite-strings -Wunused -Wshadow -Winline -Wnested-externs -Wmissing-declarations -Wmissing-prototypes -Wno-long-long -Wfloat-equal -Wno-multichar -Wsign-compare -Wundef -Wno-format-nonliteral -Wendif-labels -Wstrict-prototypes -Wdeclaration-after-statement -Wstrict-aliasing=3 -Wcast-align -Wtype-limits -Wold-style-declaration -Wmissing-parameter-type -Wempty-body -Wclobbered -Wignored-qualifiers -Wconversion -Wno-sign-conversion -Wvla -Wdouble-promotion -Wno-system-headers)
      # surprisingly, CHECK_C_COMPILER_FLAG needs a new variable to store each new
      # test result in.
      CHECK_C_COMPILER_FLAG(${_CCOPT} OPT${_CCOPT})
      if(OPT${_CCOPT})
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} ${_CCOPT}")
      endif()
    endforeach()
  endif(PICKY_COMPILER)
endif(CMAKE_COMPILER_IS_GNUCC OR CMAKE_COMPILER_IS_CLANG)

if (ENABLE_DEBUG)
  # DEBUGBUILD will be defined only for Debug builds
  if(NOT CMAKE_VERSION VERSION_LESS 3.0)
    set_property(DIRECTORY APPEND PROPERTY COMPILE_DEFINITIONS $<$<CONFIG:Debug>:DEBUGBUILD>)
  else()
    set_property(DIRECTORY APPEND PROPERTY COMPILE_DEFINITIONS_DEBUG DEBUGBUILD)
  endif()
  set(ENABLE_CURLDEBUG ON)
endif()

if (ENABLE_CURLDEBUG)
  set_property(DIRECTORY APPEND PROPERTY COMPILE_DEFINITIONS CURLDEBUG)
endif()

# For debug libs and exes, add "-d" postfix
if(NOT DEFINED CMAKE_DEBUG_POSTFIX)
  set(CMAKE_DEBUG_POSTFIX "-d")
endif()

# initialize CURL_LIBS
set(CURL_LIBS "")

include(CurlSymbolHiding)

option(HTTP_ONLY "disables all protocols except HTTP (This overrides all CURL_DISABLE_* options)" ON)
mark_as_advanced(HTTP_ONLY)
option(CURL_DISABLE_FTP "disables FTP" OFF)
mark_as_advanced(CURL_DISABLE_FTP)
option(CURL_DISABLE_LDAP "disables LDAP" OFF)
mark_as_advanced(CURL_DISABLE_LDAP)
option(CURL_DISABLE_TELNET "disables Telnet" OFF)
mark_as_advanced(CURL_DISABLE_TELNET)
option(CURL_DISABLE_DICT "disables DICT" OFF)
mark_as_advanced(CURL_DISABLE_DICT)
option(CURL_DISABLE_FILE "disables FILE" OFF)
mark_as_advanced(CURL_DISABLE_FILE)
option(CURL_DISABLE_TFTP "disables TFTP" OFF)
mark_as_advanced(CURL_DISABLE_TFTP)
option(CURL_DISABLE_HTTP "disables HTTP" OFF)
mark_as_advanced(CURL_DISABLE_HTTP)

option(CURL_DISABLE_LDAPS "to disable LDAPS" OFF)
mark_as_advanced(CURL_DISABLE_LDAPS)

option(CURL_DISABLE_RTSP "to disable RTSP" OFF)
mark_as_advanced(CURL_DISABLE_RTSP)
option(CURL_DISABLE_PROXY "to disable proxy" OFF)
mark_as_advanced(CURL_DISABLE_PROXY)
option(CURL_DISABLE_POP3 "to disable POP3" OFF)
mark_as_advanced(CURL_DISABLE_POP3)
option(CURL_DISABLE_IMAP "to disable IMAP" OFF)
mark_as_advanced(CURL_DISABLE_IMAP)
option(CURL_DISABLE_SMTP "to disable SMTP" OFF)
mark_as_advanced(CURL_DISABLE_SMTP)
option(CURL_DISABLE_GOPHER "to disable Gopher" OFF)
mark_as_advanced(CURL_DISABLE_GOPHER)

if(HTTP_ONLY)
  set(CURL_DISABLE_FTP ON)
  set(CURL_DISABLE_LDAP ON)
  set(CURL_DISABLE_LDAPS ON)
  set(CURL_DISABLE_TELNET ON)
  set(CURL_DISABLE_DICT ON)
  set(CURL_DISABLE_FILE ON)
  set(CURL_DISABLE_TFTP ON)
  set(CURL_DISABLE_RTSP ON)
  set(CURL_DISABLE_POP3 ON)
  set(CURL_DISABLE_IMAP ON)
  set(CURL_DISABLE_SMTP ON)
  set(CURL_DISABLE_GOPHER ON)
endif()

option(CURL_DISABLE_COOKIES "to disable cookies support" OFF)
mark_as_advanced(CURL_DISABLE_COOKIES)

option(CURL_DISABLE_CRYPTO_AUTH "to disable cryptographic authentication" OFF)
mark_as_advanced(CURL_DISABLE_CRYPTO_AUTH)
option(CURL_DISABLE_VERBOSE_STRINGS "to disable verbose strings" OFF)
mark_as_advanced(CURL_DISABLE_VERBOSE_STRINGS)
option(ENABLE_IPV6 "Define if you want to enable IPv6 support" ON)
mark_as_advanced(ENABLE_IPV6)
if(ENABLE_IPV6 AND NOT WIN32)
  include(CheckStructHasMember)
  check_struct_has_member("struct sockaddr_in6" sin6_addr "netinet/in.h"
                          HAVE_SOCKADDR_IN6_SIN6_ADDR)
  check_struct_has_member("struct sockaddr_in6" sin6_scope_id "netinet/in.h"
                          HAVE_SOCKADDR_IN6_SIN6_SCOPE_ID)
  if(NOT HAVE_SOCKADDR_IN6_SIN6_ADDR)
    message(WARNING "struct sockaddr_in6 not available, disabling IPv6 support")
    # Force the feature off as this name is used as guard macro...
    set(ENABLE_IPV6 OFF
        CACHE BOOL "Define if you want to enable IPv6 support" FORCE)
  endif()
endif()

CURL_NROFF_CHECK()
find_package(Perl)

CMAKE_DEPENDENT_OPTION(ENABLE_MANUAL "to provide the built-in manual"
    ON "NROFF_USEFUL;PERL_FOUND"
    OFF)

if(NOT PERL_FOUND)
  message(STATUS "Perl not found, testing disabled.")
  set(BUILD_TESTING OFF)
endif()
if(ENABLE_MANUAL)
  set(USE_MANUAL ON)
endif()

# We need ansi c-flags, especially on HP
set(CMAKE_C_FLAGS "${CMAKE_ANSI_CFLAGS} ${CMAKE_C_FLAGS}")
set(CMAKE_REQUIRED_FLAGS ${CMAKE_ANSI_CFLAGS})

if(CURL_STATIC_CRT)
  set(CMAKE_C_FLAGS_RELEASE "${CMAKE_C_FLAGS_RELEASE} /MT")
  set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} /MTd")
endif()

# Disable warnings on Borland to avoid changing 3rd party code.
if(BORLAND)
  set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -w-")
endif(BORLAND)

# If we are on AIX, do the _ALL_SOURCE magic
if(${CMAKE_SYSTEM_NAME} MATCHES AIX)
  set(_ALL_SOURCE 1)
endif(${CMAKE_SYSTEM_NAME} MATCHES AIX)

# Include all the necessary files for macros
include (CheckFunctionExists)
include (CheckIncludeFile)
include (CheckIncludeFiles)
include (CheckLibraryExists)
include (CheckSymbolExists)
include (CheckTypeSize)
include (CheckCSourceCompiles)
include (CMakeDependentOption)

# On windows preload settings
if(WIN32)
  set(CMAKE_REQUIRED_DEFINITIONS "${CMAKE_REQUIRED_DEFINITIONS} -D_WINSOCKAPI_=")
  include(${CMAKE_CURRENT_SOURCE_DIR}/CMake/Platforms/WindowsCache.cmake)
endif(WIN32)

if(ENABLE_THREADED_RESOLVER)
  find_package(Threads REQUIRED)
  if(WIN32)
    set(USE_THREADS_WIN32 ON)
  else()
    set(USE_THREADS_POSIX ${CMAKE_USE_PTHREADS_INIT})
    set(HAVE_PTHREAD_H ${CMAKE_USE_PTHREADS_INIT})
  endif()
  set(CURL_LIBS ${CURL_LIBS} ${CMAKE_THREAD_LIBS_INIT})
endif()

# Check for all needed libraries
check_library_exists_concat("dl"     dlopen       HAVE_LIBDL)
check_library_exists_concat("socket" connect      HAVE_LIBSOCKET)
check_library_exists("c" gethostbyname "" NOT_NEED_LIBNSL)

# Yellowtab Zeta needs different libraries than BeOS 5.
if(BEOS)
  set(NOT_NEED_LIBNSL 1)
  check_library_exists_concat("bind" gethostbyname HAVE_LIBBIND)
  check_library_exists_concat("bnetapi" closesocket HAVE_LIBBNETAPI)
endif(BEOS)

if(NOT NOT_NEED_LIBNSL)
  check_library_exists_concat("nsl"    gethostbyname  HAVE_LIBNSL)
endif(NOT NOT_NEED_LIBNSL)

check_function_exists(gethostname HAVE_GETHOSTNAME)

if(WIN32)
  check_library_exists_concat("ws2_32" getch        HAVE_LIBWS2_32)
  check_library_exists_concat("winmm"  getch        HAVE_LIBWINMM)
  list(APPEND CURL_LIBS "advapi32")
endif()

# check SSL libraries
# TODO support GNUTLS, NSS, POLARSSL, AXTLS, CYASSL
IF(NOT GNUTLS_FOUND)
  FIND_PACKAGE(GnuTLS REQUIRED)
ENDIF()
IF(GNUTLS_FOUND)
  include_directories(${GNUTLS_INCLUDE_DIRS})
ENDIF(GNUTLS_FOUND)

IF(NOT GCRYPT_FOUND)
  FIND_PACKAGE(GCRYPT REQUIRED)
ENDIF()
IF(GCRYPT_FOUND)
  include_directories(${GCRYPT_INCLUDE_DIRS})
ENDIF(GCRYPT_FOUND)

set(USE_GNUTLS ON)
set(USE_GNUTLS_NETTLE ON)

# Check for idn
check_library_exists_concat("idn2" idn2_lookup_ul HAVE_LIBIDN2)

# Check for symbol dlopen (same as HAVE_LIBDL)
check_library_exists("${CURL_LIBS}" dlopen "" HAVE_DLOPEN)

option(ENABLE_UNIX_SOCKETS "Define if you want Unix domain sockets support" ON)
if(ENABLE_UNIX_SOCKETS)
  include(CheckStructHasMember)
  check_struct_has_member("struct sockaddr_un" sun_path "sys/un.h" USE_UNIX_SOCKETS)
else()
  unset(USE_UNIX_SOCKETS CACHE)
endif()

set(CURL_CA_BUNDLE "auto" CACHE STRING
    "Path to the CA bundle. Set 'none' to disable or 'auto' for auto-detection. Defaults to 'auto'.")
set(CURL_CA_FALLBACK OFF CACHE BOOL
    "Set ON to use built-in CA store of TLS backend. Defaults to OFF")
set(CURL_CA_PATH "auto" CACHE STRING
    "Location of default CA path. Set 'none' to disable or 'auto' for auto-detection. Defaults to 'auto'.")

if("${CURL_CA_BUNDLE}" STREQUAL "")
    message(FATAL_ERROR "Invalid value of CURL_CA_BUNDLE. Use 'none', 'auto' or file path.")
elseif("${CURL_CA_BUNDLE}" STREQUAL "none")
    unset(CURL_CA_BUNDLE CACHE)
elseif("${CURL_CA_BUNDLE}" STREQUAL "auto")
    unset(CURL_CA_BUNDLE CACHE)
    set(CURL_CA_BUNDLE_AUTODETECT TRUE)
else()
    set(CURL_CA_BUNDLE_SET TRUE)
endif()

if("${CURL_CA_PATH}" STREQUAL "")
    message(FATAL_ERROR "Invalid value of CURL_CA_PATH. Use 'none', 'auto' or directory path.")
elseif("${CURL_CA_PATH}" STREQUAL "none")
    unset(CURL_CA_PATH CACHE)
elseif("${CURL_CA_PATH}" STREQUAL "auto")
    unset(CURL_CA_PATH CACHE)
    set(CURL_CA_PATH_AUTODETECT TRUE)
else()
    set(CURL_CA_PATH_SET TRUE)
endif()

if(CURL_CA_BUNDLE_SET AND CURL_CA_PATH_AUTODETECT)
    # Skip autodetection of unset CA path because CA bundle is set explicitly
elseif(CURL_CA_PATH_SET AND CURL_CA_BUNDLE_AUTODETECT)
    # Skip autodetection of unset CA bundle because CA path is set explicitly
elseif(CURL_CA_PATH_AUTODETECT OR CURL_CA_BUNDLE_AUTODETECT)
    # first try autodetecting a CA bundle, then a CA path

    if(CURL_CA_BUNDLE_AUTODETECT)
        set(SEARCH_CA_BUNDLE_PATHS
            /etc/ssl/certs/ca-certificates.crt
            /etc/pki/tls/certs/ca-bundle.crt
            /usr/share/ssl/certs/ca-bundle.crt
            /usr/local/share/certs/ca-root-nss.crt
            /etc/ssl/cert.pem)

        foreach(SEARCH_CA_BUNDLE_PATH ${SEARCH_CA_BUNDLE_PATHS})
            if(EXISTS "${SEARCH_CA_BUNDLE_PATH}")
                message(STATUS "Found CA bundle: ${SEARCH_CA_BUNDLE_PATH}")
                set(CURL_CA_BUNDLE "${SEARCH_CA_BUNDLE_PATH}")
                set(CURL_CA_BUNDLE_SET TRUE CACHE BOOL "Path to the CA bundle has been set")
                break()
            endif()
        endforeach()
    endif()

    if(CURL_CA_PATH_AUTODETECT AND (NOT CURL_CA_PATH_SET))
        if(EXISTS "/etc/ssl/certs")
            set(CURL_CA_PATH "/etc/ssl/certs")
            set(CURL_CA_PATH_SET TRUE CACHE BOOL "Path to the CA bundle has been set")
        endif()
    endif()
endif()

# Check for header files
if(NOT UNIX)
  check_include_file_concat("windows.h"      HAVE_WINDOWS_H)
  check_include_file_concat("winsock.h"      HAVE_WINSOCK_H)
  check_include_file_concat("ws2tcpip.h"     HAVE_WS2TCPIP_H)
  check_include_file_concat("winsock2.h"     HAVE_WINSOCK2_H)
  if(NOT CURL_WINDOWS_SSPI AND USE_OPENSSL)
    set(CURL_LIBS ${CURL_LIBS} "crypt32")
  endif()
endif(NOT UNIX)

check_include_file_concat("stdio.h"          HAVE_STDIO_H)
check_include_file_concat("inttypes.h"       HAVE_INTTYPES_H)
check_include_file_concat("sys/filio.h"      HAVE_SYS_FILIO_H)
check_include_file_concat("sys/ioctl.h"      HAVE_SYS_IOCTL_H)
check_include_file_concat("sys/param.h"      HAVE_SYS_PARAM_H)
check_include_file_concat("sys/poll.h"       HAVE_SYS_POLL_H)
check_include_file_concat("sys/resource.h"   HAVE_SYS_RESOURCE_H)
check_include_file_concat("sys/select.h"     HAVE_SYS_SELECT_H)
check_include_file_concat("sys/socket.h"     HAVE_SYS_SOCKET_H)
check_include_file_concat("sys/sockio.h"     HAVE_SYS_SOCKIO_H)
check_include_file_concat("sys/stat.h"       HAVE_SYS_STAT_H)
check_include_file_concat("sys/time.h"       HAVE_SYS_TIME_H)
check_include_file_concat("sys/types.h"      HAVE_SYS_TYPES_H)
check_include_file_concat("sys/uio.h"        HAVE_SYS_UIO_H)
check_include_file_concat("sys/un.h"         HAVE_SYS_UN_H)
check_include_file_concat("sys/utime.h"      HAVE_SYS_UTIME_H)
check_include_file_concat("sys/xattr.h"      HAVE_SYS_XATTR_H)
check_include_file_concat("alloca.h"         HAVE_ALLOCA_H)
check_include_file_concat("arpa/inet.h"      HAVE_ARPA_INET_H)
check_include_file_concat("arpa/tftp.h"      HAVE_ARPA_TFTP_H)
check_include_file_concat("assert.h"         HAVE_ASSERT_H)
check_include_file_concat("crypto.h"         HAVE_CRYPTO_H)
check_include_file_concat("des.h"            HAVE_DES_H)
check_include_file_concat("err.h"            HAVE_ERR_H)
check_include_file_concat("errno.h"          HAVE_ERRNO_H)
check_include_file_concat("fcntl.h"          HAVE_FCNTL_H)
check_include_file_concat("idn2.h"           HAVE_IDN2_H)
check_include_file_concat("ifaddrs.h"        HAVE_IFADDRS_H)
check_include_file_concat("io.h"             HAVE_IO_H)
check_include_file_concat("krb.h"            HAVE_KRB_H)
check_include_file_concat("libgen.h"         HAVE_LIBGEN_H)
check_include_file_concat("locale.h"         HAVE_LOCALE_H)
check_include_file_concat("net/if.h"         HAVE_NET_IF_H)
check_include_file_concat("netdb.h"          HAVE_NETDB_H)
check_include_file_concat("netinet/in.h"     HAVE_NETINET_IN_H)
check_include_file_concat("netinet/tcp.h"    HAVE_NETINET_TCP_H)

check_include_file_concat("pem.h"            HAVE_PEM_H)
check_include_file_concat("poll.h"           HAVE_POLL_H)
check_include_file_concat("pwd.h"            HAVE_PWD_H)
check_include_file_concat("rsa.h"            HAVE_RSA_H)
check_include_file_concat("setjmp.h"         HAVE_SETJMP_H)
check_include_file_concat("sgtty.h"          HAVE_SGTTY_H)
check_include_file_concat("signal.h"         HAVE_SIGNAL_H)
check_include_file_concat("ssl.h"            HAVE_SSL_H)
check_include_file_concat("stdbool.h"        HAVE_STDBOOL_H)
check_include_file_concat("stdint.h"         HAVE_STDINT_H)
check_include_file_concat("stdio.h"          HAVE_STDIO_H)
check_include_file_concat("stdlib.h"         HAVE_STDLIB_H)
check_include_file_concat("string.h"         HAVE_STRING_H)
check_include_file_concat("strings.h"        HAVE_STRINGS_H)
check_include_file_concat("stropts.h"        HAVE_STROPTS_H)
check_include_file_concat("termio.h"         HAVE_TERMIO_H)
check_include_file_concat("termios.h"        HAVE_TERMIOS_H)
check_include_file_concat("time.h"           HAVE_TIME_H)
check_include_file_concat("unistd.h"         HAVE_UNISTD_H)
check_include_file_concat("utime.h"          HAVE_UTIME_H)
check_include_file_concat("x509.h"           HAVE_X509_H)

check_include_file_concat("process.h"        HAVE_PROCESS_H)
check_include_file_concat("stddef.h"         HAVE_STDDEF_H)
check_include_file_concat("dlfcn.h"          HAVE_DLFCN_H)
check_include_file_concat("malloc.h"         HAVE_MALLOC_H)
check_include_file_concat("memory.h"         HAVE_MEMORY_H)
check_include_file_concat("netinet/if_ether.h" HAVE_NETINET_IF_ETHER_H)
check_include_file_concat("stdint.h"        HAVE_STDINT_H)
check_include_file_concat("sockio.h"        HAVE_SOCKIO_H)
check_include_file_concat("sys/utsname.h"   HAVE_SYS_UTSNAME_H)

check_type_size(size_t  SIZEOF_SIZE_T)
check_type_size(ssize_t  SIZEOF_SSIZE_T)
check_type_size("long long"  SIZEOF_LONG_LONG)
check_type_size("long"  SIZEOF_LONG)
check_type_size("short"  SIZEOF_SHORT)
check_type_size("int"  SIZEOF_INT)
check_type_size("__int64"  SIZEOF___INT64)
check_type_size("long double"  SIZEOF_LONG_DOUBLE)
check_type_size("time_t"  SIZEOF_TIME_T)
if(NOT HAVE_SIZEOF_SSIZE_T)
  if(SIZEOF_LONG EQUAL SIZEOF_SIZE_T)
    set(ssize_t long)
  endif(SIZEOF_LONG EQUAL SIZEOF_SIZE_T)
  if(NOT ssize_t AND SIZEOF___INT64 EQUAL SIZEOF_SIZE_T)
    set(ssize_t __int64)
  endif(NOT ssize_t AND SIZEOF___INT64 EQUAL SIZEOF_SIZE_T)
endif(NOT HAVE_SIZEOF_SSIZE_T)
# off_t is sized later, after the HAVE_FILE_OFFSET_BITS test

if(HAVE_SIZEOF_LONG_LONG)
  set(HAVE_LONGLONG 1)
  set(HAVE_LL 1)
endif(HAVE_SIZEOF_LONG_LONG)

find_file(RANDOM_FILE urandom /dev)
mark_as_advanced(RANDOM_FILE)

# Check for some functions that are used
if(HAVE_LIBWS2_32)
  set(CMAKE_REQUIRED_LIBRARIES ws2_32)
elseif(HAVE_LIBSOCKET)
  set(CMAKE_REQUIRED_LIBRARIES socket)
endif()

check_symbol_exists(basename      "${CURL_INCLUDES}" HAVE_BASENAME)
check_symbol_exists(socket        "${CURL_INCLUDES}" HAVE_SOCKET)
# poll on macOS is unreliable, it first did not exist, then was broken until
# fixed in 10.9 only to break again in 10.12.
if(NOT APPLE)
  check_symbol_exists(poll        "${CURL_INCLUDES}" HAVE_POLL)
endif()
check_symbol_exists(select        "${CURL_INCLUDES}" HAVE_SELECT)
check_symbol_exists(strdup        "${CURL_INCLUDES}" HAVE_STRDUP)
check_symbol_exists(strstr        "${CURL_INCLUDES}" HAVE_STRSTR)
check_symbol_exists(strtok_r      "${CURL_INCLUDES}" HAVE_STRTOK_R)
check_symbol_exists(strftime      "${CURL_INCLUDES}" HAVE_STRFTIME)
check_symbol_exists(uname         "${CURL_INCLUDES}" HAVE_UNAME)
check_symbol_exists(strcasecmp    "${CURL_INCLUDES}" HAVE_STRCASECMP)
check_symbol_exists(stricmp       "${CURL_INCLUDES}" HAVE_STRICMP)
check_symbol_exists(strcmpi       "${CURL_INCLUDES}" HAVE_STRCMPI)
check_symbol_exists(strncmpi      "${CURL_INCLUDES}" HAVE_STRNCMPI)
check_symbol_exists(alarm         "${CURL_INCLUDES}" HAVE_ALARM)
if(NOT HAVE_STRNCMPI)
  set(HAVE_STRCMPI)
endif(NOT HAVE_STRNCMPI)
check_symbol_exists(gethostbyaddr "${CURL_INCLUDES}" HAVE_GETHOSTBYADDR)
check_symbol_exists(gethostbyaddr_r "${CURL_INCLUDES}" HAVE_GETHOSTBYADDR_R)
check_symbol_exists(gettimeofday  "${CURL_INCLUDES}" HAVE_GETTIMEOFDAY)
check_symbol_exists(inet_addr     "${CURL_INCLUDES}" HAVE_INET_ADDR)
check_symbol_exists(inet_ntoa     "${CURL_INCLUDES}" HAVE_INET_NTOA)
check_symbol_exists(inet_ntoa_r   "${CURL_INCLUDES}" HAVE_INET_NTOA_R)
check_symbol_exists(tcsetattr     "${CURL_INCLUDES}" HAVE_TCSETATTR)
check_symbol_exists(tcgetattr     "${CURL_INCLUDES}" HAVE_TCGETATTR)
check_symbol_exists(perror        "${CURL_INCLUDES}" HAVE_PERROR)
check_symbol_exists(closesocket   "${CURL_INCLUDES}" HAVE_CLOSESOCKET)
check_symbol_exists(setvbuf       "${CURL_INCLUDES}" HAVE_SETVBUF)
check_symbol_exists(sigsetjmp     "${CURL_INCLUDES}" HAVE_SIGSETJMP)
check_symbol_exists(getpass_r     "${CURL_INCLUDES}" HAVE_GETPASS_R)
check_symbol_exists(strlcat       "${CURL_INCLUDES}" HAVE_STRLCAT)
check_symbol_exists(getpwuid      "${CURL_INCLUDES}" HAVE_GETPWUID)
check_symbol_exists(getpwuid_r    "${CURL_INCLUDES}" HAVE_GETPWUID_R)
check_symbol_exists(geteuid       "${CURL_INCLUDES}" HAVE_GETEUID)
check_symbol_exists(utime         "${CURL_INCLUDES}" HAVE_UTIME)
check_symbol_exists(gmtime_r      "${CURL_INCLUDES}" HAVE_GMTIME_R)
check_symbol_exists(localtime_r   "${CURL_INCLUDES}" HAVE_LOCALTIME_R)

check_symbol_exists(gethostbyname   "${CURL_INCLUDES}" HAVE_GETHOSTBYNAME)
check_symbol_exists(gethostbyname_r "${CURL_INCLUDES}" HAVE_GETHOSTBYNAME_R)

check_symbol_exists(signal        "${CURL_INCLUDES}" HAVE_SIGNAL_FUNC)
check_symbol_exists(SIGALRM       "${CURL_INCLUDES}" HAVE_SIGNAL_MACRO)
if(HAVE_SIGNAL_FUNC AND HAVE_SIGNAL_MACRO)
  set(HAVE_SIGNAL 1)
endif(HAVE_SIGNAL_FUNC AND HAVE_SIGNAL_MACRO)
check_symbol_exists(uname          "${CURL_INCLUDES}" HAVE_UNAME)
check_symbol_exists(strtoll        "${CURL_INCLUDES}" HAVE_STRTOLL)
check_symbol_exists(_strtoi64      "${CURL_INCLUDES}" HAVE__STRTOI64)
check_symbol_exists(strerror_r     "${CURL_INCLUDES}" HAVE_STRERROR_R)
check_symbol_exists(siginterrupt   "${CURL_INCLUDES}" HAVE_SIGINTERRUPT)
check_symbol_exists(perror         "${CURL_INCLUDES}" HAVE_PERROR)
check_symbol_exists(fork           "${CURL_INCLUDES}" HAVE_FORK)
check_symbol_exists(getaddrinfo    "${CURL_INCLUDES}" HAVE_GETADDRINFO)
check_symbol_exists(freeaddrinfo   "${CURL_INCLUDES}" HAVE_FREEADDRINFO)
check_symbol_exists(freeifaddrs    "${CURL_INCLUDES}" HAVE_FREEIFADDRS)
check_symbol_exists(pipe           "${CURL_INCLUDES}" HAVE_PIPE)
check_symbol_exists(ftruncate      "${CURL_INCLUDES}" HAVE_FTRUNCATE)
check_symbol_exists(getprotobyname "${CURL_INCLUDES}" HAVE_GETPROTOBYNAME)
check_symbol_exists(getrlimit      "${CURL_INCLUDES}" HAVE_GETRLIMIT)
check_symbol_exists(setlocale      "${CURL_INCLUDES}" HAVE_SETLOCALE)
check_symbol_exists(setmode        "${CURL_INCLUDES}" HAVE_SETMODE)
check_symbol_exists(setrlimit      "${CURL_INCLUDES}" HAVE_SETRLIMIT)
check_symbol_exists(fcntl          "${CURL_INCLUDES}" HAVE_FCNTL)
check_symbol_exists(ioctl          "${CURL_INCLUDES}" HAVE_IOCTL)
check_symbol_exists(setsockopt     "${CURL_INCLUDES}" HAVE_SETSOCKOPT)
check_function_exists(mach_absolute_time HAVE_MACH_ABSOLUTE_TIME)

# symbol exists in win32, but function does not.
if(WIN32)
  if(ENABLE_INET_PTON)
    check_function_exists(inet_pton HAVE_INET_PTON)
    # _WIN32_WINNT_VISTA (0x0600)
    add_definitions(-D_WIN32_WINNT=0x0600)
  else()
    # _WIN32_WINNT_WINXP (0x0501)
    add_definitions(-D_WIN32_WINNT=0x0501)
  endif()
else()
    check_function_exists(inet_pton HAVE_INET_PTON)
endif()

check_symbol_exists(fsetxattr "${CURL_INCLUDES}" HAVE_FSETXATTR)
if(HAVE_FSETXATTR)
  foreach(CURL_TEST HAVE_FSETXATTR_5 HAVE_FSETXATTR_6)
    curl_internal_test(${CURL_TEST})
  endforeach(CURL_TEST)
endif(HAVE_FSETXATTR)

# sigaction and sigsetjmp are special. Use special mechanism for
# detecting those, but only if previous attempt failed.
if(HAVE_SIGNAL_H)
  check_symbol_exists(sigaction "signal.h" HAVE_SIGACTION)
endif(HAVE_SIGNAL_H)

if(NOT HAVE_SIGSETJMP)
  if(HAVE_SETJMP_H)
    check_symbol_exists(sigsetjmp "setjmp.h" HAVE_MACRO_SIGSETJMP)
    if(HAVE_MACRO_SIGSETJMP)
      set(HAVE_SIGSETJMP 1)
    endif(HAVE_MACRO_SIGSETJMP)
  endif(HAVE_SETJMP_H)
endif(NOT HAVE_SIGSETJMP)

# If there is no stricmp(), do not allow LDAP to parse URLs
if(NOT HAVE_STRICMP)
  set(HAVE_LDAP_URL_PARSE 1)
endif(NOT HAVE_STRICMP)

# Do curl specific tests
foreach(CURL_TEST
    HAVE_FCNTL_O_NONBLOCK
    HAVE_IOCTLSOCKET
    HAVE_IOCTLSOCKET_CAMEL
    HAVE_IOCTLSOCKET_CAMEL_FIONBIO
    HAVE_IOCTLSOCKET_FIONBIO
    HAVE_IOCTL_FIONBIO
    HAVE_IOCTL_SIOCGIFADDR
    HAVE_SETSOCKOPT_SO_NONBLOCK
    HAVE_SOCKADDR_IN6_SIN6_SCOPE_ID
    TIME_WITH_SYS_TIME
    HAVE_O_NONBLOCK
    HAVE_GETHOSTBYADDR_R_5
    HAVE_GETHOSTBYADDR_R_7
    HAVE_GETHOSTBYADDR_R_8
    HAVE_GETHOSTBYADDR_R_5_REENTRANT
    HAVE_GETHOSTBYADDR_R_7_REENTRANT
    HAVE_GETHOSTBYADDR_R_8_REENTRANT
    HAVE_GETHOSTBYNAME_R_3
    HAVE_GETHOSTBYNAME_R_5
    HAVE_GETHOSTBYNAME_R_6
    HAVE_GETHOSTBYNAME_R_3_REENTRANT
    HAVE_GETHOSTBYNAME_R_5_REENTRANT
    HAVE_GETHOSTBYNAME_R_6_REENTRANT
    HAVE_SOCKLEN_T
    HAVE_IN_ADDR_T
    HAVE_BOOL_T
    STDC_HEADERS
    RETSIGTYPE_TEST
    HAVE_INET_NTOA_R_DECL
    HAVE_INET_NTOA_R_DECL_REENTRANT
    HAVE_GETADDRINFO
    HAVE_FILE_OFFSET_BITS
    )
  curl_internal_test(${CURL_TEST})
endforeach(CURL_TEST)

if(HAVE_FILE_OFFSET_BITS)
  set(_FILE_OFFSET_BITS 64)
  set(CMAKE_REQUIRED_FLAGS "-D_FILE_OFFSET_BITS=64")
endif(HAVE_FILE_OFFSET_BITS)
check_type_size("off_t"  SIZEOF_OFF_T)

# include this header to get the type
set(CMAKE_REQUIRED_INCLUDES "${CURL_SOURCE_DIR}/include")
set(CMAKE_EXTRA_INCLUDE_FILES "curl/system.h")
check_type_size("curl_off_t"  SIZEOF_CURL_OFF_T)
set(CMAKE_EXTRA_INCLUDE_FILES "")

set(CMAKE_REQUIRED_FLAGS)

foreach(CURL_TEST
    HAVE_GLIBC_STRERROR_R
    HAVE_POSIX_STRERROR_R
    )
  curl_internal_test(${CURL_TEST})
endforeach(CURL_TEST)

# Check for reentrant
foreach(CURL_TEST
    HAVE_GETHOSTBYADDR_R_5
    HAVE_GETHOSTBYADDR_R_7
    HAVE_GETHOSTBYADDR_R_8
    HAVE_GETHOSTBYNAME_R_3
    HAVE_GETHOSTBYNAME_R_5
    HAVE_GETHOSTBYNAME_R_6
    HAVE_INET_NTOA_R_DECL_REENTRANT)
  if(NOT ${CURL_TEST})
    if(${CURL_TEST}_REENTRANT)
      set(NEED_REENTRANT 1)
    endif(${CURL_TEST}_REENTRANT)
  endif(NOT ${CURL_TEST})
endforeach(CURL_TEST)

if(NEED_REENTRANT)
  foreach(CURL_TEST
      HAVE_GETHOSTBYADDR_R_5
      HAVE_GETHOSTBYADDR_R_7
      HAVE_GETHOSTBYADDR_R_8
      HAVE_GETHOSTBYNAME_R_3
      HAVE_GETHOSTBYNAME_R_5
      HAVE_GETHOSTBYNAME_R_6)
    set(${CURL_TEST} 0)
    if(${CURL_TEST}_REENTRANT)
      set(${CURL_TEST} 1)
    endif(${CURL_TEST}_REENTRANT)
  endforeach(CURL_TEST)
endif(NEED_REENTRANT)

if(HAVE_INET_NTOA_R_DECL_REENTRANT)
  set(HAVE_INET_NTOA_R_DECL 1)
  set(NEED_REENTRANT 1)
endif(HAVE_INET_NTOA_R_DECL_REENTRANT)

# Some other minor tests

if(NOT HAVE_IN_ADDR_T)
  set(in_addr_t "unsigned long")
endif(NOT HAVE_IN_ADDR_T)

# Fix libz / zlib.h

if(NOT CURL_SPECIAL_LIBZ)
  if(NOT HAVE_LIBZ)
    set(HAVE_ZLIB_H 0)
  endif(NOT HAVE_LIBZ)

  if(NOT HAVE_ZLIB_H)
    set(HAVE_LIBZ 0)
  endif(NOT HAVE_ZLIB_H)
endif(NOT CURL_SPECIAL_LIBZ)

# Check for nonblocking
set(HAVE_DISABLED_NONBLOCKING 1)
if(HAVE_FIONBIO OR
    HAVE_IOCTLSOCKET OR
    HAVE_IOCTLSOCKET_CASE OR
    HAVE_O_NONBLOCK)
  set(HAVE_DISABLED_NONBLOCKING)
endif(HAVE_FIONBIO OR
  HAVE_IOCTLSOCKET OR
  HAVE_IOCTLSOCKET_CASE OR
  HAVE_O_NONBLOCK)

if(RETSIGTYPE_TEST)
  set(RETSIGTYPE void)
else(RETSIGTYPE_TEST)
  set(RETSIGTYPE int)
endif(RETSIGTYPE_TEST)

if(CMAKE_COMPILER_IS_GNUCC AND APPLE)
  include(CheckCCompilerFlag)
  check_c_compiler_flag(-Wno-long-double HAVE_C_FLAG_Wno_long_double)
  if(HAVE_C_FLAG_Wno_long_double)
    # The Mac version of GCC warns about use of long double.  Disable it.
    get_source_file_property(MPRINTF_COMPILE_FLAGS mprintf.c COMPILE_FLAGS)
    if(MPRINTF_COMPILE_FLAGS)
      set(MPRINTF_COMPILE_FLAGS "${MPRINTF_COMPILE_FLAGS} -Wno-long-double")
    else(MPRINTF_COMPILE_FLAGS)
      set(MPRINTF_COMPILE_FLAGS "-Wno-long-double")
    endif(MPRINTF_COMPILE_FLAGS)
    set_source_files_properties(mprintf.c PROPERTIES
      COMPILE_FLAGS ${MPRINTF_COMPILE_FLAGS})
  endif(HAVE_C_FLAG_Wno_long_double)
endif(CMAKE_COMPILER_IS_GNUCC AND APPLE)

if(HAVE_SOCKLEN_T)
  set(CURL_TYPEOF_CURL_SOCKLEN_T "socklen_t")
  if(WIN32)
    set(CMAKE_EXTRA_INCLUDE_FILES "winsock2.h;ws2tcpip.h")
  elseif(HAVE_SYS_SOCKET_H)
    set(CMAKE_EXTRA_INCLUDE_FILES "sys/socket.h")
  endif()
  check_type_size("socklen_t" CURL_SIZEOF_CURL_SOCKLEN_T)
  set(CMAKE_EXTRA_INCLUDE_FILES)
  if(NOT HAVE_CURL_SIZEOF_CURL_SOCKLEN_T)
    message(FATAL_ERROR
     "Check for sizeof socklen_t failed, see CMakeFiles/CMakerror.log")
  endif()
else()
  set(CURL_TYPEOF_CURL_SOCKLEN_T int)
  set(CURL_SIZEOF_CURL_SOCKLEN_T ${SIZEOF_INT})
endif()

# TODO test which of these headers are required
if(WIN32)
  set(CURL_PULL_WS2TCPIP_H ${HAVE_WS2TCPIP_H})
else()
  set(CURL_PULL_SYS_TYPES_H ${HAVE_SYS_TYPES_H})
  set(CURL_PULL_SYS_SOCKET_H ${HAVE_SYS_SOCKET_H})
  set(CURL_PULL_SYS_POLL_H ${HAVE_SYS_POLL_H})
endif()
set(CURL_PULL_STDINT_H ${HAVE_STDINT_H})
set(CURL_PULL_INTTYPES_H ${HAVE_INTTYPES_H})

include(CMake/OtherTests.cmake)

add_definitions(-DHAVE_CONFIG_H)

# For Windows, all compilers used by CMake should support large files
if(WIN32)
  set(USE_WIN32_LARGE_FILES ON)

  # Use the manifest embedded in the Windows Resource
  set(CMAKE_RC_FLAGS "${CMAKE_RC_FLAGS} -DCURL_EMBED_MANIFEST")
endif(WIN32)

if(MSVC)
  # Disable default manifest added by CMake
  set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} /MANIFEST:NO")

  add_definitions(-D_CRT_SECURE_NO_DEPRECATE -D_CRT_NONSTDC_NO_DEPRECATE)
  if(CMAKE_C_FLAGS MATCHES "/W[0-4]")
    string(REGEX REPLACE "/W[0-4]" "/W4" CMAKE_C_FLAGS "${CMAKE_C_FLAGS}")
  else(CMAKE_C_FLAGS MATCHES "/W[0-4]")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} /W4")
  endif(CMAKE_C_FLAGS MATCHES "/W[0-4]")
endif(MSVC)

if(CURL_WERROR)
  if(MSVC_VERSION)
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} /WX")
  else()
    # this assumes clang or gcc style options
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Werror")
  endif()
endif(CURL_WERROR)

# Ugly (but functional) way to include "Makefile.inc" by transforming it (= regenerate it).
function(TRANSFORM_MAKEFILE_INC INPUT_FILE OUTPUT_FILE)
  file(READ ${INPUT_FILE} MAKEFILE_INC_TEXT)
  string(REPLACE "$(top_srcdir)"   "\${CURL_SOURCE_DIR}" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})
  string(REPLACE "$(top_builddir)" "\${CURL_BINARY_DIR}" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})

  string(REGEX REPLACE "\\\\\n" "!π!α!" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})
  string(REGEX REPLACE "([a-zA-Z_][a-zA-Z0-9_]*)[\t ]*=[\t ]*([^\n]*)" "SET(\\1 \\2)" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})
  string(REPLACE "!π!α!" "\n" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})

  string(REGEX REPLACE "\\$\\(([a-zA-Z_][a-zA-Z0-9_]*)\\)" "\${\\1}" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})    # Replace $() with ${}
  string(REGEX REPLACE "@([a-zA-Z_][a-zA-Z0-9_]*)@" "\${\\1}" MAKEFILE_INC_TEXT ${MAKEFILE_INC_TEXT})    # Replace @@ with ${}, even if that may not be read by CMake scripts.
  file(WRITE ${OUTPUT_FILE} ${MAKEFILE_INC_TEXT})

endfunction()

add_definitions(-DHAVE_CONFIG_H)

if(WIN32 AND NOT CYGWIN)
    set(CURL_INSTALL_CMAKE_DIR CMake)
else()
    set(CURL_INSTALL_CMAKE_DIR lib/cmake/curl)
endif()

add_subdirectory(lib)
