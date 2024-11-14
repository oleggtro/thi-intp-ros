# generated from ament/cmake/core/templates/nameConfig.cmake.in

# prevent multiple inclusion
if(_etagen_display_CONFIG_INCLUDED)
  # ensure to keep the found flag the same
  if(NOT DEFINED etagen_display_FOUND)
    # explicitly set it to FALSE, otherwise CMake will set it to TRUE
    set(etagen_display_FOUND FALSE)
  elseif(NOT etagen_display_FOUND)
    # use separate condition to avoid uninitialized variable warning
    set(etagen_display_FOUND FALSE)
  endif()
  return()
endif()
set(_etagen_display_CONFIG_INCLUDED TRUE)

# output package information
if(NOT etagen_display_FIND_QUIETLY)
  message(STATUS "Found etagen_display: 0.0.1 (${etagen_display_DIR})")
endif()

# warn when using a deprecated package
if(NOT "" STREQUAL "")
  set(_msg "Package 'etagen_display' is deprecated")
  # append custom deprecation text if available
  if(NOT "" STREQUAL "TRUE")
    set(_msg "${_msg} ()")
  endif()
  # optionally quiet the deprecation message
  if(NOT etagen_display_DEPRECATED_QUIET)
    message(DEPRECATION "${_msg}")
  endif()
endif()

# flag package as ament-based to distinguish it after being find_package()-ed
set(etagen_display_FOUND_AMENT_PACKAGE TRUE)

# include all config extra files
set(_extras "")
foreach(_extra ${_extras})
  include("${etagen_display_DIR}/${_extra}")
endforeach()
