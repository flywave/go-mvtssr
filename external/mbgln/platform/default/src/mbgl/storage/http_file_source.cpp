#include <mbgl/storage/http_file_source.hpp>
#include <mbgl/storage/resource.hpp>
#include <mbgl/storage/response.hpp>
#include <mbgl/util/logging.hpp>

#include <mbgl/util/util.hpp>
#include <mbgl/util/optional.hpp>
#include <mbgl/util/run_loop.hpp>
#include <mbgl/util/string.hpp>
#include <mbgl/util/timer.hpp>
#include <mbgl/util/chrono.hpp>
#include <mbgl/util/http_header.hpp>

#include <dlfcn.h>
#include <queue>
#include <map>
#include <cassert>
#include <cstring>
#include <cstdio>

namespace mbgl {

class HTTPFileSource::Impl {
public:
    Impl() {}
    ~Impl() {}
};

class HTTPRequest : public AsyncRequest {
public:
    HTTPRequest(HTTPFileSource::Impl*, Resource, FileSource::Callback);
    ~HTTPRequest() override;

private:
    static size_t headerCallback(char *buffer, size_t size, size_t nmemb, void *userp);
    static size_t writeCallback(void *contents, size_t size, size_t nmemb, void *userp);

    HTTPFileSource::Impl* context = nullptr;
    Resource resource;
    FileSource::Callback callback;

    // Will store the current response.
    std::shared_ptr<std::string> data;
    std::unique_ptr<Response> response;

    optional<std::string> retryAfter;
    optional<std::string> xRateLimitReset;
};

HTTPRequest::HTTPRequest(HTTPFileSource::Impl* context_, Resource resource_, FileSource::Callback callback_)
    : context(context_),
      resource(std::move(resource_)),
      callback(std::move(callback_)) {

}

HTTPRequest::~HTTPRequest() {
    
}

// This function is called when we have new data for a request. We just append it to the string
// containing the previous data.
size_t HTTPRequest::writeCallback(void *const contents, const size_t size, const size_t nmemb, void *userp) {
    assert(userp);
    auto impl = reinterpret_cast<HTTPRequest *>(userp);

    if (!impl->data) {
        impl->data = std::make_shared<std::string>();
    }

    impl->data->append(static_cast<char *>(contents), size * nmemb);
    return size * nmemb;
}

// Compares the beginning of the (non-zero-terminated!) data buffer with the (zero-terminated!)
// header string. If the data buffer contains the header string at the beginning, it returns
// the length of the header string == begin of the value, otherwise it returns npos.
// The comparison of the header is ASCII-case-insensitive.
size_t headerMatches(const char *const header, const char *const buffer, const size_t length) {
    const size_t headerLength = strlen(header);
    if (length < headerLength) {
        return std::string::npos;
    }
    size_t i = 0;
    while (i < length && i < headerLength && std::tolower(buffer[i]) == std::tolower(header[i])) {
        i++;
    }
    return i == headerLength ? i : std::string::npos;
}

size_t HTTPRequest::headerCallback(char *const , const size_t , const size_t , void *) {
    return 0;
}

HTTPFileSource::HTTPFileSource()
    : impl(std::make_unique<Impl>()) {
}

HTTPFileSource::~HTTPFileSource() = default;

std::unique_ptr<AsyncRequest> HTTPFileSource::request(const Resource& resource, Callback callback) {
    return std::make_unique<HTTPRequest>(impl.get(), resource, callback);
}

} // namespace mbgl
