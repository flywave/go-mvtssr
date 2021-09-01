set(CMAKE_C_FLAGS "-Wno-shift-negative-value")
if(UNIX OR APPLE)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC -Wall")
endif()

add_definitions(-DHAVE_CONFIG_H)

SET(jpeg_HEADER_FILES
	jerror.h
	jpegint.h
	jpeglib.h
)

SET(jpeg_SOURCE_FILES
	jcapimin.c
	 jcapistd.c
	 jccoefct.c
	 jccolor.c
	 jcdctmgr.c
	 jchuff.c
	 jchuff.h
	 jcinit.c
	 jcmainct.c
	 jcmarker.c
	 jcmaster.c
	 jcomapi.c
	 jconfig.h
	 jcparam.c
	 jcphuff.c
	 jcprepct.c
	 jcsample.c
	 jdapimin.c
	 jdapistd.c
	 jdatadst.c
	 jdatasrc.c
	 jdcoefct.c
	 jdcolor.c
	 jdct.h
	 jddctmgr.c
	 jdhuff.c
	 jdhuff.h
	 jdinput.c
	 jdmainct.c
	 jdmarker.c
	 jdmaster.c
	 jdmerge.c
	 jdphuff.c
	 jdpostct.c
	 jdsample.c
	 jerror.c
	 jerror.h
	 jfdctflt.c
	 jfdctfst.c
	 jfdctint.c
	 jidctflt.c
	 jidctfst.c
	 jidctint.c
	 jidctred.c
	 jinclude.h
	 jmemmgr.c
	 jmemnobs.c
	 jmemsys.h
	 jmorecfg.h
	 jpegint.h
	 jpeglib.h
	 jquant1.c
	 jquant2.c
	 jutils.c
	 jversion.h
)

add_library (jpeg STATIC
	${jpeg_HEADER_FILES}
	${jpeg_SOURCE_FILES}
)

SOURCE_GROUP("src" FILES ${jpeg_SOURCE_FILES})
SOURCE_GROUP("include" FILES ${jpeg_HEADER_FILES})

SET_TARGET_PROPERTIES(jpeg
    PROPERTIES PUBLIC_HEADER "${headers}")

		SET_TARGET_PROPERTIES(jpeg PROPERTIES
ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(jpeg PROPERTIES 
LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(jpeg PROPERTIES
RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})

IF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
	SET_TARGET_PROPERTIES(jpeg PROPERTIES FOLDER external)
ENDIF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
