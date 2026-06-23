import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

/// Writes a line to standard error, mirroring the Go program's `fmt.Println`
/// diagnostics so failures stay visible in CI logs without aborting the run.
func printErr(_ message: String) {
    FileHandle.standardError.write(Data((message + "\n").utf8))
}

/// Performs an HTTP GET and returns the raw response body, or `nil` on failure.
///
/// A descriptive User-Agent is always set: it is good manners for the upstream
/// APIs, and some reject requests that omit one.
func fetchData(_ urlString: String) async -> Data? {
    guard let url = URL(string: urlString) else {
        printErr("Invalid URL: \(urlString)")
        return nil
    }

    var request = URLRequest(url: url)
    request.setValue("apl9000-readme-bot (github.com/apl9000)", forHTTPHeaderField: "User-Agent")

    // Use the completion-handler API wrapped in a continuation: it is available
    // on older macOS and on swift-corelibs-foundation (Linux CI) without the
    // availability gates of the newer async `data(for:)` overload.
    return await withCheckedContinuation { continuation in
        let task = URLSession.shared.dataTask(with: request) { data, _, error in
            if let error = error {
                printErr("The HTTP request failed with error \(error)")
                continuation.resume(returning: nil)
                return
            }
            continuation.resume(returning: data)
        }
        task.resume()
    }
}

/// Fetches `urlString` and decodes the body into `T`, returning `nil` on any
/// network or decoding error.
func fetchJSON<T: Decodable>(_ urlString: String, as type: T.Type) async -> T? {
    guard let data = await fetchData(urlString) else { return nil }
    do {
        return try JSONDecoder().decode(T.self, from: data)
    } catch {
        printErr("Error unmarshalling JSON from \(urlString): \(error)")
        return nil
    }
}
