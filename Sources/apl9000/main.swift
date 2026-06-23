import Foundation

// Gather the dynamic data, then render and write the profile README.
// Mirrors the original Go `main.go`.
let profile = ProfileData(
    timestamp: getCurrentTime(),
    website: await getWebsiteData(),
    rates: await getRates()
)

let filename = "README.md"
do {
    let html = try renderReadme(profile)
    try html.write(toFile: filename, atomically: true, encoding: .utf8)
    print("README written successfully to \(filename)")
} catch {
    printErr("Error generating README: \(error)")
    exit(1)
}
