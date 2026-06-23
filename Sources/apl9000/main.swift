import Foundation

// Gather the dynamic data, then render and write the profile README.
// Mirrors the original Go `main.go`.
let profile = ProfileData(
    timestamp: getCurrentTime(),
    website: await getWebsiteData(),
    weatherForecast: await getWeatherForecast(),
    rates: await getRates()
)

let html = renderReadme(profile)

let filename = "README.md"
do {
    try html.write(toFile: filename, atomically: true, encoding: .utf8)
    print("README written successfully to \(filename)")
} catch {
    printErr("Error writing README to file: \(error)")
    exit(1)
}
