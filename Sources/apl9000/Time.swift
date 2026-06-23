import Foundation

/// Current time in America/New_York, formatted as `yyyy-MM-dd HH:mm:ss`.
/// Equivalent to the Go layout `2006-01-02 15:04:05`.
func getCurrentTime() -> String {
    let formatter = DateFormatter()
    formatter.dateFormat = "yyyy-MM-dd HH:mm:ss"
    formatter.timeZone = TimeZone(identifier: "America/New_York")
    return formatter.string(from: Date())
}
